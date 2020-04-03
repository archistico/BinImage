package main

import (
	"fmt"
	"math"
)

func suddividiBlocchi(dati string, max int) []string {
	datiLength := len(dati)
	blocchi := int(math.Ceil(float64(datiLength)/float64(max)))

	res := []string{}

	for index:=0; index<blocchi; index++ {

		if (max*(index+1))<datiLength {
			res = append(res, dati[max*index:max*(index+1)])
		} else {
			res = append(res, dati[max*index:])
		}
	}

	return res
}

func main() {
	// Divisione blocchi di codice
	dati := "12345678901234567890123456789012345"
	maxImmagine := 10

	arr := suddividiBlocchi(dati, maxImmagine)
	fmt.Println(arr)
}
