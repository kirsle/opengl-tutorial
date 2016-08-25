// Include standard headers.
#include <stdio.h>
#include <stdlib.h>
#include <string.h>

// Include GLEW. Always include it before gl.h and glfw.h, since it's a bit magic.
#include <GL/glew.h>

// Include glfw
#include <GLFW/glfw3.h>

// 3D Mathematics
#include <glm/glm.hpp>
using namespace glm;

// Include our shaders loader.
#include "shaders.cpp"

int main() {
	// Initialize GLFW
	if ( !glfwInit() )
	{
		fprintf( stderr, "Failed to initialize GLFW\n" );
		return -1;
	}

	glfwWindowHint(GLFW_SAMPLES, 4); // 4x antialiasing
	glfwWindowHint(GLFW_CONTEXT_VERSION_MAJOR, 3); // We want OpenGL 3.3
	glfwWindowHint(GLFW_CONTEXT_VERSION_MINOR, 3);
	glfwWindowHint(GLFW_OPENGL_FORWARD_COMPAT, GL_TRUE); // To make MacOS happy
	glfwWindowHint(GLFW_OPENGL_PROFILE, GLFW_OPENGL_CORE_PROFILE); // We don't want the old OpenGL

	// Open a window and create its OpenGL context
	GLFWwindow* window;
	window = glfwCreateWindow( 1024, 768, "Tutorial 01", NULL, NULL);
	if (window == NULL ) {
		fprintf(stderr, "Failed to open GLFW window.");
		glfwTerminate();
		return -1;
	}
	glfwMakeContextCurrent(window); // Initialize GLEW
	glewExperimental = true; // Needed in core profile
	if (glewInit() != GLEW_OK) {
		fprintf(stderr, "Failed to initialize GLEW");
		return -1;
	}

	// To draw a triangle, create a Vertex Array Object and set it as
	// the current one. This must be done after creating the window and
	// before any other OpenGL calls.
	GLuint VertexArrayID;
	glGenVertexArrays(1, &VertexArrayID);
	glBindVertexArray(VertexArrayID);

	// An array of 3 vectors which represents 3 vertices
	static const GLfloat g_vertex_buffer_data[] = {
		-1.0f, -1.0f, 0.0f,
		 1.0f, -1.0f, 0.0f,
		 0.0f,  1.0f, 0.0f,
	};

	////////////////////////
	// Prepare the triangle.
	////////////////////////
	GLuint vertexbuffer;

	// Generate 1 buffer, put the resulting identifier in vertexbuffer
	glGenBuffers(1, &vertexbuffer);

	// The following commands will talk about our buffer.
	glBindBuffer(GL_ARRAY_BUFFER, vertexbuffer);
	glBufferData(GL_ARRAY_BUFFER, sizeof(g_vertex_buffer_data), g_vertex_buffer_data, GL_STATIC_DRAW);

	// At this point the window would immediately exit, so wait for a keypress.

	// Create and compile our GLSL program from the shaders.
	GLuint programID = LoadShaders("SimpleVertexShader.glsl", "SimpleFragmentShader.glsl");

	// Make the default clear color dark blue.
	glClearColor(0.0f, 0.0f, 0.4f, 0.0f);

	// Ensure we can capture the escape key being pressed below
	glfwSetInputMode(window, GLFW_STICKY_KEYS, GL_TRUE);
	do {
		// Clear the screen.
		glClear(GL_COLOR_BUFFER_BIT | GL_DEPTH_BUFFER_BIT);
		glUseProgram(programID);

		// Draw the triangle.
		glEnableVertexAttribArray(0);
		glBindBuffer(GL_ARRAY_BUFFER, vertexbuffer);
		glVertexAttribPointer(
			0,        // attribute 0. No particular reason for 0 but it must match the layout in the shader.
			3,        // size
			GL_FLOAT, // type
			GL_FALSE, // normalized?
			0,        // stride
			(void*)0  // array buffer offset
		);

		// This actually does the drawing.
		// Starting from vertex 0; 3 vertices total -> 1 triangle.
		glDrawArrays(GL_TRIANGLES, 0, 3);
		glDisableVertexAttribArray(0);

		// Swap buffers
		glfwSwapBuffers(window);
		glfwPollEvents();
	} // Check if the ESC key was pressed or the window was closed
	while( glfwGetKey(window, GLFW_KEY_ESCAPE ) != GLFW_PRESS &&
	glfwWindowShouldClose(window) == 0 );
}
