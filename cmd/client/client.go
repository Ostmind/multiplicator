package main

import (
	"fmt"
	"github.com/Ostmind/multiplicator/internal/client/helpers"
	"math/rand"

	"github.com/Ostmind/multiplicator/internal/client/httprequest"
)

func main() {
	n := 5 + rand.Intn(25)

	requestURL := fmt.Sprintf("http://localhost:%d/get", 64333)

	sequence := make([]float64, n)
	multSequence := make([]float64, n)

	for i := 0; i < n; i++ {
		num := 1 + rand.Float64()*(10000)
		if num > 10000 {
			sequence[i] = 10000
		} else {
			sequence[i] = num
		}

		multSequence[i] = httprequest.MakeRequest(requestURL)

	}

	fmt.Println("RTP:", helpers.GetRTP(sequence, multSequence))
}
