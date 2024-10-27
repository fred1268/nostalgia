package demo

import (
	"math"

	"nostalgia/internal/text"

	"github.com/hajimehoshi/ebiten/v2"
)

func wriggle(def *text.STDefinition, theta float64) float64 {
	return def.Amplitude * (math.Sin(theta) + math.Cos(1.15*(theta+2*math.Pi/5))) / 1.2
}

func top(def *text.STDefinition, windowHeight int) {
	def.Y = def.Font.Height
}

func stepWriggle(def *text.STDefinition, opts *ebiten.DrawImageOptions, theta float64) {
	opts.GeoM.Scale(1, 0.5)
}

func colorWriggle(def *text.STDefinition, step int) (float32, float32, float32) {
	s := 60 - step%120
	if s < 0 {
		s = -s
	}
	f := 0.9 - float32(s)/180.0
	return 1, f, f
}

func GetWriggleDefinition(font *text.Font) *text.STDefinition {
	return &text.STDefinition{
		Text:           "STARFIELDS SINUS SCROLL TEXTS SPRITES ETC     ",
		Font:           font,
		Start:          480,
		Stop:           -1,
		SliceWidth:     2,
		StepInc:        8,
		Amplitude:      75,
		Speed:          90.0,
		TextOscillator: 0.250,
		CharOscillator: 0.008,
		CurveFn:        wriggle,
		YFn:            top,
		StepFn:         stepWriggle,
		ColorFn:        colorWriggle,
	}
}
