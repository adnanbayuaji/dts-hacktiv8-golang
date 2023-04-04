package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"time"
)

func main() {
	for true {
		kirimData(rand.Intn(100), rand.Intn(100))
		time.Sleep(15 * time.Second)
	}
}

type Cuaca struct {
	Water int `json:"water"`
	Wind  int `json:"wind"`
}

func kirimData(water int, wind int) {
	data := map[string]interface{}{
		"water": water,
		"wind":  wind,
	}

	requestJson, err := json.Marshal(data)
	if err != nil {
		log.Fatalln(err)
	}
	// responseBody := bytes.NewBuffer(requestJson)
	client := &http.Client{}
	if err != nil {
		log.Fatalln(err)
	}

	req, err := http.NewRequest("POST", "https://jsonplaceholder.typicode.com/posts", bytes.NewBuffer(requestJson))
	req.Close = true
	req.Header.Set("Content-type", "application/json")
	if err != nil {
		log.Fatalln(err)
	}

	res, err := client.Do(req)
	if err != nil {
		log.Fatalln(err)
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatalln(err)
	}

	log.Println(string(body))

	// Unmarshal result
	datas := Cuaca{}
	err = json.Unmarshal(body, &datas)
	if err != nil {
		log.Printf("Reading body failed: %s", err)
		return
	}

	var statWat string = ""
	if datas.Water <= 5 {
		statWat = "aman"
	} else if datas.Water >= 6 && datas.Water <= 8 {
		statWat = "siaga"
	} else if datas.Water > 8 {
		statWat = "bahaya"
	}

	var statWin string = ""
	if datas.Wind <= 6 {
		statWin = "aman"
	} else if datas.Wind >= 7 && datas.Wind <= 15 {
		statWin = "siaga"
	} else if datas.Wind > 15 {
		statWin = "bahaya"
	}

	fmt.Printf("status water : %s\n", statWat)
	fmt.Printf("status wind : %s\n", statWin)

}
