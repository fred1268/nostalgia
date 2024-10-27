package gfx

import (
	"image/color"
	"math"
	"math/rand/v2"
)

type star struct {
	startX float64
	startY float64
	theta  float64
	rho    float64
	color  color.RGBA
}

func newStar(maxX, maxY int) *star {
	s := &star{}
	s.randomize(maxX, maxY)
	return s
}

func (s *star) randomize(maxX, maxY int) {
	s.theta = rand.Float64() * math.Pi * 2
	s.rho = rand.Float64()
	s.startX = s.rho * math.Cos(s.theta) * float64(maxX)
	s.startY = s.rho * math.Sin(s.theta) * float64(maxY)
	c := uint8(rand.IntN(128) + 128)
	s.color = color.RGBA{c, c, c, 255}
}
