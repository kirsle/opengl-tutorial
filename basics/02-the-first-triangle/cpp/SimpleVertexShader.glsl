#version 330 core

// Declare the input data.
// "vec3" is a vector of 3 components in GLSL. It's similar (but different)
//   to the glm::vec3 we used to declare our triangle. The important thing
//   is that if we use 3 components in C++ we use 3 in the GLSL too.
// "layout(location=0)" refers to the buffer we use to feed the
//   vertexPosition_modelspace attribute. Each vertex can have numerous
//   attributes: a position, one or several colors, one or several texture
//   coordinates, lots of other things. OpenGL doesn't know what a color is,
//   it just sees a vec3. So we have to tell it which buffer corresponds to
//   which input. We do that by setting the layout to the same value as the
//   first parameter to glVertexAttribPointer. The value "0" is not important,
//   it could be 12 (but no more than glGetIntegerv(GL_MAX_VERTEX_ATTRIBS, &v))
//   the important thing is that it's the same number on both sides.
// "vertexPosition_modelspace" could have any other name. It will contain
//   the position of the vertex for each run of the vertex shader.
// "in" means that this is some input data. Soon we'll see the "out" keyword.
layout(location = 0) in vec3 vertexPosition_modelspace;

void main() {
	gl_Position.xyz = vertexPosition_modelspace;
	gl_Position.w = 1.0;
}
