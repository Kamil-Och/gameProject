package main

import (
	"fmt"
)

type Game struct {
	windowSets windowOptions
	running bool
	initizaled bool
}

func (g *Game) init() {
	fmt.Println("Game Init")
	g.windowSets.width = 300
	g.windowSets.height = g.windowSets.width/16*9
	fmt.Println("window: ", g.windowSets.width, "x", g.windowSets.height)
	g.running = true
	g.initizaled = true
}

func (g *Game) run() {
	fmt.Println("run")

	if !g.initizaled {
		g.init()
	}

	for g.running { //Game Loop
		fmt.Println("running")
	}
	
}