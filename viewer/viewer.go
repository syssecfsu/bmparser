package main

import (
	"flag"
	"image"
	"log"
	"os"

	g "github.com/AllenDang/giu"
)

const MAX_SZ = 8192

var rgba image.RGBA

func loop() {
	g.SingleWindow().Layout(
		g.Align(g.AlignCenter).To(
			g.ImageWithRgba(&rgba).Size(float32(rgba.Rect.Max.X), float32(rgba.Rect.Max.Y)),
		),
	)
}

func main() {
	log.SetFlags(0)

	var width, height int

	cmd := "dib_viewer -w <width> -h <height> <dib_file>"

	flag.IntVar(&width, "w", -1, "specify width of the image")
	flag.IntVar(&height, "h", -1, "specify height of the image")

	flag.Parse()

	if (width <= 0) || (width > MAX_SZ) || (height <= 0) || (height > MAX_SZ) {
		log.Fatalln(cmd)
	}

	args := flag.Args()

	if len(args) != 1 {
		log.Fatalln(cmd)
	}

	fc, err := os.ReadFile(args[0])

	if err != nil {
		log.Fatalln("Failed to read file", err)
	}

	rgba.Pix = fc
	rgba.Stride = 4 * width
	rgba.Rect.Min = image.Point{X: 0, Y: 0}
	rgba.Rect.Max = image.Point{X: width, Y: height}

	wnd := g.NewMasterWindow("Load Dib", width+32, height+32, g.MasterWindowFlagsNotResizable)
	wnd.Run(loop)
}
