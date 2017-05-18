package Map

type Cell struct {
	Type tCellType
}

func (c *Cell) Init(arg tCellType) (out *Cell) {
	out.Type = arg
	return
}
