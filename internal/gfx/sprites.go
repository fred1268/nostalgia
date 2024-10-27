package gfx

import (
	"math"

	"github.com/hajimehoshi/ebiten/v2"
)

const spAlphaStep = 0.075

type SPDefinition struct {
	Sprite   *ebiten.Image
	Start    int
	Stop     int
	Sprites  int
	Distance int
	Size     int
	Steps    int
	SpriteFn SpriteFunc
}

type spIteration struct {
	step       int
	posX       []float64
	posY       []float64
	scrOffset  int
	curStep    int
	colorShift float64
	opts       *ebiten.DrawImageOptions
}

type Sprites struct {
	windowWidth  int
	windowHeight int
	current      spIteration
	def          *SPDefinition
}

func NewSprites(def *SPDefinition) *Sprites {
	sp := &Sprites{
		current: spIteration{
			opts: &ebiten.DrawImageOptions{},
		},
		def: def,
	}
	minX := math.MaxFloat64
	maxX := -math.MaxFloat64
	sp.current.posX = make([]float64, 0, sp.def.Steps)
	sp.current.posY = make([]float64, 0, sp.def.Steps)
	for theta := 0.0; theta < 12*math.Pi; theta += 12 * math.Pi / float64(sp.def.Steps) {
		rho := sp.def.SpriteFn(sp.def, theta)
		x := rho * math.Cos(theta)
		sp.current.posX = append(sp.current.posX, x)
		sp.current.posY = append(sp.current.posY, rho*math.Sin(theta))
		if x > maxX {
			maxX = x
		} else if x < minX {
			minX = x
		}
	}
	sp.current.scrOffset = int(maxX+minX) / 2
	return sp
}

func (sp *Sprites) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	sp.windowWidth, sp.windowHeight = ebiten.WindowSize()
	return sp.windowWidth, sp.windowHeight
}

func (sp *Sprites) Update(step int) error {
	if step < sp.def.Start || (sp.def.Stop != -1 && sp.def.Stop < step) {
		return nil
	}
	sp.current.step++
	sp.current.colorShift += 0.033
	sp.current.curStep = sp.current.step / 2 % sp.def.Steps
	return nil
}

func (sp *Sprites) Draw(screen *ebiten.Image, step int) {
	if step < sp.def.Start || (sp.def.Stop != -1 && sp.def.Stop < step) {
		return
	}
	index := sp.current.curStep
	var alpha float32 = float32(sp.def.Sprites) * spAlphaStep
	for range sp.def.Sprites {
		sp.current.opts.GeoM.Reset()
		sp.current.opts.GeoM.Translate(float64(sp.windowWidth/2-sp.current.scrOffset)+sp.current.posX[index], float64(sp.windowHeight/2)+sp.current.posY[index])
		sp.current.opts.ColorScale.Reset()
		sp.current.opts.ColorScale.Scale(0.75+0.5*float32(math.Sin(0.67*sp.current.colorShift)), float32(0.75+0.5*math.Sin(0.47*sp.current.colorShift)), float32(0.75+0.5*math.Sin(1.27*sp.current.colorShift)), 0.5-alpha)
		screen.DrawImage(sp.def.Sprite, sp.current.opts)
		index = (index + 1) % sp.def.Steps
		alpha -= spAlphaStep
	}
}
