package main

import (
	"fmt"

	"github.com/go-gl/glfw/v3.2/glfw"
)

type windowOptions struct {
	width  int
	height int
	resizable bool
}

type gameWindow struct {
	window *glfw.Window
	width int
	height int
}

func NewWindow(options windowOptions) *gameWindow {
	w := &gameWindow{}
	w.init(options)
	return w
}

func (w *gameWindow) init(options windowOptions) {
	fmt.Println("GameWindow Init")
	w.width = options.width
	w.height = options.height

	if !options.resizable {
		glfw.WindowHint(glfw.Resizable, glfw.False)
	} else {
		glfw.WindowHint(glfw.Resizable, glfw.True)
	}

	glfw.WindowHint(glfw.ContextVersionMajor, 4)
	glfw.WindowHint(glfw.ContextVersionMinor, 1)
	glfw.WindowHint(glfw.OpenGLProfile, glfw.OpenGLCoreProfile)
	glfw.WindowHint(glfw.OpenGLForwardCompatible, glfw.True)

	w.window, _ = glfw.CreateWindow(w.width, w.height, "Testing", nil, nil)
	w.window.MakeContextCurrent()
}

func (w *gameWindow) shouldClose() bool {
	return w.window.ShouldClose()
}

func (w *gameWindow) swapBuffers() {
	w.window.SwapBuffers()
}

func (w *gameWindow) pollEvents() {
	glfw.PollEvents()
}

