package main

import (
	"fmt"
	enums "github.com/jdobner/iv-educative-go/enums"
	"os"
)

type module string

const (
	ENUMS   = "enums"
	STRUCTS = "structs"
)

func main() {
	fmt.Println("hello world")
	switch pkg := os.Args[1]; pkg {
	case ENUMS:
		enums.Main()
	default:
		panic("nothing defined")
	}
}
