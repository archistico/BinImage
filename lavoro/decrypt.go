package main

import (
	"fmt"
	"github.com/alexflint/go-arg"
	"github.com/goccy/go-yaml"
	"image"
	"image/png"
	"io"
	"io/ioutil"
	"math"
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

func getBytes(file io.Reader) []byte {
	img, _, _ := image.Decode(file)

	var lista []byte

	/*
	bounds := img.Bounds()
	width, height := bounds.Max.X, bounds.Max.Y
	for y := 0; y < height; y++ {
			for x := 0; x < width; x++ {
	 */
	bounds := img.Bounds()
	width, height := bounds.Max.X, bounds.Max.Y
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			R,G,B,A := img.At(x, y).RGBA()
			lista = append(lista, byte(math.Round(float64(R)/257)))
			lista = append(lista, byte(math.Round(float64(G)/257)))
			lista = append(lista, byte(math.Round(float64(B)/257)))
			lista = append(lista, byte(math.Round(float64(A)/257)))

		}
	}
	return lista
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
	var bLetti []byte
	image.RegisterFormat("png", "png", png.Decode, png.DecodeConfig)

	// CREAZIONE NOMI IMMAGINI DA LEGGERE
	for c:=0; c<fileConfYaml.NumeroImmagini ; c++ {
		nf := fmt.Sprintf("%s-%03d.%s", fileConfYaml.NomeImmagine, c, fileConfYaml.EstensioneImmagine)
		file, err := os.Open(nf)
		if err != nil {
			fmt.Println("Error: File could not be opened")
			os.Exit(1)
		}

		defer file.Close()
		bLetti = append(bLetti, getBytes(file)...)
	}
	bScrivere := bLetti[:fileConfYaml.DataLength]
	err = ioutil.WriteFile(fileConfYaml.NomeFile, bScrivere, 0644)
	check(err)
}
