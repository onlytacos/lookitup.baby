package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Printf("HONKS ONLINE ON PORT %v", port)
	http.Handle("/kubecon", http.RedirectHandler("https://www.youtube.com/watch?v=ZY4kTDCsJ44", 302))
	http.Handle("/", http.RedirectHandler("https://twitter.com/IanColdwater/status/1292895288546545666", 302))
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", port), nil))
}
