package main

import (
    "fmt"
    "log"
    "net/http"
    "os"
    "text/template"
)

var maintenanceFlagPath = "/path/to/your/webserver/maintenance.flag" // Update with correct path

func main() {
    // Define file server for static files
    fs := http.FileServer(http.Dir("./static"))
    http.Handle("/static/", http.StripPrefix("/static/", fs))

    // Serve the homepage
    http.HandleFunc("/", serveHomePage)

    // Start the server
    fmt.Println("Starting server on :8080...")
    log.Fatal(http.ListenAndServe(":8080", nil))
}

func serveHomePage(w http.ResponseWriter, r *http.Request) {
    // Check if the maintenance flag file exists
    if _, err := os.Stat(maintenanceFlagPath); err == nil {
        // Serve the maintenance page if the flag file exists
        tmpl, err := template.ParseFiles("templates/maintenance.html")
        if err != nil {
            http.Error(w, "Maintenance Page Not Found", http.StatusNotFound)
            return
        }
        tmpl.Execute(w, nil)
        return
    }

    // Serve the regular home page if not in maintenance mode
    tmpl, err := template.ParseFiles("templates/home.html")
    if err != nil {
        http.Error(w, "Home Page Not Found", http.StatusNotFound)
        return
    }
    tmpl.Execute(w, nil)
}
