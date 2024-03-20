package main

import (
	"fmt"
	"runtime"

	//"github.com/go-gl/glfw/v3.2/glfw"
)

//TODO: ADD glfw window implementation and options

func init() {
	fmt.Println("init")
	runtime.LockOSThread()
}

func main() {
	fmt.Println("main")
	game := Game{}
	game.run()
}