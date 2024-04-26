package main

import (
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func MainHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Received request: %v %v %v\n", r.Method, r.URL, 200)
	io.WriteString(w, "Hello Magican")
}

func main() {
	godotenv.Load(".dev.env")
	port := os.Getenv("PORT")

	http.HandleFunc("/", MainHandler)
	http.ListenAndServe(":"+port, nil)
}
