package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/davidbyttow/govips/v2/vips"
)

func main() {
	// vips.Startup(&vips.Config{
	// 	ConcurrencyLevel: 1,
	// 	MaxCacheSize:     0,
	// 	MaxCacheMem:      0,
	// 	MaxCacheFiles:    0,
	// })

	fmt.Printf("Running test: %s", os.Args[1])

	switch os.Args[1] {
	case "resize":
		resize()
	default:
		panic("Unknown test")
	}

	// for i := 0; i < 10; i++ {
	// 	fmt.Printf("Run step %d", i)
	// 	err := resize()
	// 	if err != nil {
	// 		panic(err)
	// 	}
	// }

	// var stats vips.MemoryStats
	// vips.ReadVipsMemStats(&stats)
	// fmt.Printf("memory stats: %w", stats)

	// vips.PrintObjectReport("Objec report")

	// vips.PrintCache()

	// vips.Shutdown()

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
	src, err := vips.NewImageFromFile("images/image50.jpg")
	if err != nil {
		panic(err)
	}
	defer src.Close()

	dst, _, err := src.ExportNative()
	if err != nil {
		panic(err)
	}

	err = os.WriteFile("out/resize.jpg", dst, 0644)
	if err != nil {
		panic(err)
	}
}

func thumbnail() error {
	return nil
}

func thumbnailStep() error {

	src, err := vips.NewThumbnailFromFile("images/image50.jpg", 5000, 3333, vips.InterestingNone)
	if err != nil {
		return err
	}
	defer src.Close()

	dst, _, err := src.ExportNative()
	if err != nil {
		return err
	}

	err = os.WriteFile("out/thumbnail.jpg", dst, 0644)
	if err != nil {
		return err
	}

	fmt.Printf("got %d bytes\n", len(dst))

	return nil
}
