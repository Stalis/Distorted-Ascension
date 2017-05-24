package DEEngine

func Start() {

	var tt TileTypes
	tt.GetDataFromJSON()

	var Chunk sChunk
	Chunk.GetDataFromCSV(tt)
	Chunk.Print()

	var Graphics sGraphics
	Graphics.LoadChunk(Chunk.GetMap())
	Graphics.start()
}
