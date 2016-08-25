package shaders

import (
	"fmt"
	"io/ioutil"
	"github.com/go-gl/gl/v3.3-core/gl"
)

func LoadShaders(VertexFilePath, FragmentFilePath string) uint32 {
	// Create the shaders.
	var VertexShaderID uint32 = gl.CreateShader(gl.VERTEX_SHADER)
	var FragmentShaderID uint32 = gl.CreateShader(gl.FRAGMENT_SHADER)

	// Read the vertex shader code from the file.
	VertexShaderCode, err := ioutil.ReadFile(VertexFilePath)
	if err != nil {
		panic(err)
	}

	// Read the fragment shader code from the file.
	FragmentShaderCode, err := ioutil.ReadFile(FragmentFilePath)
	if err != nil {
		panic(err)
	}

	var Result int32 = gl.FALSE
	var InfoLogLength int32

	// Compile the vertex shader.
	fmt.Printf("Compiling vertex shader\n")
	glvsrc, vfree := gl.Strs(string(VertexShaderCode)+"\x00")
	gl.ShaderSource(VertexShaderID, 1, glvsrc, nil)
	gl.CompileShader(VertexShaderID)
	defer vfree()
	defer gl.DeleteShader(VertexShaderID)

	// Check vertex shader.
	gl.GetShaderiv(VertexShaderID, gl.COMPILE_STATUS, &Result)
	gl.GetShaderiv(VertexShaderID, gl.INFO_LOG_LENGTH, &InfoLogLength)
	if InfoLogLength > 0 {
		var VertexShaderErrorMessage uint8
		gl.GetShaderInfoLog(VertexShaderID, InfoLogLength, nil, &VertexShaderErrorMessage)
		fmt.Printf("Error: %s (compile status: %s)\n", VertexShaderErrorMessage, Result)
	}

	// Compile the fragment shader.
	fmt.Printf("Compiling fragment shader\n")
	glfsrc, ffree := gl.Strs(string(FragmentShaderCode)+"\x00")
	gl.ShaderSource(FragmentShaderID, 1, glfsrc, nil)
	gl.CompileShader(FragmentShaderID)
	defer ffree()
	defer gl.DeleteShader(FragmentShaderID)

	// Check fragment shader.
	gl.GetShaderiv(FragmentShaderID, gl.COMPILE_STATUS, &Result)
	gl.GetShaderiv(FragmentShaderID, gl.INFO_LOG_LENGTH, &InfoLogLength)
	if InfoLogLength > 0 {
		var FragmentShaderErrorMessage uint8
		gl.GetShaderInfoLog(FragmentShaderID, InfoLogLength, nil, &FragmentShaderErrorMessage)
		fmt.Printf("Error: %s\n", FragmentShaderErrorMessage)
	}

	// Link the program
	fmt.Printf("Linking program\n")
	ProgramID := gl.CreateProgram()
	gl.AttachShader(ProgramID, VertexShaderID)
	gl.AttachShader(ProgramID, FragmentShaderID)
	gl.LinkProgram(ProgramID)
	defer gl.DetachShader(ProgramID, VertexShaderID)
	defer gl.DetachShader(ProgramID, FragmentShaderID)

	// Check the program.
	gl.GetProgramiv(ProgramID, gl.LINK_STATUS, &Result)
	gl.GetProgramiv(ProgramID, gl.INFO_LOG_LENGTH, &InfoLogLength)
	if InfoLogLength > 0 {
		var ProgramErrorMessage uint8
		gl.GetProgramInfoLog(ProgramID, InfoLogLength, nil, &ProgramErrorMessage)
		fmt.Printf("Error: %s\n", ProgramErrorMessage)
	}

	return ProgramID
}
