package demo

import (
	"math"

	"nostalgia/internal/gfx"

	"github.com/hajimehoshi/ebiten/v2"
)

func sprite1(def *gfx.SPDefinition, theta float64) float64 {
	return float64(def.Size) * math.Sin(2.25*theta) / (2 + math.Cos(3*theta))
}

func sprite2(def *gfx.SPDefinition, theta float64) float64 {
	return float64(def.Size) * math.Cos(7*theta) * math.Sin(5*theta)
}

func GetSpritesDefinition1(sprite *ebiten.Image) *gfx.SPDefinition {
	return &gfx.SPDefinition{
		Sprite:   sprite,
		Start:    780,
		Stop:     -1,
		Sprites:  16,
		Size:     360,
		Steps:    768,
		SpriteFn: sprite1,
	}
}

func GetSpritesDefinition2(sprite *ebiten.Image) *gfx.SPDefinition {
	return &gfx.SPDefinition{
		Sprite:   sprite,
		Start:    960,
		Stop:     -1,
		Sprites:  16,
		Size:     480,
		Steps:    2048,
		SpriteFn: sprite2,
	}
}
