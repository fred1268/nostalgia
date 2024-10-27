package demo

import (
	"math"

	"nostalgia/internal/text"

	"github.com/hajimehoshi/ebiten/v2"
)

func bounce(def *text.STDefinition, theta float64) float64 {
	return def.Amplitude * (math.Sin(theta) / (2 + math.Cos(theta+math.Pi/3)))
}

func bottomBounce(def *text.STDefinition, windowHeight int) {
	def.Y = windowHeight - def.Font.Height - int(def.Amplitude)
}

func stepBounce(def *text.STDefinition, opts *ebiten.DrawImageOptions, theta float64) {
	opts.GeoM.Scale(1, 1+0.75*math.Sin(theta))
}

func GetBounceDefinition(font *text.Font) *text.STDefinition {
	return &text.STDefinition{
		Text:           "A TRIBUTE TO THE DARK PRIESTS     ",
		Font:           font,
		Start:          240,
		Stop:           -1,
		SliceWidth:     2,
		StepInc:        5,
		Amplitude:      120,
		Speed:          100.0,
		TextOscillator: 0.125,
		CharOscillator: 0.008,
		CurveFn:        bounce,
		YFn:            bottomBounce,
		StepFn:         stepBounce,
	}
}
