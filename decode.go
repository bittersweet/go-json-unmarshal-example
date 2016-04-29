package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
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

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal("ReadAll", err)
	}

	p := map[string]string{}
	err = json.Unmarshal(body, &p)
	fmt.Printf("%#v\n", p)
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
