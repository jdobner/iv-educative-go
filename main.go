package main

import (
	enums "github.com/jdobner/iv-educative-go/enums"
	"github.com/jdobner/iv-educative-go/structs"
	"os"
)

type module string

const (
	ENUMS     = "enums"
	STRUCTS   = "structs"
	FUNCTIONS = "functions"
)

func main() {
	switch pkg := os.Args[1]; pkg {
	case ENUMS:
		enums.Main()
	case STRUCTS:
		structs.Main()
	case FUNCTIONS:
		runFunctions()
	default:
		panic("nothing defined")
	}
}
