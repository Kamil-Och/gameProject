package main

import (
	"fmt"
	"sync"
	"time"
	"slices"

	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/go-gl/glfw/v3.2/glfw"
	"runtime"
)

type Game struct {
	window gameWindow
	world World

	running bool
	initizaled bool
	test bool

	inputEvent chan []glfw.Key
	renerInfo chan World
	
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
	g.world = *NewWorld()

	g.running = true
	g.initizaled = true

	g.inputEvent = make(chan []glfw.Key, 1)
	g.renerInfo = make(chan World, 1)
}

func (g *Game) update () {
	fmt.Println("start update")
	keysRec := []glfw.Key{}

	for g.running {
		startTime := time.Now()

		//Select to get the keys
		select {
			case keys := <-g.inputEvent:
				fmt.Println("keys", keys)
				keysRec = keys
			default:
		}

		//Input logic
		if slices.Contains(keysRec, glfw.KeyW) {
			fmt.Println("")
			g.running = false
		}
		//Update the world

		

		//Select to send the world to render
		select{
			case g.renerInfo <- g.world:
			default:
		}
		
		elapsedTime := time.Since(startTime)
		if elapsedTime < (1000/60) * time.Millisecond {
			time.Sleep(1000/60 * time.Millisecond - elapsedTime)
		}else {
			fmt.Println("update time over", elapsedTime)
		}
	}
	fmt.Println("update end")
}

func (g *Game) render () {
	runtime.LockOSThread()

	err := glfw.Init()
	if err != nil {panic(err)}

	g.window = *NewWindow(windowOptions{800, 600, false})

	err = gl.Init()
	if err != nil {panic(err)}

	keys := []glfw.Key{}

	for g.running {
		fmt.Println("render start")
		//Select to get the world
		select {
			case g.world = <-g.renerInfo:
			default:
		}
		//Clear the screen
		gl.ClearColor(0.2, 0.5, 0.3, 1.0)
		gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)

		//Render the world


		//Swap the buffers
		g.window.pollEvents()
		g.window.swapBuffers()
		//Get the keys
		keys = []glfw.Key{}

		if g.window.window.GetKey(glfw.KeyEscape) == glfw.Press {
			keys = append(keys, glfw.KeyEscape)
		} 
		if g.window.window.GetKey(glfw.KeyW) == glfw.Press {
			keys = append(keys, glfw.KeyW)
		}

		//Select to send the keys to update
		select {
			case g.inputEvent <- keys:
			default:
		}
	}
	fmt.Println("render end")
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