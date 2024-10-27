package main

import (
	"bytes"
	"errors"
	"log"
	"os"
	"path"

	"nostalgia/internal/cfg"
	"nostalgia/internal/demo"
	"nostalgia/internal/gfx"
	"nostalgia/internal/text"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/audio"
	"github.com/hajimehoshi/ebiten/v2/audio/vorbis"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type entity interface {
	Update(step int) error
	Draw(screen *ebiten.Image, step int)
	Layout(width, height int) (int, int)
}

type Game struct {
	step       int
	stop       chan struct{}
	entities   []entity
	RotateText *text.RotateText
}

func newGame() *Game {
	return &Game{}
}

func (g *Game) Update() error {
	g.step++
	for _, entity := range g.entities {
		if err := entity.Update(g.step); err != nil {
			return err
		}
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	for _, entity := range g.entities {
		entity.Draw(screen, g.step)
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	var w, h int
	for _, entity := range g.entities {
		w, h = entity.Layout(outsideWidth, outsideHeight)
	}
	return w, h
}

func (g *Game) playback(path string, stop chan struct{}) error {
	audioContext := audio.NewContext(44100)
	fi, err := os.Stat(path)
	if err != nil {
		return err
	}
	buf := make([]byte, fi.Size())
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	n, err := file.Read(buf)
	if err != nil {
		return err
	}
	if n != int(fi.Size()) {
		return errors.New("file size mismatch")
	}
	stream, err := vorbis.DecodeF32(bytes.NewReader(buf))
	if err != nil {
		return err
	}
	player, err := audioContext.NewPlayerF32(stream)
	if err != nil {
		return err
	}
	for {
		select {
		case <-stop:
			if err := player.Close(); err != nil {
				return err
			}
			if err := file.Close(); err != nil {
				return err
			}
			return nil
		default:
			player.Play()
			if !player.IsPlaying() {
				if err := player.SetPosition(0); err != nil {
					return err
				}
			}
		}
	}
}

func main() {
	ebiten.SetWindowSize(cfg.WindowWidth, cfg.WindowHeight)
	ebiten.SetWindowTitle("Nostalgia - A tribute to the Dark Priests")
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)
	game := newGame()

	game.stop = make(chan struct{})

	go game.playback("assets/music/Of Far Different Nature - Make It Fit (CC-BY).ogg", game.stop)

	font := text.NewFont("assets/font")
	if err := font.Load(); err != nil {
		log.Fatal(err)
	}

	sprite1, _, err := ebitenutil.NewImageFromFile(path.Join("assets/images/sprite64.png"))
	if err != nil {
		log.Fatal(err)
	}

	sprite2, _, err := ebitenutil.NewImageFromFile(path.Join("assets/images/sprite48.png"))
	if err != nil {
		log.Fatal(err)
	}

	sText := text.NewScrollText(demo.GetBounceDefinition(font))
	game.entities = append(game.entities, sText)

	sText = text.NewScrollText(demo.GetWriggleDefinition(font))
	game.entities = append(game.entities, sText)

	rText := text.NewRotateText(demo.GetRotateDefinition(font))
	game.entities = append(game.entities, rText)

	// f := gfx.NewFlipper(demo.GetFlipperDefinition())
	// game.entities = append(game.entities, f)

	sf := gfx.NewStarField(demo.GetStarFieldDefinition())
	game.entities = append(game.entities, sf)

	sp := gfx.NewSprites(demo.GetSpritesDefinition1(sprite1))
	game.entities = append(game.entities, sp)

	sp = gfx.NewSprites(demo.GetSpritesDefinition2(sprite2))
	game.entities = append(game.entities, sp)

	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}

	close(game.stop)
}
