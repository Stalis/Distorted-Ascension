package DEEngine

import (
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
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
			if len(tile.Source) == 0 {
				log.Panicf("Tile on x=%d y=%d is empty", x, y)
			}
			currColor := color.RGBA{}
			currColor.A = 255
			switch tile.Source {
			case "#":
				currColor.R = 144
				currColor.G = 173
				currColor.B = 0
			case "@":
				currColor.R = 78
				currColor.G = 88
				currColor.B = 155
			case "%":
				currColor.R = 209
				currColor.G = 178
				currColor.B = 200
			}
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
		//tile.Img.Fill(tile.Color)
		screen.DrawImage(tile.Img, opts)
	}
	ebitenutil.DebugPrint(screen, "MapTest")
	return nil
}

func (g *sGraphics) start() {
	err := ebiten.Run(g.update, (16*32)+32, (16*32)+32, 1, "MapTest")
	errcheck(err)
}
