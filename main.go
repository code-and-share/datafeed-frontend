package main

import (
	"html/template"
	"log"
	"net/http"
	"time"
)

type WebData struct {
	Title  string
	Image1 string
	Image2 string
	Image3 string
	Image4 string
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	image_folder := "/static/"
	tmpl, _ := template.ParseFiles("templates/layout.html", "templates/home.html")
	now := time.Now()
	wd := WebData{
		Title:  now.String(),
		Image1: image_folder + "mountain.jpg",
		Image2: image_folder + "forrest.jpg",
		Image3: image_folder + "coffee.jpg",
		Image4: image_folder + "beach.jpg",
	}
	if err := tmpl.Execute(w, &wd); err != nil {
		log.Println(err.Error())
		http.Error(w, http.StatusText(500), 500)
	}
}

func main() {
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))
	http.HandleFunc("/", homeHandler)
	port := ":9000"
	log.Println("Listening on port ", port)
	http.ListenAndServe(port, nil)
}
