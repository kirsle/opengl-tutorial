# (Go) OpenGL Tutorial

This repo is my effort to learn OpenGL 3.3, particularly in Go.

This uses the tutorial from <http://www.opengl-tutorial.org>.

For the most part the code in this repo is written in Go, with occasional
C++ versions here or there in case it was difficult to translate into Go.

As an example, in the `basics/01-opening-a-window` there is the C++
version from the tutorial and the Go version that does all of the same
things in roughly the same amount of code.

Useful links:

* The OpenGL tutorial: <http://www.opengl-tutorial.org>
* GoDoc for GL 3.3 core: <https://godoc.org/github.com/go-gl/gl/v3.3-core/gl>
* GoDoc for glfw: <https://godoc.org/github.com/go-gl/glfw/v3.2/glfw>

# Go vs C++

Some key differences between C++ and Go for OpenGL:

* In C++ most functions begin with `gl` or `glfw`, in Go these are the
  module name. For example `glClearColor` becomes `gl.ClearColor`
* When the GL bindings in Go want a pointer (for example Go gives an
  error like "cannot use &f (type \*float32) as type unsafe.Pointer in
  argument to gl.VertexAttribPointer"), use `gl.Ptr(&f)` to wrap it.
  Do not use `unsafe.Pointer` as the error message suggests.
* When loading shaders (e.g. in `02-the-first-triangle/go`) use `gl.Strs`
  to convert a Go string into a C string. The Go string must be null
  terminated. The `gl.Strs` function returns the C-string and a function
  you should call to free up the memory of the string afterward (you can
  just `defer` that call).

# Setup Steps

## Linux Distributions

### Fedora

```bash
# For C++ and compiler tools.
sudo dnf install gcc-c++ cmake make

# For OpenGL stuff
sudo dnf install libX11-devel libXi-devel libXrandr-devel glew-devel \
  glfw-devel glm-devel
```

Note: the OpenGL requirements were found by trial-and-error based on what libs
the C++ code needed to import.

### Ubuntu

(From <http://www.opengl-tutorial.org/beginners-tutorials/tutorial-1-opening-a-window/>)

```bash
# For C++ and compiler tools.
sudo apt-get install cmake make g++

# For OpenGL stuff
sudo apt-get install libx11-dev libxi-dev libgl1-mesa-dev libglu1-mesa-dev \
  libxrandr-dev libxext-dev libxi-dev
```

## Go

```bash
# OpenGL 3.3
go get -u github.com/go-gl/gl/v3.3-core/gl

# Go-GLFW only supports up to OpenGL 3.2 (the C++ tutorial used 3.3)
go get -u github.com/go-gl/glfw/v3.2/glfw
```
