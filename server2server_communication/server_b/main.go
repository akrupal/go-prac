package main

import (
	"encoding/json"
	"net/http"
)

func main() {
	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		resp := map[string]string{"message": "Hello from Server B!"}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(resp)
	})

	http.ListenAndServe(":8081", nil)
}
