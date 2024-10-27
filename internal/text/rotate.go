package text

import (
	"math"

	"github.com/hajimehoshi/ebiten/v2"
)

const rtStepInc = 2

type RTDefinition struct {
	Text  string
	Font  *Font
	Start int
	Stop  int
	Y     int
	Scale float64
}

type rtIteration struct {
	step int
	opts *ebiten.DrawImageOptions
}

type RotateText struct {
	windowWidth  int
	windowHeight int
	current      rtIteration
	def          *RTDefinition
}

func NewRotateText(def *RTDefinition) *RotateText {
	t := &RotateText{
		current: rtIteration{
			opts: &ebiten.DrawImageOptions{},
		},
		def: def,
	}
	return t
}

func (t *RotateText) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	t.windowWidth, t.windowHeight = ebiten.WindowSize()
	return t.windowWidth, t.windowHeight
}

func (t *RotateText) Update(step int) error {
	if step < t.def.Start || (t.def.Stop != -1 && t.def.Stop < step) {
		return nil
	}
	t.current.step += rtStepInc
	return nil
}

func (t *RotateText) getChar(index int) byte {
	return t.def.Text[(index+len(t.def.Text))%len(t.def.Text)]
}

func (t *RotateText) drawChar(screen *ebiten.Image, index int) {
	ch := t.getChar(index)
	t.current.opts.GeoM.Reset()
	t.current.opts.ColorScale.Reset()
	t.current.opts.GeoM.Translate(-float64(t.def.Font.Width/2), -float64(t.def.Font.Height/2))
	t.current.opts.GeoM.Rotate(2 * math.Pi * float64(index+t.current.step%200) / float64(len(t.def.Text)+200))
	t.current.opts.GeoM.Scale(t.def.Scale, t.def.Scale)
	scale := 0.50 + math.Sin(float64(t.current.step)/100)/4
	t.current.opts.GeoM.Scale(scale, scale)
	t.current.opts.ColorScale.ScaleAlpha(float32(scale))
	t.current.opts.GeoM.Translate(float64(t.windowWidth/2)-float64((len(t.def.Text)-2*index-1)*t.def.Font.Width/2)*t.def.Scale, float64(t.windowHeight/2))
	screen.DrawImage(t.def.Font.letters[ch-0x41], t.current.opts)
}

func (t *RotateText) Draw(screen *ebiten.Image, step int) {
	if step < t.def.Start || (t.def.Stop != -1 && t.def.Stop < step) {
		return
	}
	for index := 0; index < len(t.def.Text); index++ {
		t.drawChar(screen, index)
	}
}
