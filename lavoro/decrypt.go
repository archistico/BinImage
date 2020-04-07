package main

import (
	"fmt"
	"github.com/alexflint/go-arg"
	"github.com/goccy/go-yaml"
	"image"
	"image/png"
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
	upLeft := image.Point{0, 0}
	lowRight := image.Point{fileConfYaml.LarghezzaImmagine, fileConfYaml.AltezzaImmagine}

	// CREAZIONE NOMI IMMAGINI DA LEGGERE
	for c:=0; c<fileConfYaml.NumeroImmagini ; c++ {
		nf := fmt.Sprintf("%s-%03d.%s\n", fileConfYaml.NomeImmagine, c, fileConfYaml.EstensioneImmagine)
		fmt.Println(nf)

		f, _ := os.Open(nf)
		img := image.NewRGBA(image.Rectangle{upLeft, lowRight})
		png.Encode(f, img)
		fmt.Println(img.RGBAAt(0,0))
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
