package main

import (
	"fmt"
	"runtime"

	//"github.com/go-gl/glfw/v3.2/glfw"
)

//TODO: ADD glfw window implementation and options
//TODO: ADD concurrency to the game loop
//TODO: UPDATE[ logic 60 fps ] and RENDER [ rendering to screen unlimited ][ buffer? ] functions in game

func init() {
	fmt.Println("init")
	runtime.LockOSThread()
}

func main() {
	fmt.Println("main")
	game := Game{}
	game.run()
}