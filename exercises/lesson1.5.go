// Измените палитру программы lissajous так, чтобы изображение было зеленого цвета на черном фоне, чтобы быть более похожим на экран осциллографа.
// Чтобы создать веб-цвет #RRGGBB, воспользуйтесь инструкцией color . RGBA{0xRR, 0xGG,0xBB,0xff}, в которой каждая пара шестнадцатеричных цифр
// представляет яркость красного, зеленого и синего компонентов пикселя.
package main

import (
	"image"
	"image/color"
	"image/gif"
	"log"
	"math"
	"math/rand"
	"os"
	"time"
)

var palette = []color.Color{color.Black, color.RGBA{G: 255, A: 255}}

const (
	BackGroundColorIndex = 0
	FigureColorIndex     = 1
)

func main() {
	f, err := os.Create("lissajous_oscilloscope.gif")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	lissajous(f)
}

func lissajous(out *os.File) {
	const (
		cycles  = 5     // Количество полных колебаний x
		res     = 0.001 // Угловое разрешение
		size    = 100   // Канва изображения охватывает [size..+size]
		nFrames = 64    // Количество кадров в анимации
		delay   = 8     // Задержка междку кадрами(единица - 10мс)
	)
	rand.Seed(time.Now().UnixNano())
	freq := rand.Float64() * 3.0 // Относительная частота колебаний y
	anim := gif.GIF{LoopCount: nFrames}
	phase := 0.0
	for i := 0; i < nFrames; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)

		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), FigureColorIndex)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	if err := gif.EncodeAll(out, &anim); err != nil {
		log.Fatal(err)
	}

}
