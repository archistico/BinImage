package main

import (
	"fmt"
	"math"
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
		numImage := calcNumberImageRequired(d, calcNumberByte(i[c]))
		numBlost := calcNumberByteLost(d, calcNumberByte(i[c]))
		fmt.Printf("formato %s \tnum:%d \tlost:%s\n", i[c].name, numImage, FormatNumber(int64(numBlost)))
	}

	// calcolo per ogni formato il numero di byte vuoti
	//performance := make([]int, 7)

	return i[0]
}

type Format struct {
	name string
	w int
	h int
}

func main() {
	var images = []Format{
		{"qvga", 320, 240},
		{"vga", 640, 480},
		{"svga",800, 600},
		{"xga",1024, 768},
		{"hd720",1280, 720},
		{"hd1080",1920, 1080},
		{"wqhd",2560, 1140},
	}

	dataLength :=2000000

	format := choiseFormat(dataLength, images)
	fmt.Printf("Format choise: %s %d %d\n", format.name, format.w, format.h)
}
