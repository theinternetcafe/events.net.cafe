package main

import (
	"log"
	"net/http"
	"os"
)

func createReactHandler(path string) http.Handler {
	// Assuming React files are in the '../webgui' directory relative to where your Go code is
	fs := http.FileServer(http.Dir(path))
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		r.URL.Path = "/" + r.URL.Path
		fs.ServeHTTP(w, r)
	})
}

func main() {
	// Set address based on environment variable
	addr := os.Getenv("WEBGUI_ADDR")
	if addr == "" {
		addr = "127.0.0.1:8080"
	}

	// Set path based on environment variable
	path := os.Getenv("WEBGUI_PATH")
	if path == "" {
		path = "./webgui-app"
	}

	http.Handle("/", createReactHandler(path))
	log.Println("Starting web server on :8080...")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
