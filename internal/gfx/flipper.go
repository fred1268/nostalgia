package gfx

import (
	"image/color"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
)

type FDefinition struct {
	Start int
	Stop  int
}

type fIteration struct {
	step  int
	theta float64
	opts  *ebiten.DrawImageOptions
}

type Flipper struct {
	windowWidth  int
	windowHeight int
	current      fIteration
	def          *FDefinition
}

func NewFlipper(def *FDefinition) *Flipper {
	return &Flipper{
		current: fIteration{
			opts: &ebiten.DrawImageOptions{},
		},
		def: def,
	}
}

func (f *Flipper) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	f.windowWidth, f.windowHeight = ebiten.WindowSize()
	return f.windowWidth, f.windowHeight
}

func (f *Flipper) Update(step int) error {
	if step < f.def.Start || (f.def.Stop != -1 && f.def.Stop < step) {
		return nil
	}
	f.current.step++
	f.current.theta += 0.02
	return nil
}

func (f *Flipper) Draw(screen *ebiten.Image, step int) {
	if step < f.def.Start || (f.def.Stop != -1 && f.def.Stop < step) {
		return
	}
	f.current.opts.GeoM.Reset()
	f.current.opts.GeoM.Translate(-float64(f.windowWidth)/2, -float64(f.windowHeight)/2)
	f.current.opts.GeoM.Scale(1, math.Cos(f.current.theta))
	f.current.opts.GeoM.Translate(float64(f.windowWidth)/2, float64(f.windowHeight)/2)
	scr := ebiten.NewImageFromImage(screen)
	screen.Fill(color.Black)
	screen.DrawImage(scr, f.current.opts)
}
