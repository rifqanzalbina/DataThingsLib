package main

import (
	"fmt"
	"log"

	"gorgonia.org/gorgonia"
	"gorgonia.org/tensor"
)

func main2() {
	g := gorgonia.NewGraph()

	// Membuat tensor
	data := tensor.New(tensor.WithShape(4), tensor.WithBacking([]float32{1.0, -1.0, 0.0, 3.0}))
	node := gorgonia.NodeFromAny(g, data)

	// Menggunakan fungsi aktivasi sigmoid
	sigmoidResult, err := gorgonia.Sigmoid(node)
	if err != nil {
		log.Fatal(err)
	}

	// Membuat mesin dan menjalankannya
	vm := gorgonia.NewTapeMachine(g)
	if err := vm.RunAll(); err != nil {
		log.Fatal(err)
	}

	// Menampilkan hasil
	fmt.Println("The Result is : ")
	fmt.Println(sigmoidResult.Value().Data())
}
