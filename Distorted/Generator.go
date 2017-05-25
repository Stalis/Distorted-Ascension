package Distorted

import (
	"math/rand"
	"time"
	"fmt"
)

type Area struct {
	BaseTile sTile
}

type Generator struct {
	MainChunk sChunk
	Tile sTile
	Surroundings Area
	Environment Area
}

func NewGenerator(Gen sTile, Surr sTile, Env sTile) *Generator {
	return &Generator{
		Tile: Gen,
		Surroundings: Area{BaseTile:Surr},
		Environment: Area{BaseTile:Env},
	}
}

func (g *Generator) Start() {
	sX, sY := g.getObjectCoords()
	mX, mY := g.getChunkDims()
	fmt.Printf("sX=%d sY=%d \t mX=%d mY=%d\n", sX, sY, mX, mY)
	random := rand.New(rand.NewSource(int64(time.Now().Nanosecond())))
	r := (random.Intn(mX) + random.Intn(mY))/2
	for x, row := range g.MainChunk.Map {
		for y, _ := range row {
			if (x > sX - r && x < sX + r) && (y > sY -r && y < sY + r){
				g.MainChunk.Map[x][y] = g.Surroundings.BaseTile
			} else {
				g.MainChunk.Map[x][y] = g.Environment.BaseTile
			}
		}
	}
	if g.Tile.Name != ""{
		fmt.Println("Yep")
		g.MainChunk.Map[sX][sY] = g.Tile
	} else {
		g.MainChunk.Map[sX][sY] = g.Surroundings.BaseTile
	}
}

func (g Generator) getChunkDims() (int, int){
	return len(g.MainChunk.Map), len(g.MainChunk.Map[0])
}

func (g Generator) getObjectCoords() (int, int){
	return GaussGen(len(g.MainChunk.Map)), GaussGen(len(g.MainChunk.Map[0]))
}

func GaussGen(length int) int {

	weights := make([]int, length)
	var weightSum int

	for i := 0; i < length/2; i++ {
		if i == 0 {
			weights[i] = 1
			weights[len(weights)-1] = 1
		} else {
			weights[i] = weights[i-1] + i + 1
			weights[length-1-i] = weights[i]
		}
	}
	if length%2 == 1 {
		weights[length/2+1] = weights[length/2] + (weights[length/2] + 1)
	}
	for _, v := range weights {
		weightSum += v
	}

	for _, v := range weights {
		weightSum += v
	}
	random := rand.New(rand.NewSource(int64(time.Now().Nanosecond())))
	r := random.Intn(weightSum)

	currWeight := 0
	for i, w := range weights {
		if currWeight >= r {
			return i
		} else {
			currWeight += w
		}
	}
	return length - 1
}