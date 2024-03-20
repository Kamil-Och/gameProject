package main

import (
	"fmt"
	"sync"
	"time"
)

type Game struct {
	windowSets windowOptions
	running bool
	initizaled bool
	test bool
	mu sync.Mutex
}

func (g *Game) init() {
	fmt.Println("Game Init")
	g.windowSets.width = 300
	g.windowSets.height = g.windowSets.width/16*9
	fmt.Println("window: ", g.windowSets.width, "x", g.windowSets.height)
	g.running = true
	g.initizaled = true
}

func (g *Game) update () {
	for g.running {
		fmt.Println("update")
		g.mu.Lock()
		g.test = !g.test
		g.mu.Unlock()
		time.Sleep(5 * time.Second)
	}
}

func (g *Game) render () {
	
	fmt.Println("render")

	for g.running {
		
		fmt.Println("render", g.test)
		time.Sleep(1 * time.Second)
	}
}

func (g *Game) run() {
	fmt.Println("run")

	if !g.initizaled {
		g.init()
	}

	go g.update()

	go g.render()

	for g.running {}
}