package main

import (
	"flag"
	"io"
	"log"
	"os"

	shapepipe "github.com/lindell/shape-pipe/pkg/shape-pipe"
	"github.com/lindell/shape-pipe/pkg/shapes"
)

var availableShapes = map[string]shapepipe.Shape{
	"tree": shapes.Tree,
}

func main() {
	shapeName := flag.String("shape", "tree", "defines the shape to be used")
	flag.Parse()

	shape, ok := availableShapes[*shapeName]
	if !ok {
		log.Fatalf("could not find shape: %s", *shapeName)
	}

	reader := &shapepipe.ShapeReader{
		Shape:  shape,
		Reader: os.Stdin,
	}

	buf := make([]byte, 1)
	_, err := io.CopyBuffer(os.Stdout, reader, buf)
	if err != nil {
		log.Fatal(err)
	}
}
