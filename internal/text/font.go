package text

import (
	"errors"
	"io/fs"
	"path"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Font struct {
	path    string
	Width   int
	Height  int
	letters []*ebiten.Image
}

func NewFont(path string) *Font {
	return &Font{path: path}
}

func (f *Font) Load() error {
	letters := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	for _, letter := range letters {
		img, _, err := ebitenutil.NewImageFromFile(path.Join(f.path, string(letter)+".png"))
		if err != nil {
			if errors.Is(err, fs.ErrNotExist) {
				continue
			}
			return err
		}
		f.letters = append(f.letters, img)
		f.Width = img.Bounds().Dx()
		f.Height = img.Bounds().Dy()
	}
	return nil
}
