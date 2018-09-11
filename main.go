package main

import (
	"image"
	"image/color"
	"image/jpeg"
	"log"
	"math"
	"os"
	"sync"
)

func main() {
	args := os.Args[1:]
	convert(args[0])
}

//
func convert(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	img, err := jpeg.Decode(file)
	if err != nil {
		log.Fatal(err)
	}

	workers := 5
	b := img.Bounds()
	h := int(math.Ceil(float64(b.Max.Y / workers)))
	imgSet := image.NewRGBA(b)

	var wg sync.WaitGroup
	wg.Add(workers)
	for i := 0; i < workers; i++ {
		go func(i int) {
			miny, maxy := i*h, (i+1)*h
			if maxy > b.Max.Y {
				maxy = b.Max.Y
			}
			for y := miny; y < maxy; y++ {
				for x := 0; x < b.Max.X; x++ {
					_, g, _, _ := img.At(x, y).RGBA()
					imgSet.Set(x, y, color.Gray{uint8(g / 256)})
				}
			}
			wg.Done()
		}(i)
	}
	wg.Wait()

	outFile, err := os.Create("output.jpg")
	if err != nil {
		log.Fatal(err)
	}
	defer outFile.Close()
	jpeg.Encode(outFile, imgSet, nil)
}

//
func convert2(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	img, err := jpeg.Decode(file)
	if err != nil {
		log.Fatal(err)
	}

	b := img.Bounds()
	imgSet := image.NewRGBA(b)
	for y := 0; y < b.Max.Y; y++ {
		for x := 0; x < b.Max.X; x++ {
			_, g, _, _ := img.At(x, y).RGBA()
			imgSet.Set(x, y, color.Gray{uint8(g / 256)})
		}
	}

	outFile, err := os.Create("output.jpg")
	if err != nil {
		log.Fatal(err)
	}
	defer outFile.Close()
	jpeg.Encode(outFile, imgSet, nil)
}
