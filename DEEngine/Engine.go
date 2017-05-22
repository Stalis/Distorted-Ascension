package DEEngine

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
