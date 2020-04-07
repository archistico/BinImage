package main

import (
	"fmt"
	"github.com/alexflint/go-arg"
	"github.com/goccy/go-yaml"
	"image"
	"image/png"
	"io"
	"io/ioutil"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

type FileConf struct {
	NomeFile string `yaml:"NomeFile"`
	DataLength int `yaml:"DataLength"`
	NomeImmagine string `yaml:"NomeImmagine"`
	EstensioneImmagine  string `yaml:"EstensioneImmagine"`
	LarghezzaImmagine int `yaml:"LarghezzaImmagine"`
	AltezzaImmagine int `yaml:"AltezzaImmagine"`
	NumeroImmagini int `yaml:"NumeroImmagini"`
	Sha1 string `yaml:"Sha1"`
}

// Get the bi-dimensional pixel array
func getPixels(file io.Reader) ([][]Pixel, error) {
	img, _, err := image.Decode(file)

	if err != nil {
		return nil, err
	}

	bounds := img.Bounds()
	width, height := bounds.Max.X, bounds.Max.Y

	var pixels [][]Pixel
	for y := 0; y < height; y++ {
		var row []Pixel
		for x := 0; x < width; x++ {
			row = append(row, rgbaToPixel(img.At(x, y).RGBA()))
		}
		pixels = append(pixels, row)
	}

	return pixels, nil
}

// img.At(x, y).RGBA() returns four uint32 values; we want a Pixel
func rgbaToPixel(r uint32, g uint32, b uint32, a uint32) Pixel {
	return Pixel{int(r / 257), int(g / 257), int(b / 257), int(a / 257)}
}

// Pixel struct example
type Pixel struct {
	R int
	G int
	B int
	A int
}

func main() {

	// VARIABILI
	var fileConfYaml FileConf

	// LETTURA DEGLI ARGOMENTI DA CONSOLE
	var args struct {
		NomeFileYaml string `arg:"positional, required" help:"Nome del file yaml"`
	}
	arg.MustParse(&args)

	// LETTURA FILE CONFIGURAZIONE
	dat, err := ioutil.ReadFile(args.NomeFileYaml)
	check(err)
	errUnmarshal := yaml.Unmarshal([]byte(dat), &fileConfYaml)
	check(errUnmarshal)

	// VARIABILI
	// var imgs []image.Image
	image.RegisterFormat("png", "png", png.Decode, png.DecodeConfig)

	// CREAZIONE NOMI IMMAGINI DA LEGGERE
	for c:=0; c<fileConfYaml.NumeroImmagini ; c++ {
		nf := fmt.Sprintf("%s-%03d.%s", fileConfYaml.NomeImmagine, c, fileConfYaml.EstensioneImmagine)
		fmt.Println(nf)

		file, err := os.Open(nf)

		if err != nil {
			fmt.Println("Error: File could not be opened")
			os.Exit(1)
		}

		defer file.Close()

		pixels, err := getPixels(file)

		if err != nil {
			fmt.Println("Error: Image could not be decoded")
			os.Exit(1)
		}

		fmt.Println(pixels)




	}



	//// PER OGNI IMMAGINE LEGGI TUTTO
	//indexData := 0
	//var byteLetti []byte
	//
	//for y := 0; y < fileConfYaml.AltezzaImmagine; y++ {
	//	for x := 0; x < fileConfYaml.LarghezzaImmagine; x++ {
	//
	//		colore := imgs[0].RGBAAt(x,y)
	//
	//		if indexData < fileConfYaml.DataLength {
	//			byteLetti = append(byteLetti, colore.R)
	//		}
	//
	//		indexData++
	//
	//		if indexData < fileConfYaml.DataLength {
	//			byteLetti = append(byteLetti, colore.G)
	//		}
	//
	//		indexData++
	//
	//		if indexData < fileConfYaml.DataLength {
	//			byteLetti = append(byteLetti, colore.B)
	//		}
	//
	//		indexData++
	//
	//		if indexData < fileConfYaml.DataLength {
	//			byteLetti = append(byteLetti, colore.A)
	//		}
	//
	//		indexData++
	//	}
	//}

}
