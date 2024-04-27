package main

import (
	"encoding/json"
	"fmt"
	"math/rand/v2"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
)

type RandomData struct {
	Value1 int    `json:"value1"`
	Value2 string `json:"value2"`
}

func generateRandomData() RandomData {
	return RandomData{
		Value1: rand.IntN(100),
		Value2: fmt.Sprintf("random-string-%d", rand.IntN(50)),
	}
}

func eventsHandler(w http.ResponseWriter, r *http.Request) {
	// Inside your eventsHandler function:
	fmt.Println("EVENTS HANDLER")
	requiredAPIKey := "SSE_API_KEY" // Replace with your preferred method of storing the key
	clientAPIKey := r.Header.Get("Authorization")  
	if requiredAPIKey != clientAPIKey {
		http.Error(w, "Invalid API Key", http.StatusUnauthorized)
		return
	}
	flusher, ok := w.(http.Flusher)
	if !ok {
		http.Error(w, "Streaming not supported", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")

	for {
		data := generateRandomData()
		jsonData, _ := json.Marshal(data)
		fmt.Println("SOO3 ", string(jsonData))
		fmt.Fprintf(w, "data: %s\n\n", jsonData)
		flusher.Flush()
		time.Sleep(2 * time.Second)
	}
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/events", eventsHandler)
	port := os.Getenv("PORT") 
    if port == "" {
        port = "8080"
    }
	fmt.Println("Server listening at 8080")
	http.ListenAndServe(":"+port, r)
}
