package main

type WeightedEdge interface {
	From() int
	To() int
	Weight() int
}

type Edge struct {
	F, T, W int
}

func (e Edge) From() int {
	return e.F
}

func (e Edge) To() int {
	return e.T
}

func (e Edge) Weight() int {
	return e.W
}
