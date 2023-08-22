package main

import (
	"log"
	"net/http"
	"os"

	"github.com/emersion/go-webdav"
)

func main() {
	// Set address based on environment variable
	addr := os.Getenv("WEBDAV_ADDR")
	if addr == "" {
		addr = "127.0.0.1:8081"
	}

	// Set path based on environment variable
	path := os.Getenv("WEBDAV_PATH")
	if path == "" {
		path = "/tmp/webdav"
	}

	// Create a handler
	handler := webdav.Handler{
		FileSystem: webdav.LocalFileSystem(path),
	}
	log.Printf("WebDAV server running on %s", addr)
	log.Fatal(http.ListenAndServe(addr, &handler))
}
