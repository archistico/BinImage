package main

import (
	"fmt"
	"math"
)

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

		//fmt.Println(dati[max*index:])
	}

	return res
}

func main() {
	// Divisione blocchi di codice
	dati := []byte("12345678901234567890123456789012345")
	maxImmagine := 10
	fmt.Println("Dati: ", dati)

	arr := suddividiBlocchi(dati, maxImmagine)
	fmt.Println("Suddivisione: ", arr)
}
