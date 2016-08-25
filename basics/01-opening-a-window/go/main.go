package main

import (
	"runtime"
	"github.com/go-gl/glfw/v3.2/glfw"
)

func init() {
	// This is needed to arrange that main() runs on the main thread.
	// See documentation for functions that are only allowed to be called from
	// the main thread.
	runtime.LockOSThread()
}

func main() {
	err := glfw.Init()
	if err != nil {
		panic(err)
	}
	defer glfw.Terminate()

	glfw.WindowHint(glfw.Samples, 4) // 4x antialiasing
	glfw.WindowHint(glfw.ContextVersionMajor, 3) // We want OpenGL 3.3
	glfw.WindowHint(glfw.ContextVersionMinor, 3)
	glfw.WindowHint(glfw.OpenGLForwardCompatible, glfw.True) // To make MacOS happy
	glfw.WindowHint(glfw.OpenGLProfile, glfw.OpenGLCoreProfile) // We don't want the old OpenGL

	// Open a window and create its OpenGL context.
	window, err := glfw.CreateWindow(1024, 768, "Tutorial 01", nil , nil)
	if err != nil {
		panic(err)
	}

	// Initialize GLEW.
	window.MakeContextCurrent()

	// Ensure we can capture the escape key being pressed below
	window.SetInputMode(glfw.StickyKeysMode, glfw.True)
	for window.GetKey(glfw.KeyEscape) != glfw.Press && !window.ShouldClose() {
		// Draw nothing. That comes in tutorial 2!

		// Swap buffers.
		window.SwapBuffers()
		glfw.PollEvents()
	}
}
