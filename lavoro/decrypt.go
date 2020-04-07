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

func getBytes(file io.Reader) []byte {
	img, _ := png.Decode(file)

	var lista []byte

	/*
	bounds := img.Bounds()
	width, height := bounds.Max.X, bounds.Max.Y
	for y := 0; y < height; y++ {
			for x := 0; x < width; x++ {
	 */
	for y := 0; y < 1; y++ {
		for x := 0; x < 16; x++ {

			R,G,B,A := img.At(x, y).RGBA()

			fmt.Printf("R %#18b | %d", R, R)
			//R |= R >> 8
			fmt.Printf(" R %#18b | %d\n", byte(R/257), byte(R/257))

			fmt.Printf("G %#18b | %d", G, G)
			//G |= G >> 8
			fmt.Printf(" G %#18b | %d\n", byte(G/257), byte(G/257))

			fmt.Printf("B %#18b | %d", B, B)
			//B |= B >> 8
			fmt.Printf(" B %#18b | %d\n", byte(B/257), byte(B/257))

			fmt.Printf("A %#18b | %d", A, A)
			//A |= A >> 8
			fmt.Printf(" A %#18b | %d\n\n", byte(A/257), byte(A/257))


			lista = append(lista, byte(R))
			lista = append(lista, byte(G))
			lista = append(lista, byte(B))
			lista = append(lista, byte(A))
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
