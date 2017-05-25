package Distorted

import (
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten"
)

const (
	TILE_DIMENSION = 16
)

type sGTile struct {
	X     float64
	Y     float64
	Color color.RGBA
	Img   *ebiten.Image
}

type sGraphics struct {
	Tiles [1024]sGTile
}

func (g *sGraphics) LoadChunk(chunk [32][32]sTile) error {
	counter := 0
	for x, row := range chunk {
		for y, tile := range row {
			if len(tile.Name) == 0 {
				log.Panicf("Tile on x=%d y=%d is empty", x, y)
			}
			currColor := color.RGBA{}
			currColor = tile.Color

			cell, err := ebiten.NewImage(TILE_DIMENSION, TILE_DIMENSION, ebiten.FilterNearest)
			errcheck(err)
			cell.Fill(currColor)
			g.Tiles[counter] = sGTile{
				X:     float64((x * TILE_DIMENSION) + x),
				Y:     float64((y * TILE_DIMENSION) + y),
				Color: currColor,
				Img:   cell,
			}

			counter++
		}
	}
	return nil
}

func (g *sGraphics) update(screen *ebiten.Image) error {
	screen.Fill(color.NRGBA{0x00, 0x00, 0x00, 0xff})
	if ebiten.IsRunningSlowly() {
		return nil
	}
	for _, tile := range g.Tiles {
		opts := &ebiten.DrawImageOptions{}
		opts.GeoM.Translate(tile.X, tile.Y)
		screen.DrawImage(tile.Img, opts)
	}
	return nil
}

func (g *sGraphics) start() {
	err := ebiten.Run(g.update, (16*32)+32, (16*32)+32, 1, "MapTest")
	errcheck(err)
}
