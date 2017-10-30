package main

import (
	"encoding/json"
	"flag"
	"log"
	"net/http"

	"./utils"
)

func main() {
	var url string
	flag.StringVar(&url, "url", "", "取得対象URL")
	flag.Parse()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		result, err := utils.GetContent(url)
		if err != nil {
			log.Fatal(err)
		}
		json.NewEncoder(w).Encode(result)
	})
	http.ListenAndServe(":8080", nil)
}
