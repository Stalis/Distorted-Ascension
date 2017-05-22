package DEEngine

import (
	"fmt"
	"math/rand"
	"time"
)

type sTileType struct {
	Name       string
	Prevalence int32
	Source     string
}

func NewTileType(Name string, Prevalence int32, Source string) *sTileType {
	if Prevalence <= 0 {
		return nil
	}
	if len(Name) == 0 || len(Source) == 0 {
		return nil
	}
	return &sTileType{Name: Name, Prevalence: Prevalence, Source: Source}
}

type sTile struct {
	sTileType
}

func NewTile(Type sTileType) *sTile {
	return &sTile{Type}
}

type sChunk struct {
	Map           [32][32]sTile
	TileSource    [3]string
	TileTypeNames [3]string
	TileTypes     [3]sTileType
}

func (s sChunk) GetMap() [32][32]sTile {
	return s.Map
}

func (s *sChunk) Generate() {
	var WeightSum int32
	for _, t := range s.TileTypes {
		WeightSum += t.Prevalence
	}
	random := rand.New(rand.NewSource(int64(time.Now().Nanosecond())))

	for x := 0; x <= 31; x++ {
		for y := 0; y <= 31; y++ {
			r := random.Int31n(WeightSum)
			var CurrWeight int32
			for _, item := range s.TileTypes {
				CurrWeight += item.Prevalence
				if CurrWeight >= r {
					s.Map[x][y] = *NewTile(item)
					break
				}
			}
		}
	}
}

func (s *sChunk) Print() {
	for _, row := range &s.Map {
		srow := ""
		for _, tile := range row {
			srow += tile.Source + " "
		}
		fmt.Println(srow)
	}
}
