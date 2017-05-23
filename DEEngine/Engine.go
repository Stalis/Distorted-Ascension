package DEEngine

func Start() {

	var tt TileTypes
	tt.GetDataFromJSON()

	var Chunk sChunk
	Chunk.GetDataFromCSV(tt)
	Chunk.Print()

	var Graphics sGraphics
	Graphics.Chunk = Chunk.GetMap()
	Graphics.start()
}
