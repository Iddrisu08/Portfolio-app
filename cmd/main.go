package main

import (
    "io"
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
    http.Handle("/static/", http.StripPrefix("/static/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        log.Printf("Serving static file: %s", r.URL.Path)
        http.FileServer(statikFS).ServeHTTP(w, r)
    })))

    // Default route for the home page
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        f, err := statikFS.Open("/templates/home.html")
        if err != nil {
            log.Printf("Error loading home page: %v", err)
            http.Error(w, "Error loading home page", http.StatusInternalServerError)
            return
        }
        defer f.Close()

        log.Println("Serving home page")

        // Stream the file content to the response writer
        _, err = io.Copy(w, f)
        if err != nil {
            log.Printf("Error serving home page: %v", err)
        }
    })

    // Start the HTTP server
    log.Println("Server running at http://localhost:8080")
    err = http.ListenAndServe(":8080", nil)
    if err != nil {
        log.Fatalf("Server failed to start: %v", err)
    }
}




