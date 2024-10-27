package demo

import "nostalgia/internal/text"

func GetRotateDefinition(font *text.Font) *text.RTDefinition {
	return &text.RTDefinition{
		Text:  "NOSTALGIA",
		Font:  font,
		Start: 120,
		Stop:  -1,
		Scale: 0.4,
	}
}
