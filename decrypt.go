package main

import (
	"crypto/sha1"
	"fmt"
	"github.com/alexflint/go-arg"
	"github.com/goccy/go-yaml"
	"image"
	"image/png"
	"io"
	"io/ioutil"
	"log"
	"os"
	"strings"
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

	bounds := img.Bounds()
	width, height := bounds.Max.X, bounds.Max.Y
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {

			R,G,B,_ := img.At(x, y).RGBA()

			lista = append(lista, byte(R/257))
			lista = append(lista, byte(G/257))
			lista = append(lista, byte(B/257))
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

	// ----------- CONTROLLO HASH --------------

	file, err := os.Open(fileConfYaml.NomeFile)
	defer file.Close()

	h := sha1.New()
	if _, err := io.Copy(h, file); err != nil {
		log.Fatal(err)
	}

	hash := fmt.Sprintf("%x", h.Sum(nil))
	if strings.Compare(hash, fileConfYaml.Sha1)==0 {
		fmt.Println("HASH FILE ESPORTATO OK")
	} else {
		fmt.Println("ATTENZIONE! HASH FILE ESPORTATO NON CORRISPONDE")
	}
}
