package Map

//go:generate stringer -type=tCellType
type tCellType int

const (
	CELL_LAND_1 tCellType = iota
	CELL_LAND_2
	CELL_LAND_3
)
