package main

import (
	"fmt"
	"github.com/alexflint/go-arg"
	"math"
	"sort"
	"strconv"
)

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

func main() {

	var args struct {
		NomeFile string `arg:"-n" help:"il nome del file da convertire"`
		NumeroByte int  `arg:"-b" help:"numero di byte"`
	}
	arg.MustParse(&args)

	dataLength :=int(args.NumeroByte)

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

	format := choiseFormat(dataLength, images)
	maxByteInFormat := calcNumberByte(format)
	format.immagini = calcNumberImageRequired(dataLength, maxByteInFormat)
	format.lost = calcNumberByteLost(dataLength, maxByteInFormat)
	fmt.Printf("Choise: %s | w:%d | h:%d | images:%d | lost:%s\n", format.name, format.w, format.h, format.immagini, FormatNumber(int64(format.lost)))
}
