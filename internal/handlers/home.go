package handlers

import (
    "html/template"
    "net/http"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
    // Load templates
    tmpl, err := template.ParseFiles("internal/templates/layout.html", "internal/templates/home.html")
    if err != nil {
        http.Error(w, "Unable to load templates", http.StatusInternalServerError)
        return
    }

    // Pass data to templates
    data := map[string]string{
        "Title":       "My Portfolio",
        "Description": "Welcome to my portfolio website! Here you can learn about me and my work.",
    }
    tmpl.Execute(w, data)
}

