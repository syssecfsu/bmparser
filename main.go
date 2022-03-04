package main

import (
	"flag"
	"fmt"
	"image"
	"log"
	"os"

	_ "golang.org/x/image/bmp"
)

func main() {
	log.SetFlags(0)

	var outfile string

	flag.StringVar(&outfile, "o", "raw.dib", "specify output file")
	flag.Parse()

	args := flag.Args()

	if len(args) != 1 {
		log.Fatalln("bmparser -o output_file input_file")
	}

	fp, err := os.Open(args[0])

	if err != nil {
		log.Fatalln("Failed to open file", err)
	}

	defer fp.Close()

	img, format, err := image.Decode(fp)

	if err != nil {
		log.Fatalln("Failed to decode file", err)
	}

	bounds := img.Bounds()
	fmt.Println(format, bounds)

	rgba := image.NewRGBA(bounds)

	for i := bounds.Min.X; i < bounds.Max.X; i++ {
		for j := bounds.Min.Y; j < bounds.Max.Y; j++ {
			color := img.At(i, j)
			rgba.Set(i, j, color)
		}
	}

	err = os.WriteFile(outfile, rgba.Pix, 0644)

	if err != nil {
		log.Fatalln("Failed to write file", err)
	}
}
