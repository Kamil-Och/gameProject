package main

import (
	"fmt"
	//"runtime"
)

func init() {
	fmt.Println("init")
	//runtime.LockOSThread()
}

func main() {
	fmt.Println("main")
	game := NewGame()
	defer game.Close()

	game.run()
}