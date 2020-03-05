// Mandelbrot emits a PNG image of the Mandelbrot fractal.
package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"io"
	"math/cmplx"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(write http.ResponseWriter, req *http.Request) {
		write.Header().Set("Content-Type", "image/png")
		printFX(write)
	})
	http.ListenAndServe("localhost:6060", nil)
}

func printFX(write io.Writer) {
	const (
		xmin, ymin, xmax, ymax = -2, -2, +2, +2
		width, height          = 1024, 1024
	)
	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		y := float64(py)/height*(ymax-ymin) + ymin
		for px := 0; px < width; px++ {
			x := float64(px)/width*(xmax-xmin) + xmin
			z := complex(x, y) // Image point (px, py) represents complex value z.
			img.Set(px, py, mandelbrot(z))
		}
	}
	png.Encode(write, img)
	// NOTE: ignoring errors
}
func mandelbrot(z complex128) color.Color {
	const iterations = 200
	const contrast = 15
	var v complex128
	for n := uint8(0); n < iterations; n++ {
		v = v*v + z
		if cmplx.Abs(v) > 2 {
			return color.Gray{200 - contrast*n}
		}
	}
	return color.RGBA{0x11, 0xF8, 0x0F, 3}
}

func funcName1() {
	var x complex128 = 1 + 2i
	var y complex128 = complex(3, 4)
	fmt.Println(x * y)
	fmt.Println(real(x * y))
	fmt.Println(imag(x * y))
}
