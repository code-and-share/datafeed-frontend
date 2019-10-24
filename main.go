package main

import (
	"html/template"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

type WebData struct {
	Title  string
	Image1 string
	Image2 string
	Image3 string
	Image4 string
}

var wd = WebData{
	Title:  now,
	Image1: image_folder + "mountain001.jpg",
	Image2: image_folder + "forrest001.jpg",
	Image3: image_folder + "rain001.jpg",
	Image4: image_folder + "beach001.jpg",
}

var image_folder string = "/raw/"

var now string = time.Now().String()

func homeHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, _ := template.ParseFiles("templates/layout.html", "templates/home.html")
	//log.Println("Called home")
	//log.Println(wd)
	if err := tmpl.Execute(w, &wd); err != nil {
		log.Println(err.Error())
		http.Error(w, http.StatusText(500), 500)
	}
}

func postHandler(w http.ResponseWriter, r *http.Request) {
	// Must call ParseForm() before working with data
	r.ParseForm()
	// Log all data. Form is a map[]
	log.Println("Called post")
	log.Println(r.Form)
	wd = WebData{
		Title:  now,
		Image1: image_folder + "mountain002.jpg",
		Image2: image_folder + "forrest002.jpg",
		Image3: image_folder + "rain002.jpg",
		Image4: image_folder + "beach002.jpg",
	}
	w.WriteHeader(200)
	w.Write([]byte("ok cool"))
}

func Router() *mux.Router {
	router := mux.NewRouter()
	router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./static/"))))
	router.PathPrefix("/raw/").Handler(http.StripPrefix("/raw/", http.FileServer(http.Dir("./raw/"))))

	router.HandleFunc("/", homeHandler).Methods("GET")
	router.HandleFunc("/post", postHandler).Methods("POST")
	return router
}

func main() {
	PORT := "9000"
	log.Println("Serving on port " + PORT + "...")
	log.Fatal(http.ListenAndServe(":"+PORT, Router()))
}
