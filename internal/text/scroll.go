package text

import (
	"fmt"
	"image"

	"nostalgia/internal/cfg"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

const stDebug = false

type STDefinition struct {
	Text           string
	Font           *Font
	Start          int
	Stop           int
	Y              int
	SliceWidth     int
	StepInc        int
	Amplitude      float64
	Speed          float64
	TextOscillator float64
	CharOscillator float64
	CurveFn        CurveFunc
	YFn            YFunc
	StepFn         StepFunc
	ColorFn        ColorFunc
}

type stIteration struct {
	step       int
	theta      float64
	scrOffset  int
	fontOffset int
	firstChar  int
	chars      int
	opts       *ebiten.DrawImageOptions
}

type ScrollText struct {
	windowWidth  int
	windowHeight int
	current      stIteration
	def          *STDefinition
}

func NewScrollText(def *STDefinition) *ScrollText {
	t := &ScrollText{
		windowWidth:  cfg.WindowWidth,
		windowHeight: cfg.WindowHeight,
		current: stIteration{
			opts: &ebiten.DrawImageOptions{},
		},
		def: def,
	}
	t.def.YFn(t.def, t.windowHeight)
	return t
}

func (t *ScrollText) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	t.windowWidth, t.windowHeight = ebiten.WindowSize()
	t.def.YFn(t.def, t.windowHeight)
	return t.windowWidth, t.windowHeight
}

func (t *ScrollText) Update(step int) error {
	if step < t.def.Start || (t.def.Stop != -1 && t.def.Stop < step) {
		return nil
	}
	if stDebug {
		if ebiten.IsKeyPressed(ebiten.KeyRight) {
			t.current.step -= t.def.StepInc
		}
		if ebiten.IsKeyPressed(ebiten.KeyLeft) {
			t.current.step += t.def.StepInc
		}
		if ebiten.IsKeyPressed(ebiten.KeySpace) {
			if t.def.StepInc == 0 {
				t.def.StepInc = 5
			} else {
				t.def.StepInc = 0
			}
		}
	}
	t.current.step += t.def.StepInc
	t.current.theta += t.def.TextOscillator * t.def.Speed / 100
	t.current.scrOffset = max(0, t.windowWidth-t.current.step)
	t.current.fontOffset = max(0, t.current.step-t.windowWidth) % t.def.Font.Width
	t.current.firstChar = (max(0, t.current.step-t.windowWidth) / t.def.Font.Width) % len(t.def.Text)
	t.current.chars = (min(t.windowWidth, t.current.step)+t.current.fontOffset)/t.def.Font.Width + 1
	return nil
}

func (t *ScrollText) getChar(index int) byte {
	return t.def.Text[(index+len(t.def.Text))%len(t.def.Text)]
}

func (t *ScrollText) drawChar(screen *ebiten.Image, index int, scrOffset int) int {
	fontOffset := t.current.fontOffset
	if index != 0 {
		fontOffset = 0
	}
	if t.def.ColorFn != nil {
		t.current.opts.ColorScale.Reset()
		r, g, b := t.def.ColorFn(t.def, t.current.step)
		t.current.opts.ColorScale.Scale(r, g, b, 1)
	}
	nSlices := (t.def.Font.Width - fontOffset) / t.def.SliceWidth
	for slice := 0; slice < nSlices; slice++ {
		t.current.opts.GeoM.Reset()
		theta := t.current.theta + float64(scrOffset+t.def.SliceWidth*slice)*t.def.CharOscillator*t.def.Speed/100
		if t.def.StepFn != nil {
			t.current.opts.GeoM.Translate(-float64(t.def.Font.Width/2), -float64(t.def.Font.Height/2))
			t.def.StepFn(t.def, t.current.opts, t.current.theta)
			t.current.opts.GeoM.Translate(float64(t.def.Font.Width/2), float64(t.def.Font.Height/2))
		}
		t.current.opts.GeoM.Translate(float64(scrOffset+t.def.SliceWidth*slice), float64(t.def.Y)+t.def.CurveFn(t.def, theta))
		if t.getChar(t.current.firstChar+index) != 0x20 {
			screen.DrawImage(t.def.Font.letters[t.getChar(t.current.firstChar+index)-0x41].SubImage(image.Rect(fontOffset+t.def.SliceWidth*slice, 0, fontOffset+t.def.SliceWidth*(slice+1), t.def.Font.Height)).(*ebiten.Image), t.current.opts)
		}
	}
	return t.def.Font.Width - fontOffset
}

func (t *ScrollText) Draw(screen *ebiten.Image, step int) {
	if step < t.def.Start || (t.def.Stop != -1 && t.def.Stop < step) {
		return
	}
	if stDebug {
		text := fmt.Sprintf("step: %d, theta: %.2f, fontOffset: %d, screenOffset: %d\nfirst: %d (%c), chars: %d (%c)",
			t.current.step, t.current.theta, t.current.fontOffset, t.current.scrOffset, t.current.firstChar, t.getChar(t.current.firstChar), t.current.chars,
			t.getChar(t.current.firstChar+t.current.chars))
		ebitenutil.DebugPrint(screen, text)
	}
	scrOffset := t.current.scrOffset
	for n := 0; n < t.current.chars; n++ {
		scrOffset += t.drawChar(screen, n, scrOffset)
	}
}
