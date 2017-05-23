package DEEngine

import (
	"image/color"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
)

type sGraphics struct {
	Chunk [32][32]sTile
	Tiles [1024]struct {
		Img  *ebiten.Image
		opts *ebiten.DrawImageOptions
	}
}

func (g *sGraphics) update(screen *ebiten.Image) error {
	//screen.Fill(color.NRGBA{0x00, 0x00, 0x00, 0xff})
	ebitenutil.DebugPrint(screen, "MapTest")
	counter := 0
	//currX := 1.0
	//currY := 1.0
	for x, row := range g.Chunk {
		for y, tile := range row {
			cell, err := ebiten.NewImage(16, 16, ebiten.FilterNearest)
			errcheck(err)
			switch tile.Source {
			case "@":
				cell.Fill(color.NRGBA64{R: 78, G: 88, B: 155, A: 127})
			case "#":
				cell.Fill(color.NRGBA64{R: 144, G: 173, B: 0, A: 127})
			case "%":
				cell.Fill(color.NRGBA64{R: 209, G: 178, B: 200, A: 127})
			}
			opts := &ebiten.DrawImageOptions{}
			opts.GeoM.Translate(float64(x*16), float64(y*16))
			//currY += 16
			a := &struct {
				Img  *ebiten.Image
				opts *ebiten.DrawImageOptions
			}{
				cell,
				opts,
			}
			counter++
		}
		//currX += 16
	}

	for _, tile := range g.Tiles {
		screen.DrawImage(tile.Img, tile.opts)
	}
	return nil
}

func (g *sGraphics) start() {
	err := ebiten.Run(g.update, (16*16)+16, (16*16)+16, 2, "MapTest")
	errcheck(err)
}
