// Пришлось немного изменить код из книги, потому что .gif вообще не работал
// Не смог увидеть результат, решил создать файл и внести генерацию фигуры Лиссажу в него
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

var palette = []color.Color{color.White, color.Black}

const (
	whiteIndex = 0
	blackIndex = 1
)

func main() {
	f, err := os.Create("lissajous.gif")
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
		nframes = 64    // Количество кадров в анимации
		delay   = 8     // Задержка междку кадрами(единица - 10мс)
	)
	rand.Seed(time.Now().UnixNano())
	freq := rand.Float64() * 3.0 // Относительная частота колебаний y
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(
				size+int(x*size+0.5),
				size+int(y*size+0.5),
				blackIndex,
			)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	if err := gif.EncodeAll(out, &anim); err != nil {
		log.Fatal(err)
	}

}
