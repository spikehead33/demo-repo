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
		log.Print("got / requests\n")
		io.WriteString(w, fmt.Sprintf("Hello, %v", os.Getenv("TARGET")))
	})

	err := http.ListenAndServe(":3000", mux)
	if err != nil {
		panic(err)
	}
}
