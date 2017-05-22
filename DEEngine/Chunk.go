package DEEngine

import (
	"bytes"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
)

func errcheck(err error) {
	if err != nil {
		log.Println("Error:", err)
		return
	}
}

type sTileType struct {
	Name       string `json:"name,omitempty"`
	Prevalence int32  `json:"prevalence,omitempty"`
	Source     string `json:"source,omitempty"`
}

type sTile struct {
	sTileType
}

func NewTile(Type sTileType) *sTile {
	return &sTile{Type}
}

type sChunk struct {
	Map [32][32]sTile
}

func (s sChunk) GetMap() [32][32]sTile {
	return s.Map
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

func (s *sChunk) GetDataFromCSV(types TileTypes) {
	file, err := ioutil.ReadFile("map.csv")
	defer errcheck(err)
	r := csv.NewReader(bytes.NewReader(file))
	grid, err := r.ReadAll()
	for x, row := range grid {
		for y, sindex := range row {
			index, err := strconv.Atoi(sindex)
			s.Map[x][y] = *NewTile(types.types[index-1])
			errcheck(err)
		}
	}
}

type TileTypes struct {
	types []sTileType `json:"types,omitempty"`
}

func (t *TileTypes) GetDataFromJSON() {
	file, err := ioutil.ReadFile("tiles.json")
	err = json.Unmarshal(file, &t.types)
	defer errcheck(err)
}

func (t TileTypes) GetTypes() []sTileType {
	return t.types
}
