package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Payload struct {
	Icon  string `json:"icon.png"`
	Pass  string `json:"pass.json"`
	Strip string `json:"strip.png"`
}

func handleReceive(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	p := map[string]string{}
	err := json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		log.Fatal("NewDecoder", err)
	}
	for key, value := range p {
		fmt.Printf("key: %s value: %s\n", key, value)
	}

	// Build response object
	resp := map[string]string{
		"status": "ok",
	}

	output, err := json.MarshalIndent(resp, "", "  ")
	if err != nil {
		log.Fatal("MarshalIndent", err)
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Write(output)
}

func main() {
	http.HandleFunc("/receive", handleReceive)

	fmt.Println("Listening on port 8000")
	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		log.Fatal("ListenAndServe", err)
	}
}
