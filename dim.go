package main

import (
	"fmt"
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

		fmt.Printf("formato %s \tnum:%d \tlost:%s\n", i[c].name, immagini, FormatNumber(int64(lost)))
	}

	// fai una lista dei primi due minimi
	sort.Sort(ByNumeroImmagini(i))
	fmt.Println(i)

	// calcolo per ogni formato il numero di byte vuoti
	//performance := make([]int, 7)

	return i[0]
}

type Format struct {
	name string
	w int
	h int
	immagini int
	lost int
}

type ByNumeroImmagini []Format
type ByLossByte []Format

func (a ByNumeroImmagini) Len() int           { return len(a) }
func (a ByNumeroImmagini) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByNumeroImmagini) Less(i, j int) bool { return a[i].immagini < a[j].immagini }

func main() {
	var images = []Format{
		{"qvga", 320, 240, 0,0},
		{"vga", 640, 480, 0,0},
		{"svga",800, 600, 0,0},
		{"xga",1024, 768, 0,0},
		{"hd720",1280, 720, 0,0},
		{"hd1080",1920, 1080, 0,0},
		{"wqhd",2560, 1140, 0,0},
	}

	dataLength :=2000000

	format := choiseFormat(dataLength, images)
	fmt.Printf("Format choise: %s %d %d\n", format.name, format.w, format.h)
}
