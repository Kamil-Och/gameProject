package main

import (
	"fmt"
	"runtime"
)

//TODO: ADD glfw window implementation and options

func init() {
	fmt.Println("init")
	runtime.LockOSThread()
}

func main() {
	fmt.Println("main")
	game := NewGame()

	game.run()
}