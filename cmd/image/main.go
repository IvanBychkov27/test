// Lissajous генерирует анимированный GIF из случайных фигур Лиссажу.
//  $ go build imag.go
//  $ ./main > out.gif

package main

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"log"
	"math"
	"math/rand"
	"os"
	"time"
)

var palette = []color.Color{color.White, color.RGBA{255, 0, 0, 255}}

const (
	//whiteIndex = 0
	blackIndex = 1
)

func main() {

	lissajous(os.Stdout)

}

func lissajous(out io.Writer) {
	const (
		cycles  = 5     // Количество полных колебаний x
		res     = 0.001 // Угловое разрешение
		size    = 200   // Канва изображения охватывает [size..+size]
		nframes = 64    // Количество кадров анимации
		delay   = 10    // Задержка между кадрами (единица - 10мс)
	)
	rand.Seed(time.Now().UTC().UnixNano())
	freq := rand.Float64() * 3.0 // Относительная частота колебаний у
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0 // Разность фаз
	//colorIndex := color.RGBA{255, 20, 20, 255}
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), blackIndex)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	err := gif.EncodeAll(out, &anim)
	if err != nil {
		log.Println(err.Error())
	}
}
