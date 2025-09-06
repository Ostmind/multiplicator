package httprequest

import (
	"encoding/json"
	"log"
	"net/http"
)

type response struct {
	Number float64 `json:"result"`
}

func MakeRequest(URL string) float64 {
	var resp response
	res, err := http.Get(URL)
	if err != nil {
		log.Fatal("cannot get multiplier for numbers")
	}

	err = json.NewDecoder(res.Body).Decode(&resp)
	if err != nil {
		log.Fatal("cannot get multiplier for numbers")
	}

	return resp.Number
}
