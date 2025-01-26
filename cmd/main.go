package main

import (
    "log"
    "net/http"

    "github.com/rakyll/statik/fs" // Import the statik package
    _ "portfolio/statik"         // Import the statik package (generated)
)

func main() {
    // Create a statik file system
    statikFS, err := fs.New()
    if err != nil {
        log.Fatalf("Failed to initialize statik file system: %v", err)
    }

    // Serve static files from the embedded file system
    http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(statikFS)))

    // Default route for the home page
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        http.ServeFile(w, r, "internal/templates/home.html")
    })

    // Start the HTTP server
    log.Println("Server running at http://localhost:8080")
    err = http.ListenAndServe(":8080", nil)
    if err != nil {
        log.Fatalf("Server failed to start: %v", err)
    }
}


