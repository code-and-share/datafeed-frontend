package main

import (
	"fmt"
	"html/template"
	"io"
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

var phase int = 1
var phases_results []string

var image_folder string = "/raw/"

var now string = time.Now().String()

func HealthHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")

	io.WriteString(w, `{"alive": true}`)

}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, _ := template.ParseFiles("templates/layout.html", "templates/home.html")
	//log.Println("Called home")
	//log.Println(wd)
	if err := tmpl.Execute(w, &wd); err != nil {
		log.Println(err.Error())
		http.Error(w, http.StatusText(500), 500)
	}
}

func PostHandler(w http.ResponseWriter, r *http.Request) {
	// Must call ParseForm() before working with data
	r.ParseForm()
	// Log all data. Form is a map[]
	log.Println("Called post")
	log.Println(r.Form)
	if r.Form.Get("restart") == "true" {
		phase = 1
	} else {
		phases_results = append(phases_results, r.Form.Get("selected"))
		phase += 1
		log.Println(phases_results)
	}
	phase_string := fmt.Sprintf("%03d", phase)
	wd = WebData{
		Title:  now,
		Image1: image_folder + "mountain" + phase_string + ".jpg",
		Image2: image_folder + "forrest" + phase_string + ".jpg",
		Image3: image_folder + "rain" + phase_string + ".jpg",
		Image4: image_folder + "beach" + phase_string + ".jpg",
	}
	w.WriteHeader(200)
	w.Write([]byte("ok cool"))
}

func Router() *mux.Router {
	router := mux.NewRouter()
	router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./static/"))))
	router.PathPrefix("/raw/").Handler(http.StripPrefix("/raw/", http.FileServer(http.Dir("./raw/"))))

	router.HandleFunc("/", HomeHandler).Methods("GET")
	router.HandleFunc("/post", PostHandler).Methods("POST")
	router.HandleFunc("/health", HealthHandler).Methods("GET")
	return router
}

func main() {
	PORT := "9000"
	log.Println("Serving on port " + PORT + "...")
	log.Fatal(http.ListenAndServe(":"+PORT, Router()))
}
