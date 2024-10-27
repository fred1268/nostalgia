package demo

import "nostalgia/internal/gfx"

func GetStarFieldDefinition() *gfx.SFDefinition {
	return &gfx.SFDefinition{
		Start: 0,
		Stop:  -1,
		Stars: 500,
	}
}
