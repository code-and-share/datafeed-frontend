package main

import (
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"strconv"
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

type Selection struct {
	position int
	selected string
}

type Results struct {
	session_id string
	selections []Selection
}

var results []Results

var phase int = 1

var wd = WebData{
	Title:  strconv.Itoa(phase),
	Image1: image_folder + "mountain001.jpg",
	Image2: image_folder + "forest001.jpg",
	Image3: image_folder + "rain001.jpg",
	Image4: image_folder + "beach001.jpg",
}

var phases_results []string

var image_folder string = "/raw/"

var now string = time.Now().String()

func HealthHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")

	io.WriteString(w, `{"alive": true}`)

}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	cookie_phase, _ := r.Cookie("phase")
	if cookie_phase == nil {
		phase = 1
	} else {
		phase, _ = strconv.Atoi(cookie_phase.Value)
	}
	phase_string := fmt.Sprintf("%03d", phase)
	wd = WebData{
		Title:  phase_string,
		Image1: image_folder + "mountain" + phase_string + ".jpg",
		Image2: image_folder + "forest" + phase_string + ".jpg",
		Image3: image_folder + "rain" + phase_string + ".jpg",
		Image4: image_folder + "beach" + phase_string + ".jpg",
	}
	tmpl, _ := template.ParseFiles("templates/layout.html", "templates/home.html")
	if err := tmpl.Execute(w, &wd); err != nil {
		log.Println(err.Error())
		http.Error(w, http.StatusText(500), 500)
	}
}

func PostHandler(w http.ResponseWriter, r *http.Request) {
	cookie_session, _ := r.Cookie("session_id")
	cookie_phase, _ := r.Cookie("phase")
	// Must call ParseForm() before working with data
	r.ParseForm()
	// Log all data. Form is a map[]
	log.Println("Called post")
	log.Println(r.Form)
	if r.Form.Get("restart") != "true" {
		//phases_results = append(phases_results, r.Form.Get("selected"))
		//log.Println(phases_results)
		ManageResult(cookie_session.Value, cookie_phase.Value, r.Form.Get("selected"))
	}
	log.Println("phase cookie says " + cookie_phase.Value)
	w.WriteHeader(200)
	w.Write([]byte("ok cool"))
}

func ContainsResultSession(session string) (bool, int) {
	for ix, a := range results {
		if a.session_id == session {
			return true, ix
		}
	}
	return false, 0
}

func ManageResult(session string, phase string, selected string) {
	session_exists, session_ix := ContainsResultSession(session)
	if session_exists {
		log.Println("session " + session + " exists")
		phase_int, _ := strconv.Atoi(phase)
		this_selection := Selection{
			position: phase_int,
			selected: selected,
		}
		results[session_ix].selections = append(results[session_ix].selections, this_selection)
		log.Println(results[session_ix])
	} else {
		phase_int, _ := strconv.Atoi(phase)
		this_selection := Selection{
			position: phase_int,
			selected: selected,
		}
		selections := []Selection{}
		selections = append(selections, this_selection)
		these_results := Results{
			session_id: session,
			selections: selections,
		}
		results = append(results, these_results)
		log.Println(results)
	}
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
