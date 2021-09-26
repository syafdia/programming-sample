package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func main() {
	resp, err := http.Get("http://0.0.0.0:8001/movies/tt0372784")
	if err != nil {
		log.Fatalf("can not get movie data, err: %v", err)
	}

	defer resp.Body.Close()

	var body interface{}
	err = json.NewDecoder(resp.Body).Decode(&body)
	if err != nil {
		log.Fatalf("failed when decode body, err: %v", err)
	}

	log.Println("got response from FindMultiple, response:", body)
}
