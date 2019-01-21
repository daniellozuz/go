package main

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"math"
	"math/rand"
	"os"
)

var palette = []color.Color{
	color.White,
	color.RGBA{0x00, 0xf0, 0x00, 0xff},
	color.RGBA{0x00, 0xd0, 0x00, 0xff},
	color.RGBA{0x00, 0xb0, 0x00, 0xff},
	color.RGBA{0x00, 0x90, 0x00, 0xff},
	color.RGBA{0x00, 0x70, 0x00, 0xff},
	color.RGBA{0x00, 0x50, 0x00, 0xff},
	color.RGBA{0x00, 0x30, 0x00, 0xff},
	color.RGBA{0x00, 0x10, 0x00, 0xff},
	color.RGBA{0x00, 0x10, 0x00, 0xff},
	color.RGBA{0x00, 0x30, 0x00, 0xff},
	color.RGBA{0x00, 0x50, 0x00, 0xff},
	color.RGBA{0x00, 0x70, 0x00, 0xff},
	color.RGBA{0x00, 0x90, 0x00, 0xff},
	color.RGBA{0x00, 0xb0, 0x00, 0xff},
	color.RGBA{0x00, 0xd0, 0x00, 0xff},
	color.RGBA{0x00, 0xf0, 0x00, 0xff},
}

func main() {
	lissajous(os.Stdout)
}

func lissajous(out io.Writer) {
	const (
		cycles  = 5     // number of complete x oscillator revolutions
		res     = 0.001 // angular resolution
		size    = 100   // image canvas covers [-size..+size]
		nframes = 64    // number of animation frames
		delay   = 8     // delay between frames in 10ms units
	)
	freq := rand.Float64() * 3.0 // relative frequency of y oscillator
	anim := gif.GIF{LoopCount: nframes}
	phaseDifference := 0.0
	paletteIndex := 0
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phaseDifference)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), uint8(1+paletteIndex))
		}
		phaseDifference += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
		paletteIndex = (paletteIndex + 1) % 16
	}
	gif.EncodeAll(out, &anim)
}
