package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

var DEBUG bool = true

func debug(to_print ...interface{}) {
	if DEBUG {
		fmt.Println(to_print...)
	}
}

type HelloworldResponse struct {
	Message string
}

func helloworldHandler(w http.ResponseWriter, r *http.Request) {

	// response as struct
	response := HelloworldResponse{"Hello world!"}

	// create json response from struct
	response_json, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	// send response
	_, err = w.Write(response_json)
	if err != nil {
		fmt.Println("Unable to write response")
	}
}

func main() {
	if len(os.Args) < 1 || os.Args[1] == "-h" || os.Args[1] == "--help" {
		fmt.Println("Usage: devserver.exe <port>")
		return
	}

	port := os.Args[1]
	http.HandleFunc("/helloworld", helloworldHandler)
	fmt.Println("HTTP server listening on port " + port)

	log.Fatal(http.ListenAndServe(":"+port, nil))
}
