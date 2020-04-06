package main

import (
	"bufio"
	"crypto/sha1"
	"fmt"
	"github.com/alexflint/go-arg"
	"github.com/goccy/go-yaml"
	"image"
	"image/color"
	"image/png"
	"io"
	"io/ioutil"
	"log"
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

func suddividiBlocchi(dati []byte, max int) [][]byte {
	datiLength := len(dati)
	blocchi := int(math.Ceil(float64(datiLength)/float64(max)))

	res := [][]byte{}

	for index:=0; index<blocchi; index++ {
		in:=max*index
		fi:=max*(index+1)
		if (fi)<datiLength {
			res = append(res, dati[in:fi])
		} else {
			res = append(res, dati[in:])
		}
	}

	return res
}

func EncodeImage(dati []byte, nomeFile string, width int, height int) {
	datiLength := len(dati)

	upLeft := image.Point{0, 0}
	lowRight := image.Point{width, height}

	img := image.NewRGBA(image.Rectangle{upLeft, lowRight})

	indexData := 0

	// Set color for each pixel.
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {

			var Red byte = 0
			var Green byte = 0
			var Blue byte = 0
			var Alpha byte = 0

			if indexData < datiLength {
				Red = dati[indexData]
			}

			indexData++

			if indexData < datiLength {
				Green = dati[indexData]
			}

			indexData++

			if indexData < datiLength {
				Blue = dati[indexData]
			}

			indexData++

			if indexData < datiLength {
				Alpha = dati[indexData]
			}

			indexData++

			img.Set(x, y, color.RGBA{Red, Green, Blue, Alpha})
		}
	}

	f, _ := os.Create(nomeFile)
	png.Encode(f, img)
}

func ConsoleFormat(f []Format) int {
	for index:=0; index<len(f); index++ {
		fmt.Printf("%d) %7s | w:%4d | h:%4d\n", index, f[index].name, f[index].w, f[index].h)
	}

	fmt.Printf("Scegli il formato desiderato: ")
	reader := bufio.NewReader(os.Stdin)
	char, _, err := reader.ReadRune()

	if err != nil {
		fmt.Println(err)
	}

	res:=0
	switch char {
		case '1':
			res = 1
		case '2':
			res =  2
		case '3':
			res =  3
		case '4':
			res =  4
		case '5':
			res =  5
		case '6':
			res =  6
	}
	return res
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
	LarghezzaImmagine int `yaml:"LarghezzaImmagine"`
	AltezzaImmagine int `yaml:"AltezzaImmagine"`
	NumeroImmagini int `yaml:"NumeroImmagini"`
	Sha1 string `yaml:"Sha1"`
}

// ----------------------------------------
// -------------- MAIN --------------------
// ----------------------------------------

func main() {

	var args struct {
		NomeFile string `arg:"positional, required" help:"Nome del file da convertire"`
		Formato bool `arg:"-f" help:"Seleziona il formato di output. Altrimenti sceglie in automatico"`
	}
	arg.MustParse(&args)

	// ---------------- TITOLO -------------------

	titolo:= `|---------------------------------------------------------|
|  _____ _         _            ___            _          |
| |   __| |_ ___ _| |___    ___|  _|   ___ ___| |___ ___  |
| |__   |   | .'| . | -_|  | . |  _|  |  _| . | | . |  _| |
| |_____|_|_|__,|___|___|  |___|_|    |___|___|_|___|_|   |
|---------------------------------------------------------|
|                                    by Emilie Rollandin  |`

	fmt.Println(titolo)

	NomeFile:=args.NomeFile
	file, err := os.Open(NomeFile)
	defer file.Close()
	check(err)
	f, err := file.Stat()
	check(err)
	dataLength := int(f.Size())

	// ----------- CREAZIONE HASH --------------

	h := sha1.New()
	if _, err := io.Copy(h, file); err != nil {
		log.Fatal(err)
	}

	hash := fmt.Sprintf("%x", h.Sum(nil))

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
	fmt.Println("|                                                         |")
	fmt.Printf("|  DIMENSIONE FILE INPUT: %23s [byte]  |\n", FormatNumber(int64(dataLength)))

	// FORZATURA DEL FORMATO
	var formatoImmagini Format
	if args.Formato {
		indexFormato:=ConsoleFormat(images)
		formatoImmagini = images[indexFormato]
	} else {
		formatoImmagini = choiseFormat(dataLength, images)
	}

	// CALCOLO DATI PRINCIPALI

	maxByteInFormat := calcNumberByte(formatoImmagini)
	formatoImmagini.immagini = calcNumberImageRequired(dataLength, maxByteInFormat)
	formatoImmagini.lost = calcNumberByteLost(dataLength, maxByteInFormat)
	fmt.Println("|---------------------------------------------------------|")
	fmt.Printf("|  FORMATO IMMAGINE: %6s | %4d x %4d | IMMAGINI:%4d |\n", strings.ToUpper(formatoImmagini.name), formatoImmagini.w, formatoImmagini.h, formatoImmagini.immagini)

	// ---------- SCRITTURA FILE CONFIGURAZIONE ---------------

	NomeConf := strings.TrimSuffix(NomeFile, filepath.Ext(NomeFile))
	FileConfName := NomeConf + ".yaml"
	NomeImmagine := fmt.Sprintf("%s-%dx%d", NomeConf, formatoImmagini.w, formatoImmagini.h)
	EstensioneImmagine := "png"

	conf := FileConf{
		NomeFile:   NomeFile,
		DataLength: dataLength,
		NomeImmagine: NomeImmagine,
		EstensioneImmagine: EstensioneImmagine,
		LarghezzaImmagine: formatoImmagini.w,
		AltezzaImmagine: formatoImmagini.h,
		NumeroImmagini: formatoImmagini.immagini,
		Sha1: hash,
	}

	bytes, errMarshal := yaml.Marshal(conf)
	check(errMarshal)
	err = ioutil.WriteFile(FileConfName, bytes, 0644)
	check(err)

	testofileconf:= `|---------------------------------------------------------|
|  FILE DI CONFIGURAZIONE YAML ESPORTATO                  |`
	fmt.Print(testofileconf)
	fmt.Println()

	// ---------- DIVIDO I DATI IN ENTRATA ---------------
	// creo un array di byte per ogni immagine
	dati, err := ioutil.ReadFile(conf.NomeFile)
	check(err)
	blocchi := suddividiBlocchi(dati, maxByteInFormat)

	// ---------- CODIFICA ---------------
	// mando solo una sequenza di byte
	for c:=0; c<len(blocchi); c++ {
		n:=fmt.Sprintf("%s-%03d.%s", conf.NomeImmagine, c, conf.EstensioneImmagine)
		EncodeImage(blocchi[c], n, formatoImmagini.w, formatoImmagini.h)
	}

	// ----------- CHIUSURA --------------
	chiusura:= `|---------------------------------------------------------|
|  FILE IMMAGINI ESPORTATI                                |
|---------------------------------------------------------|`
	fmt.Print(chiusura)
}

