package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/h2non/bimg"
)

func main() {
	fmt.Printf("Running test: %s\n", os.Args[1])

	switch os.Args[1] {
	case "resize":
		resize()
	default:
		panic("Unknown test")
	}
}

func resize() {
	iterations, err := strconv.Atoi(os.Args[2])
	if err != nil {
		panic(err)
	}

	for i := 0; i < iterations; i++ {
		fmt.Printf("Iteration %d\n", i)
		resizeStep()
	}
}

func resizeStep() {
	buffer, err := bimg.Read("images/image50.jpg")
	if err != nil {
		panic(err)
	}

	newImage, err := bimg.NewImage(buffer).Resize(5000, 3333)
	if err != nil {
		panic(err)
	}

	bimg.Write("out/resize.jpg", newImage)
}
