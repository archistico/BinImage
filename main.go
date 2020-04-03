package main

import (
	"fmt"
	"github.com/alexflint/go-arg"
	"github.com/goccy/go-yaml"
	"image"
	"image/color"
	"image/png"
	"io/ioutil"
	"math"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func FormatNumber(n int64) string {
	if n < 0 {
		return "-" + FormatNumber(-n)
	}

	in := strconv.FormatInt(n, 10)
	numOfCommas := (len(in) - 1) / 3

	out := make([]byte, len(in)+numOfCommas)

	for i, j, k := len(in)-1, len(out)-1, 0; ; i, j = i-1, j-1 {
		out[j] = in[i]
		if i == 0 {
			return string(out)
		}
		if k++; k == 3 {
			j, k = j-1, 0
			out[j] = '.'
		}
	}
}

func calcNumberByte(i Format) int {
	return i.w*i.h*4
}

func calcNumberImageRequired(data int, maxbyte int) int {
	return int(math.Ceil(float64(data)/float64(maxbyte)))
}

func calcNumberByteLost(data int, maxbyte int) int {
	return (maxbyte*calcNumberImageRequired(data, maxbyte))-data
}

func choiseFormat(d int, i []Format) Format {
	// Scelgo i formati in cui ci vogliono meno immagini
	// tra quelle scelgo quelle che lasciano vuoti meno byte
	//numeroimmagini := make([]int, 7)

	for c := 0; c < len(i); c++ {
		immagini := calcNumberImageRequired(d, calcNumberByte(i[c]))
		lost := calcNumberByteLost(d, calcNumberByte(i[c]))

		i[c].immagini = immagini
		i[c].lost = lost
	}

	// fai una lista dei primi due minimi
	sort.Sort(ByNumeroImmagini(i))

	minimo := 0
	minimo2 := 0

	for c := 0; c < len(i); c++ {
		if c == 0 {
			minimo = i[c].immagini
		}
		if minimo != i[c].immagini {
			if minimo2 == 0 {
				if i[c].immagini > minimo {
					minimo2 = i[c].immagini
				}
			}
		}
	}

	// aggiungi a ris solo i formati con immagini == a minimo e minimo2
	ris := []Format {}
	for c := 0; c < len(i); c++ {
		if i[c].immagini == minimo || i[c].immagini == minimo2 {
			ris = append(ris, i[c])
		}
	}

	// ordino in base ai byte lost
	sort.Sort(ByLostByte(ris))

	return ris[0]
}

func EncodeImage(conf FileConf, form Format) {
	divinacommedia, err := ioutil.ReadFile(conf.NomeFile)
	check(err)
	divinacommediaLength := len(divinacommedia)

	upLeft := image.Point{0, 0}
	lowRight := image.Point{form.w, form.h}

	img := image.NewRGBA(image.Rectangle{upLeft, lowRight})

	indexData := 0

	// Set color for each pixel.
	for y := 0; y < form.h; y++ {
		for x := 0; x < form.w; x++ {

			var Red byte = 0
			var Green byte = 0
			var Blue byte = 0
			var Alpha byte = 0

			if indexData < divinacommediaLength {
				Red = divinacommedia[indexData]
			}

			indexData++

			if indexData < divinacommediaLength {
				Green = divinacommedia[indexData]
			}

			indexData++

			if indexData < divinacommediaLength {
				Blue = divinacommedia[indexData]
			}

			indexData++

			if indexData < divinacommediaLength {
				Alpha = divinacommedia[indexData]
			}

			indexData++

			img.Set(x, y, color.RGBA{Red, Green, Blue, Alpha})
		}
	}

	f, _ := os.Create(conf.NomeImmagine+conf.EstensioneImmagine)
	png.Encode(f, img)
}


// ------------------- TYPE ----------------------

type Format struct {
	name string
	w int
	h int
	immagini int
	lost int
}

type ByNumeroImmagini []Format
type ByLostByte []Format

func (a ByNumeroImmagini) Len() int           { return len(a) }
func (a ByNumeroImmagini) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByNumeroImmagini) Less(i, j int) bool { return a[i].immagini < a[j].immagini }

func (a ByLostByte) Len() int           { return len(a) }
func (a ByLostByte) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByLostByte) Less(i, j int) bool { return a[i].lost < a[j].lost }

type FileConf struct {
	NomeFile string `yaml:"NomeFile"`
	DataLength int `yaml:"DataLength"`
	NomeImmagine string `yaml:"NomeImmagine"`
	EstensioneImmagine  string `yaml:"EstensioneImmagine"`
	NumeroImmagini int `yaml:"NumeroImmagini"`
}

func main() {

	var args struct {
		NomeFile string `arg:"positional, required" help:"Nome del file da convertire"`
	}
	arg.MustParse(&args)

	NomeFile:=args.NomeFile
	file, err := os.Open(NomeFile)
	check(err)
	f, err := file.Stat()
	check(err)
	dataLength := int(f.Size())

	// ---------- SCELTA FORMATO ---------------

	var images = []Format{
		{"qvga", 320, 240, 0,0},
		{"vga", 640, 480, 0,0},
		{"svga",800, 600, 0,0},
		{"xga",1024, 768, 0,0},
		{"hd720",1280, 720, 0,0},
		{"hd1080",1920, 1080, 0,0},
		{"wqhd",2560, 1140, 0,0},
	}

	fmt.Printf("Data: %s [byte]\n", FormatNumber(int64(dataLength)))

	formatoImmagini := choiseFormat(dataLength, images)
	maxByteInFormat := calcNumberByte(formatoImmagini)
	formatoImmagini.immagini = calcNumberImageRequired(dataLength, maxByteInFormat)
	formatoImmagini.lost = calcNumberByteLost(dataLength, maxByteInFormat)
	fmt.Printf("Choise: %s | w:%d | h:%d | images:%d | lost:%s\n", formatoImmagini.name, formatoImmagini.w, formatoImmagini.h, formatoImmagini.immagini, FormatNumber(int64(formatoImmagini.lost)))

	// ---------- SCRITTURA FILE CONFIGURAZIONE ---------------

	NomeConf := strings.TrimSuffix(NomeFile, filepath.Ext(NomeFile))
	FileConfName := NomeConf + ".yaml"
	NomeImmagine := NomeConf
	EstensioneImmagine := ".png"

	conf := FileConf{
		NomeFile:   NomeFile,
		DataLength: dataLength,
		NomeImmagine: NomeImmagine,
		EstensioneImmagine: EstensioneImmagine,
		NumeroImmagini: formatoImmagini.immagini,
	}

	bytes, errMarshal := yaml.Marshal(conf)
	check(errMarshal)
	err = ioutil.WriteFile(FileConfName, bytes, 0644)
	check(err)

	// ---------- CODIFICA ---------------
	EncodeImage(conf, formatoImmagini)

	// ---------------------------------------
	fmt.Print("-- Conversione completata --")
}
