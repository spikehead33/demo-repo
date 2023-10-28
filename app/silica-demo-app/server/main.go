package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		targetVar := os.Getenv("TARGET")

		log.Print(fmt.Sprintf("got / requests: Print %v\n", targetVar))
		io.WriteString(w, fmt.Sprintf("Hello, %v", targetVar))
	})

	err := http.ListenAndServe(":3000", mux)
	if err != nil {
		panic(err)
	}
}
