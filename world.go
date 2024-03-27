package main

type World struct {
	x int
	y int
}

func NewWorld() *World {
	w := &World{}
	w.init()
	return w
}

func (w *World) init() {
	w.x = 0
	w.y = 0
}
