package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"server_communication/request"
)

var httpClient = &http.Client{Timeout: 2 * time.Second}
var reqBuilder = request.NewHttpRequestBuilder(httpClient)

func main() {
	http.HandleFunc("/call-server-b", func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		var response map[string]string
		err := reqBuilder.NewRequestWithContext(ctx).
			ResponseAs(&response).
			Post("http://localhost:8081/hello")

		if err != nil {
			http.Error(w, "Failed to call server B: "+err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{
			"fromServerA": "Successfully received response",
			"serverBsaid": response["message"],
		})
	})

	log.Println("Server A running on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
