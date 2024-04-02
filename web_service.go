package main

import (
	"fmt"
	"net/http"
)

// Handler untuk endpoint pertama (/hello)
func helloHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "NAMA SAYA LEON BROOOO")
}

// Handler untuk endpoint kedua (/info)
func infoHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "ONE THING IS TRUE, YOU WILL BECOME A CHAMPION ONE DAY!!")
}

func main() {
    // Registrasi handler untuk masing-masing endpoint
    http.HandleFunc("/hello", helloHandler)
    http.HandleFunc("/info", infoHandler)

    // Mulai server web pada port 8080
    fmt.Println("Server berjalan di http://localhost:8080")
    http.ListenAndServe(":8080", nil)
}
