package vibely

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"time"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

const baseUrl = "https://genius.com"

func search(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	query := vars["value"]
	url := baseUrl + "/api/search/song?q=" + url.QueryEscape(query)
	resp, err := http.Get(url)
	if err != nil {
		log.Println("Error searching for songs on Genius, ", err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("Error parsing response, ", err)
	}
	var data map[string]interface{}
	err = json.Unmarshal(body, &data)
	if err != nil {
		log.Println("Error unmarshalling response from search ", err)
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}

func getSong(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	query := vars["path"]
	url := baseUrl + "/" + query
	res := crawlGetSong(url)
	fmt.Println(res)
}

func index(w http.ResponseWriter, r *http.Request) {
	indexFile, err := os.Open("./static/index.html")
	if err != nil {
		io.WriteString(w, "error reading index")
		return
	}
	defer indexFile.Close()

	io.Copy(w, indexFile)
}

func Start() {
	//create data.json if it doesn't exit
	// ensureDataExists()

	r := mux.NewRouter()
	_ = godotenv.Load()
	srv := &http.Server{
		Handler:      r,
		Addr:         "127.0.0.1:8996",
		WriteTimeout: 60 * time.Second,
		ReadTimeout:  60 * time.Second,
	}

	r.HandleFunc("/", index)
	r.Methods("GET").Path("/searchSongs{value}").HandlerFunc(search)
	r.Methods("GET").Path("/getSong{path}").HandlerFunc(getSong)
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))
	log.Printf("Server listening on %s\n", srv.Addr)
	log.Fatal(srv.ListenAndServe())

}
