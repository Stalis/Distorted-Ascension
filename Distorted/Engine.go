package Distorted

func Start() {

	var tt TileTypes
	tt.GetDataFromJSON()

	var Chunk sChunk
	//Chunk.GetDataFromCSV(tt)
	Chunk.GetDataFromGenerator(1,2,3,tt)

	var Graphics sGraphics
	Graphics.LoadChunk(Chunk.GetMap())
	Graphics.start()
}
