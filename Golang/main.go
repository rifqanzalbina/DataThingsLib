package main

import (
	"fmt"
	"log"

	. "gorgonia.org/gorgonia"
	"gorgonia.org/tensor"
)

func main() {
	g := NewGraph()

	// Membuat dua tensor
	a := tensor.New(tensor.WithShape(2, 2), tensor.WithBacking([]float32{1, 2, 3, 4}))
	b := tensor.New(tensor.WithShape(2, 2), tensor.WithBacking([]float32{5, 6, 7, 8}))

	// Membuat dua node yang mewakili tensor dalam graf
	nodeA := NodeFromAny(g, a, WithName("a"))
	nodeB := NodeFromAny(g, b, WithName("b"))

	// Melakukan operasi penambahan
	nodeC, err := Add(nodeA, nodeB)
	if err != nil {
		log.Fatal(err)
	}

	// Evaluasi nodeC
	machine := NewTapeMachine(g)
	if err := machine.RunAll(); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%v\n", nodeC.Value())
}
