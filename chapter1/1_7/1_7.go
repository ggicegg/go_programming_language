/**
1-7的示例代码，主要有web请求和处理
*/
package main

import (
	"fmt"
	"image"
	"image/color"
	"image/gif"
	"io"
	"log"
	"math"
	"math/rand"
	"net/http"
	"strconv"
)

func main() {

	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		url := request.URL
		writer.Write([]byte(url.Path))
	})

	http.HandleFunc("/count", func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte("hello"))
	})

	http.HandleFunc("/paint", func(writer http.ResponseWriter, request *http.Request) {

		err := request.ParseForm()
		if err != nil {
			log.Fatal(err)
		}
		cycles := request.Form["cycles"]
		if cycles != nil && len(cycles) > 0 {
			i, err := strconv.Atoi(cycles[0])
			if err != nil {
				log.Fatal(err)
			}
			lissajous(writer, i)
		}

	})

	http.HandleFunc("/print_head", handler)

	err := http.ListenAndServe("localhost:8080", nil)
	fmt.Println(err)

}

/**
打印gif
*/
func lissajous(out io.Writer, cycl int) {
	var palette = []color.Color{color.White, color.Black, color.RGBA{0xAB, 0xBC, 0xCF, 3}}

	const (
		whiteIndex = 0 // first color in palette
		blackIndex = 1 // next color in palette
		greenIndex = 2
	)
	//cycles  := cycl     // number of complete x oscillator revolutions

	const (
		res     = 0.001 // angular resolution
		size    = 100   // image canvas covers [-size..+size]
		nframes = 64    // number of animation frames
		delay   = 8     // delay between frames in 10ms units
	)
	freq := rand.Float64() * 3.0 // relative frequency of y oscillator
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0 // phase difference
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < float64(cycl)*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), greenIndex)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim) // NOTE: ignoring encoding errors
}

// handler echoes the HTTP request.
func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%s %s %s\n", r.Method, r.URL, r.Proto)
	for k, v := range r.Header {
		fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
	}
	fmt.Fprintf(w, "Host = %q\n", r.Host)
	fmt.Fprintf(w, "RemoteAddr = %q\n", r.RemoteAddr)
	if err := r.ParseForm(); err != nil {
		log.Print(err)
	}
	for k, v := range r.Form {
		fmt.Fprintf(w, "Form[%q] = %q\n", k, v)
	}
}
