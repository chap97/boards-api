package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type MessageRequest struct {
	Body   string
	Sender string
}

func main() {
	http.HandleFunc("/Message", func(rw http.ResponseWriter, r *http.Request) {
		fmt.Println(r.Method)
		rw.Header().Set("Access-Control-Allow-Origin", "*")
		rw.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		rw.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
		if r.Method == "POST" {
			body, _ := io.ReadAll(r.Body)
			r.Body.Close()

			var msg MessageRequest
			err := json.Unmarshal(body, &msg)
			if err != nil {
				http.Error(rw, err.Error(), http.StatusInternalServerError)
			} else {
				fmt.Println(msg)
			}
		}
	})
	log.Fatal(http.ListenAndServe(":7286", nil))
}
