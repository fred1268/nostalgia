package gfx

import (
	"math"
	"math/rand/v2"

	"nostalgia/internal/cfg"

	"github.com/hajimehoshi/ebiten/v2"
)

const sfNursery = 10

type SFDefinition struct {
	Start int
	Stop  int
	Stars int
}

type sfIteration struct {
	step  int
	stars []*star
	opts  *ebiten.DrawImageOptions
}

type StarField struct {
	windowWidth  int
	windowHeight int
	current      sfIteration
	def          *SFDefinition
}

func NewStarField(def *SFDefinition) *StarField {
	sf := &StarField{
		current: sfIteration{
			opts: &ebiten.DrawImageOptions{},
		},
		def: def,
	}
	for range sf.def.Stars {
		sf.current.stars = append(sf.current.stars, newStar(cfg.WindowWidth, cfg.WindowHeight))
	}
	return sf
}

func (sf *StarField) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	sf.windowWidth, sf.windowHeight = ebiten.WindowSize()
	return sf.windowWidth, sf.windowHeight
}

func (sf *StarField) Update(step int) error {
	if step < sf.def.Start || (sf.def.Stop != -1 && sf.def.Stop < step) {
		return nil
	}
	sf.current.step++
	for _, star := range sf.current.stars {
		star.rho += rand.Float64() * 3
	}
	return nil
}

func (sf *StarField) Draw(screen *ebiten.Image, step int) {
	if step < sf.def.Start || (sf.def.Stop != -1 && sf.def.Stop < step) {
		return
	}
	for _, star := range sf.current.stars {
		x := sf.windowWidth/2 + int(star.startX+star.rho*math.Cos(star.theta))
		y := sf.windowHeight/2 + int(star.startY+star.rho*math.Sin(star.theta))
		if x < 0 || x > sf.windowWidth || y < 0 || y > sf.windowHeight {
			star.randomize(sf.windowWidth/sfNursery, sf.windowHeight/sfNursery)
			x = sf.windowWidth/2 + int(star.startX+star.rho*math.Cos(star.theta))
			y = sf.windowHeight/2 + int(star.startY+star.rho*math.Sin(star.theta))
		}
		screen.Set(x, y, star.color)
	}
}
