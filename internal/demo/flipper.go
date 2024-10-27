package demo

import "nostalgia/internal/gfx"

func GetFlipperDefinition() *gfx.FDefinition {
	return &gfx.FDefinition{
		Start: 60 * 5,
		Stop:  -1,
	}
}
