package main

import (
	b64 "encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

const baseUrl = "https://api.spotify.com/v1/"

// Genius Provider.
var accessToken = ""

func readBodyData(body io.ReadCloser) []byte {
	result, err := ioutil.ReadAll(body)
	if err != nil {
		log.Println("error parsing access token response ", err)
		return nil
	}
	return result
}

func authorize() {
	stringReq := fmt.Sprintf("%s:%s", os.Getenv("CLIENT_ID"), os.Getenv("CLIENT_SECRET"))
	encodedReq := b64.URLEncoding.EncodeToString([]byte(stringReq))
	data := url.Values{}
	data.Set("grant_type", "client_credentials")

	client := &http.Client{}
	req, err := http.NewRequest("POST", "https://accounts.spotify.com/api/token", strings.NewReader(data.Encode()))
	if err != nil {
		fmt.Println(err)
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Authorization", fmt.Sprintf("Basic %s", encodedReq))
	resp, err := client.Do(req)
	if err != nil {
		log.Println("Error searching for songs on Genius, ", err)
	}
	body := readBodyData(resp.Body)
	var result map[string]interface{}
	fmt.Println(string(body))
	err = json.Unmarshal(body, &result)
	if err != nil {
		log.Println("Error unmarshalling access token response ", err)
	}
	accessToken = result["access_token"].(string)
}

func search(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	query := vars["value"]
	url := baseUrl + "search?q=" + url.QueryEscape(query) + "&type=track"
	if accessToken == "" {
		authorize()
	}
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", accessToken))
	resp, err := client.Do(req)
	if err != nil {
		log.Println("Error searching for songs on Genius, ", err)
	}
	var data map[string]interface{}
	body := readBodyData(resp.Body)
	err = json.Unmarshal(body, &data)
	if err != nil {
		log.Println("Error unmarshalling response from search ", err)
	}
	toReturn := data["tracks"]
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(toReturn)
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

func main() {
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
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))
	log.Printf("Server listening on %s\n", srv.Addr)
	log.Fatal(srv.ListenAndServe())

}
