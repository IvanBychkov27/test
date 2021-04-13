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
	"os"
	"time"
)

//--------------------------------------------  цвет фона  -------------------------  цвет линии -------------
var palette = []color.Color{color.RGBA{0, 0, 0, 255}, color.RGBA{0, 255, 0, 255}} // цвет фона и линии

const blackIndex = 1

func main() {
	fN := "fig01.gif"
	fileName := "cmd/image2/a/" + fN
	buildGifFile(fileName)
}

func buildGifFile(fileName string) {
	file, err := os.Create(fileName)
	if err != nil {
		log.Println(err.Error())
	} else {
		log.Printf("Создан файл %s\n", fileName)
	}
	defer file.Close()

	gifFigure(file)
	//lissajous(file)
}

const (
	cycles  = 100   // Количество полных колебаний x
	res     = 0.001 // Угловое разрешение
	sizeX   = 400   // Канва изображения охватывает [size..+size]
	sizeY   = 300   // Канва изображения охватывает [size..+size]
	nframes = 74    // Количество кадров анимации
	delay   = 6     // Задержка между кадрами (единица - 10мс)
)

func gifFigure(out io.Writer) {
	rand.Seed(time.Now().UTC().UnixNano())
	var (
		//xn, yn, xk, yk, x1, y1, x2, y2 float64 // прямоугольник
		r, xr, yr float64 // круг
	)
	anim := gif.GIF{LoopCount: nframes}
	rect := image.Rect(0, 0, sizeX+1, sizeY+1)

	c1 := Circle{
		Radius: 20,          // радиус круга
		X:      100,         // координата x центра круга
		Y:      100,         // координата y центра круга
		fill:   true,        // закраска
		move:   1 * math.Pi, // направление движения - [0:2]*math.Pi : 0 - вправо; 0,5 - вниз ; 1 - влево; 1,5 - вверх; 2 - вправо
	}

	for step := 0; step < nframes; step++ {
		img := image.NewPaletted(rect, palette)

		// ---- Круг 1 -------------------
		r, xr, yr = 20, 200, 200
		xr, yr = c1.coordCircle(3.18)
		img = animCircle(img, step, r, xr, yr, true)
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)

		//img = animLine(img, 0, 0, 0, sizeY)
		//anim.Delay = append(anim.Delay, delay)
		//anim.Image = append(anim.Image, img)
		//
		//img = animLine(img, 0, 0, sizeX, 0)
		//anim.Delay = append(anim.Delay, delay)
		//anim.Image = append(anim.Image, img)

		//img = animLine(img, 0, sizeY, sizeX, sizeY)
		//anim.Delay = append(anim.Delay, delay)
		//anim.Image = append(anim.Image, img)
		//
		//img = animLine(img, sizeX, 0, sizeX, sizeY)
		//anim.Delay = append(anim.Delay, delay)
		//anim.Image = append(anim.Image, img)

		// ---- Круг 2 -------------------
		//r, xr, yr = 20, 200, -50
		//yr += float64(step) * 2 * math.Pi
		//img = animCircle(img, step, r, xr, yr, true)
		//anim.Delay = append(anim.Delay, delay)
		//anim.Image = append(anim.Image, img)

		// ---- Прямоугольник -------------------
		//x1, y1, x2, y2 = 400, 175, 450, 225
		//x1 -= float64(step) * 2 * math.Pi
		//x2 -= float64(step) * 2 * math.Pi
		//xn, yn, xk, yk = x1, y1, x1, y2 // 1
		//img = animLine(img, xn, yn, xk, yk)
		//anim.Delay = append(anim.Delay, delay)
		//anim.Image = append(anim.Image, img)

		//xn, yn, xk, yk = x1, y2, x2, y2 // 2
		//img = animLine(img, xn, yn, xk, yk)
		//anim.Delay = append(anim.Delay, delay)
		//anim.Image = append(anim.Image, img)
		//
		//xn, yn, xk, yk = x2, y2, x2, y1 // 3
		//img = animLine(img, xn, yn, xk, yk)
		//anim.Delay = append(anim.Delay, delay)
		//anim.Image = append(anim.Image, img)
		//
		//xn, yn, xk, yk = x2, y1, x1, y1 // 4
		//img = animLine(img, xn, yn, xk, yk)
		//anim.Delay = append(anim.Delay, delay)
		//anim.Image = append(anim.Image, img)
	}
	err := gif.EncodeAll(out, &anim)
	if err != nil {
		log.Println(err.Error())
	}
}

func animLine(img *image.Paletted, x1, y1, x2, y2 float64) *image.Paletted {
	var x, y float64
	lenLine := math.Sqrt((x2-x1)*(x2-x1) + (y2-y1)*(y2-y1))
	for step := 0.0; step < lenLine; step += 0.01 {
		//x, y = line(step, x1, y1, x2, y2)
		x = x1 + step*(x2-x1)/lenLine
		y = y1 + step*(y2-y1)/lenLine
		img.SetColorIndex(int(x), int(y), blackIndex)
	}
	return img
}

func line(step float64, x1, y1, x2, y2 float64) (x, y float64) {
	var dx, dy, lenLine float64
	lenLine = math.Sqrt((x2-x1)*(x2-x1) + (y2-y1)*(y2-y1))
	dx = (x2 - x1) / lenLine
	dy = (y2 - y1) / lenLine

	x = x1 + dx*step
	y = y1 + dy*step
	return x, y
}

type Circle struct {
	Radius float64
	X      float64
	Y      float64
	fill   bool
	move   float64 // [0:2] * math.Pi : 0 - вправо; 0,5 - вниз; 1 - влево; 1,5 - вверх; 2 - вправо
}

func (c *Circle) coordCircle(speed float64) (float64, float64) {
	change := true
	speed *= math.Pi
	direction := math.Pi * 1

	dx := math.Cos(c.move)
	dy := math.Sin(c.move)

	x := c.X + dx*speed
	y := c.Y + dy*speed

	if x+c.Radius > sizeX {
		c.move += direction
		change = false
		fmt.Println("moveX+ = ", c.move/math.Pi)
	}
	if x-c.Radius < 0 {
		c.move -= direction
		change = false
		fmt.Println("moveX- = ", c.move/math.Pi)
	}
	if y+c.Radius > sizeY && change {
		c.move += direction
		fmt.Println("moveY+ = ", c.move/math.Pi)
	}
	if y-c.Radius < 0 && change {
		c.move -= direction
		fmt.Println("moveY- = ", c.move/math.Pi)
	}

	c.X, c.Y = x, y
	return x, y
}

func animCircle(img *image.Paletted, frame int, radius, xr, yr float64, fill bool) *image.Paletted {
	var r float64
	r = radius
	for step := 0.0; step < radius; step += 0.001 {
		x, y := circle(step, r)
		x += float64(frame) * math.Pi * 0.02
		img.SetColorIndex(int(xr+x), int(yr+y), blackIndex)

		if fill { // заливка круга цветом
			r += 0.5
			if r > radius {
				r = 0.5
			}
		}
	}
	return img
}

func circle(step float64, r float64) (x, y float64) {
	x = math.Sin(step) * r
	y = math.Cos(step) * r
	return x, y
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
