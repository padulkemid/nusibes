package main

import (
	"encoding/json"
	"log"
	"net/http"
)

const PORT = ":8080"

func main() {
	http.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		p := map[string]interface{}{
			"data": "pong!",
		}

		j, _ := json.Marshal(p)

		w.Header().Set("Content-Type", "application/json")
		w.Write(j)
	})

	log.Fatal(http.ListenAndServe(PORT, nil))
}
