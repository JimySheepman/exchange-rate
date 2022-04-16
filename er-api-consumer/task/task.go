package task

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"er-api-consumer/model"
)

const (
	MAIN_PATH  = "https://v6.exchangerate-api.com/v6/"
	QUERY_PARH = "/latest/USD"
)

func Task() {
	var rates model.Rates

	res, err := http.Get(MAIN_PATH + os.Getenv("API_KEY") + QUERY_PARH)
	if err != nil {
		log.Fatal(err)
	}

	req, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal("Could not complete read from request body")
	}

	err = json.Unmarshal(req, &rates)
	if err != nil {
		log.Fatal("Could not complete unmarshal body")
	}

	fmt.Printf("%s", rates)
}
