package main

import (
	"database/sql"
	//"encoding/json"
	"errors"
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"strconv"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

// Struct that stores elements that the frontend will show
type WebData struct {
	Title  string
	Image1 string
	Image2 string
	Image3 string
	Image4 string
}

type ResultData struct {
	Text template.HTML
}

type Selection struct {
	position int
	selected string
}
type PhaseObject struct {
	position int
	object   string
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

var rd = ResultData{
	Text: template.HTML("Empty"),
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
	var session = ""
	cookie_session, _ := r.Cookie("session_id")
	if cookie_session != nil {
		session = cookie_session.Value
	}
	cookie_phase, _ := r.Cookie("phase")
	if cookie_phase == nil {
		phase = 1
	} else {
		var err error
		phase, err = strconv.Atoi(cookie_phase.Value)
		if err != nil {
			phase = 1
		}
	}
	PhaseBackend(session, phase, w)
}

func PostHandler(w http.ResponseWriter, r *http.Request) {
	var session = ""
	cookie_session, _ := r.Cookie("session_id")
	if cookie_session != nil {
		session = cookie_session.Value
	}
	cookie_phase, _ := r.Cookie("phase")
	if cookie_phase == nil {
		phase = 1
	} else {
		var err error
		phase, err = strconv.Atoi(cookie_phase.Value)
		if err != nil {
			phase = 1
		}
	}
	// Must call ParseForm() before working with data
	r.ParseForm()
	// Log all data. Form is a map[]
	log.Println("Called post")
	log.Println(r.Form)
	if r.Form.Get("restart") != "true" {
		//phases_results = append(phases_results, r.Form.Get("selected"))
		//log.Println(phases_results)
		ManageResult(session, phase, r.Form.Get("selected"))
	}
	log.Println("phase cookie says " + strconv.Itoa(phase))
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

func ManageResult(session string, phase int, selected string) {
	session_exists, session_ix := ContainsResultSession(session)
	if session_exists {
		log.Println("session " + session + " exists")
		this_selection := Selection{
			position: phase,
			selected: selected,
		}
		results[session_ix].selections = append(results[session_ix].selections, this_selection)
		log.Println(results[session_ix])
	} else {
		this_selection := Selection{
			position: phase,
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

func PhaseBackend(session string, phase int, w http.ResponseWriter) {
	switch phase {
	case 1, 2, 3, 4:
		var err error
		wd, err = PhaseDB(session, phase)
		if err != nil {
			log.Println("ERROR: " + err.Error())
		}
		tmpl, _ := template.ParseFiles("templates/selection_layout.html", "templates/selection.html")
		if err := tmpl.Execute(w, &wd); err != nil {
			log.Println(err.Error())
			http.Error(w, http.StatusText(500), 500)
		}
	default:
		// TODO: Use a second template for the results and Parse it here
		session_exists, session_ix := ContainsResultSession(session)
		if session_exists {
			result_all := `Here are your selections: <br>`
			selections_matrix := results[session_ix].selections
			for _, v := range selections_matrix {
				result_all = result_all + strconv.Itoa(v.position) + ` -> ` + v.selected + `<br>`
			}
			rd = ResultData{
				Text: template.HTML(result_all),
			}
		}
		tmpl, _ := template.ParseFiles("templates/result_layout.html", "templates/result.html")
		if err := tmpl.Execute(w, &rd); err != nil {
			log.Println(err.Error())
			http.Error(w, http.StatusText(500), 500)
		}
	}
}

func PhaseDB(session string, phase int) (WebData, error) {
	var chosen_path = "testpath_002"
	var DBerr error
	var result WebData
	phase_string := fmt.Sprintf("%03d", phase)
	dbInfo := "root:secret@tcp(127.0.0.1:3306)/Paths"
	db, err := sql.Open("mysql", dbInfo)
	if err != nil {
		DBerr = errors.New("Error running the query: " + err.Error())
		log.Println("ERROR: " + err.Error())
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		DBerr = errors.New("Error reaching the DB: " + err.Error())
		log.Println("ERROR: " + err.Error())
	} else {
		select_phase, err := db.Query("SELECT t1.pos, objects.content FROM objects, phases, JSON_TABLE(phases.objects, '$[*]' COLUMNS(pos INT PATH '$.position', obj VARCHAR(255) PATH '$.object')) AS t1 WHERE phases.id in (SELECT phase_id from paths WHERE name = '" + chosen_path + "' AND phase_order = " + strconv.Itoa(phase) + ") AND objects.name = t1.obj COLLATE utf8mb4_general_ci;")
		defer select_phase.Close()
		if err != nil {
			DBerr = errors.New("Error running the query: " + err.Error())
			log.Println("ERROR: " + err.Error())
		}
		var objects []PhaseObject
		for select_phase.Next() {
			var position int
			var object string

			if err := select_phase.Scan(&position, &object); err != nil {
				log.Println("ERROR: " + err.Error())
			}

			this_object := PhaseObject{
				position: position,
				object:   object,
			}
			objects = append(objects, this_object)
		}
		log.Println(objects)
		// TODO: use the object position instead of a fixed index
		result = WebData{
			Title:  phase_string,
			Image1: image_folder + objects[0].object,
			Image2: image_folder + objects[1].object,
			Image3: image_folder + objects[2].object,
			Image4: image_folder + objects[3].object,
		}
	}

	fmt.Print("result:")
	fmt.Println(result)
	fmt.Print("DBerr:")
	fmt.Println(DBerr)
	return result, DBerr
	// be careful deferring Queries if you are using transactions

}

func main() {
	PORT := "9000"
	log.Println("Serving on port " + PORT + "...")
	log.Fatal(http.ListenAndServe(":"+PORT, Router()))
}
