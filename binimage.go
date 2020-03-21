package main

import (
	"fmt"
	"github.com/goccy/go-yaml"
	"image"
	"image/color"
	"image/png"
	"io/ioutil"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	fmt.Println("---BINIMAGE---")
    const fileConf = "file.yaml"

	// SCRITTURA FILE CONFIGURAZIONE

	var v struct {
		Lunghezza int `yaml:"Lunghezza"`
		NomeFile string `yaml:"NomeFile"`
	}

	/*
	v.Lunghezza = 100010001
	v.NomeFile = "nomefile.txt"

	bytes, errMarshal := yaml.Marshal(v)
	check(errMarshal)

	fmt.Println(string(bytes))

	err := ioutil.WriteFile(fileConf, bytes, 0644)
	check(err)
	*/

	// LETTURA FILE CONFIGURAZIONE

	dat, err := ioutil.ReadFile(fileConf)
	check(err)

	errUnmarshal := yaml.Unmarshal([]byte(dat), &v)
	check(errUnmarshal)

	fmt.Println(v.NomeFile)

	// Create image

	divinacommedia, err := ioutil.ReadFile("divinacommedia.txt")
	check(err)
	divinacommedia_length := len(divinacommedia)

	width := 500
	height := 500

	upLeft := image.Point{0, 0}
	lowRight := image.Point{width, height}

	img := image.NewRGBA(image.Rectangle{upLeft, lowRight})

	indexData := 0

	// Set color for each pixel.
	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {

			var Red byte = 0
			var Green byte = 0
			var Blue byte = 0
			var Alpha byte = 0

			if indexData < divinacommedia_length {
				Red = divinacommedia[indexData]
			}

			indexData++

			if indexData < divinacommedia_length {
				Green = divinacommedia[indexData]
			}

			indexData++

			if indexData < divinacommedia_length {
				Blue = divinacommedia[indexData]
			}

			indexData++

			if indexData < divinacommedia_length {
				Alpha = divinacommedia[indexData]
			}

			indexData++

			img.Set(x, y, color.RGBA{Red, Green, Blue, Alpha})
		}
	}

	colore := img.RGBAAt(0,0)
	valore1 := colore.R
	valore2 := colore.G
	valore3 := colore.B
	valore4 := colore.A

	fmt.Printf("Il valore del pixel: %d %d %d %d\n", valore1, valore2, valore3, valore4)

	colore = img.RGBAAt(0,1)
	valore1 = colore.R
	valore2 = colore.G
	valore3 = colore.B
	valore4 = colore.A

	fmt.Printf("Il valore del pixel: %s %s %s %s\n", string(valore1), string(valore2), string(valore3), string(valore4))

	colore = img.RGBAAt(0,2)
	valore1 = colore.R
	valore2 = colore.G
	valore3 = colore.B
	valore4 = colore.A

	fmt.Printf("Il valore del pixel: %s %s %s %s\n", string(valore1), string(valore2), string(valore3), string(valore4))

	// Encode as PNG.
	f, _ := os.Create("image.png")
	png.Encode(f, img)

}
