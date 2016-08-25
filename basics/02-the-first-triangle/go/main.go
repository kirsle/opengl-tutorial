package main

import (
	"runtime"
	"local/shaders"
	"github.com/go-gl/gl/v3.3-core/gl"
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

	// This seems to be needed in Go-GL so things like GenVertexArrays
	// don't crash. Put gl.Init *after* creating the window.
	err = gl.Init()
	if err != nil {
		panic(err)
	}

	// Initialize GLEW.
	window.MakeContextCurrent()

	// To draw a triangle, create a Vertex Array Object and set it as
	// the current one. This must be done after creating the window and
	// before any other OpenGL calls.
	var VertexArrayID uint32
	gl.GenVertexArrays(1, &VertexArrayID)
	gl.BindVertexArray(VertexArrayID)

	// An array of 3 vectors which represent 3 vertices for our triangle.
	g_vertex_buffer_data := []float32{
		-1.0, -1.0, 0.0,
		 1.0, -1.0, 0.0,
		 0.0,  1.0, 0.0,
	}

	////////////////////////
	// Prepare the triangle.
	////////////////////////
	var vertexbuffer uint32
	gl.GenBuffers(1, &vertexbuffer)
	gl.BindBuffer(gl.ARRAY_BUFFER, vertexbuffer)
	gl.BufferData(
		gl.ARRAY_BUFFER,
		len(g_vertex_buffer_data)*4,  // 4 bits per pixel
		gl.Ptr(g_vertex_buffer_data),
		gl.STATIC_DRAW,
	)

	// Create and compile our GLSL program from the shaders.
	ProgramID := shaders.LoadShaders("../cpp/SimpleVertexShader.glsl", "../cpp/SimpleFragmentShader.glsl")

	// Make the default clear color dark blue.
	gl.ClearColor(0.0, 0.0, 0.4, 0.0)

	// Ensure we can capture the escape key being pressed below
	window.SetInputMode(glfw.StickyKeysMode, glfw.True)
	for window.GetKey(glfw.KeyEscape) != glfw.Press && !window.ShouldClose() {
		// Clear screen and use GLSL shaders.
		gl.Clear(gl.COLOR_BUFFER_BIT)
		gl.UseProgram(ProgramID)

		// Draw the triangle.
		gl.EnableVertexAttribArray(0)
		gl.BindBuffer(gl.ARRAY_BUFFER, vertexbuffer)
		var f float32 = 0.0
		gl.VertexAttribPointer(
			0,        // attribute 0. No particular reason for 0, but must match the layout in the shader.
			3,        // size
			gl.FLOAT, // type (GL_FLOAT)
			false,    // normalized?
			0,        // stride
			gl.Ptr(&f),
		)
		gl.DrawArrays(gl.TRIANGLES, 0, 3)
		gl.DisableVertexAttribArray(0)

		// Swap buffers.
		window.SwapBuffers()
		glfw.PollEvents()
	}
}
