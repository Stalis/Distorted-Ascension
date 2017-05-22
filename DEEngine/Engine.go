package DEEngine

import "math/rand"
import "time"
import "fmt"

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

func Start() {
	var Chunk sChunk
	Chunk.TileSource = [3]string{"#", "@", "%"}
	Chunk.TileTypeNames = [3]string{"Land1", "Land2", "Land3"}
	Chunk.TileTypes = [3]sTileType{
		*NewTileType(Chunk.TileTypeNames[0], 50, Chunk.TileSource[0]),
		*NewTileType(Chunk.TileTypeNames[1], 50, Chunk.TileSource[1]),
		*NewTileType(Chunk.TileTypeNames[2], 50, Chunk.TileSource[2]),
	}
	Chunk.Generate()
	Chunk.Print()
}
