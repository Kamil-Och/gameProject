package main

type World struct {
	x float64
	y float64
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

type GameObject struct {
	x float64
	y float64
	//verticies?
	//texture?
	//shader?
	//etc
}