package main

import (
	"fmt"
	"sync"
	"time"
	
	_ "github.com/go-gl/glfw/v3.2/glfw"
)

type Game struct {
	window gameWindow
	running bool
	initizaled bool
	test bool
	mu sync.Mutex
}

func NewGame() *Game {
	g := &Game{}
	g.init()
	return g
}

func (g *Game) init() {
	fmt.Println("Game Init")
	
	g.window = *NewWindow(windowOptions{800, 600, false})
	
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
	// err := glfw.Init()
	// if err != nil {panic(err)}
	// defer glfw.Terminate()

	go g.update()

	go g.render()

	for g.running {
		if g.window.shouldClose() {
			g.running = false
		}

		g.window.swapBuffers()
		g.window.pollEvents()
	}
}