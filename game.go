package main

import (
	"fmt"
	"sync"
	"time"

	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/go-gl/glfw/v3.2/glfw"
	"runtime"
)

type Game struct {
	window gameWindow
	running bool
	initizaled bool
	test bool

	inputEvent chan interface{}
	renerInfo chan interface{}
	
	mu sync.Mutex
	wg sync.WaitGroup
}

func NewGame() *Game {
	g := &Game{}
	g.init()
	return g
}

func (g *Game) init() {
	fmt.Println("Game Init")

	g.running = true
	g.initizaled = true
}

func (g *Game) update () {
	fmt.Println("start update")
	for g.running {

		startTime := time.Now()
		fmt.Println("update")
		g.mu.Lock()
		
		g.test = !g.test

		g.mu.Unlock()
		elapsedTime := time.Since(startTime)
		if elapsedTime < (1000/60) * time.Millisecond {
			time.Sleep(1000/60 * time.Millisecond - elapsedTime)
		}
	}
}

func (g *Game) render () {
	runtime.LockOSThread()

	err := glfw.Init()
	if err != nil {panic(err)}

	g.window = *NewWindow(windowOptions{800, 600, false})

	err = gl.Init()
	if err != nil {panic(err)}
	
	for g.running {
		fmt.Println("render", g.test)

		gl.ClearColor(0.2, 0.5, 0.3, 1.0)
		gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)

		//TODO: Add to the render process
		g.window.pollEvents()
		g.window.swapBuffers()
	}
}
func (g *Game) run() {
	fmt.Println("run")

	if !g.initizaled {
		g.init()
	}

	g.wg.Add(2)
	
	go func() {
		defer g.wg.Done()
		g.render()
	}()

	go func() {
		defer g.wg.Done()
		g.update()
	}()

	g.wg.Wait()
}

func (g *Game) Close() {
	glfw.Terminate()
}