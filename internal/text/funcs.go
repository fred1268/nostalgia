package text

import "github.com/hajimehoshi/ebiten/v2"

type (
	CurveFunc func(def *STDefinition, theta float64) float64
	YFunc     func(def *STDefinition, windowHeight int)
	StepFunc  func(def *STDefinition, opts *ebiten.DrawImageOptions, theta float64)
	ColorFunc func(def *STDefinition, step int) (float32, float32, float32)
)
