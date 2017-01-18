package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"sync"
	"time"
	"strings"
)

type comment struct {
	ID     int64  `json:"id"`
	Author string `json:"author"`
	Text   string `json:"text"`
}

type company struct {
	ID          int64  `json:"id"`
	Name        string `json:"name"`
	Industries  int32 `json:"industries"`
	Website     string `json:"website"`
	FoundDate   string `json:"foundDate"`
	StockCode   string `json:"stockCode"`
	Desc        string `json:"desc"`
}

const dataFile = "./server/data/comments.json"
const companiesFile = "./server/data/companies.json"

var commentMutex = new(sync.Mutex)

// Handle comments
func handleComments(w http.ResponseWriter, r *http.Request) {
	// Since multiple requests could come in at once, ensure we have a lock
	// around all file operations
	commentMutex.Lock()
	defer commentMutex.Unlock()

	// Stat the file, so we can find its current permissions
	fi, err := os.Stat(dataFile)
	if err != nil {
		http.Error(w, fmt.Sprintf("Unable to stat the data file (%s): %s", dataFile, err), http.StatusInternalServerError)
		return
	}

	// Read the comments from the file.
	commentData, err := ioutil.ReadFile(dataFile)
	if err != nil {
		http.Error(w, fmt.Sprintf("Unable to read the data file (%s): %s", dataFile, err), http.StatusInternalServerError)
		return
	}

	switch r.Method {
	case "POST":
		// Decode the JSON data
		var comments []comment
		if err := json.Unmarshal(commentData, &comments); err != nil {
			http.Error(w, fmt.Sprintf("Unable to Unmarshal comments from data file (%s): %s", dataFile, err), http.StatusInternalServerError)
			return
		}

		// Add a new comment to the in memory slice of comments
		comments = append(comments, comment{ID: time.Now().UnixNano() / 1000000, Author: r.FormValue("author"), Text: r.FormValue("text")})

		// Marshal the comments to indented json.
		commentData, err = json.MarshalIndent(comments, "", "    ")
		if err != nil {
			http.Error(w, fmt.Sprintf("Unable to marshal comments to json: %s", err), http.StatusInternalServerError)
			return
		}

		// Write out the comments to the file, preserving permissions
		err := ioutil.WriteFile(dataFile, commentData, fi.Mode())
		if err != nil {
			http.Error(w, fmt.Sprintf("Unable to write comments to data file (%s): %s", dataFile, err), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Cache-Control", "no-cache")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		io.Copy(w, bytes.NewReader(commentData))

	case "GET":
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Cache-Control", "no-cache")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		// stream the contents of the file to the response
		io.Copy(w, bytes.NewReader(commentData))

	default:
		// Don't know the method, so error
		http.Error(w, fmt.Sprintf("Unsupported method: %s", r.Method), http.StatusMethodNotAllowed)
	}
}


func Filter(s []company, fn func(company) bool) []company {
	var p []company // == nil
	for _, v := range s {
		if fn(v) {
			p = append(p, v)
		}
	}
	return p
}


// Handle companies
func handleCompanies(w http.ResponseWriter, r *http.Request) {
	// Stat the file, so we can find its current permissions
	_, err := os.Stat(companiesFile)
	if err != nil {
		errMsg := fmt.Sprintf("Unable to stat the data file (%s): %s", companiesFile, err)
		log.Printf(errMsg)
		http.Error(w, errMsg, http.StatusInternalServerError)
		return
	}

	// Read the comments from the file.
	data, err := ioutil.ReadFile(companiesFile) //TODO: not load file for each REST request
	if err != nil {
		http.Error(w, fmt.Sprintf("Unable to read the data file (%s): %s", companiesFile, err), http.StatusInternalServerError)
		return
	}

	var companies []company
	if err := json.Unmarshal(data, &companies); err != nil {
		http.Error(w, fmt.Sprintf("Unable to Unmarshal companies from data file (%s): %s", companiesFile, err), http.StatusInternalServerError)
		return
	}

	switch r.Method {
	case "GET":
		keywords := r.URL.Query()["keywords"]
		resultCompanies := companies
		if keywords != nil {
			resultCompanies = Filter(companies, func(c company) bool {
				return strings.Contains(c.Name, keywords[0]) || strings.Contains(c.Desc, keywords[0])
			} ) //TODO: ghetto!
		}
		resultData, err := json.Marshal(resultCompanies)
		if err != nil {
			http.Error(w, fmt.Sprintf("Failed to marshaling json: %s", err), http.StatusInternalServerError)
		}

		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Cache-Control", "no-cache")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		// stream the contents of the file to the response
		io.Copy(w, bytes.NewReader(resultData))

	default:
		// Don't know the method, so error
		http.Error(w, fmt.Sprintf("Unsupported method: %s", r.Method), http.StatusMethodNotAllowed)
	}
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}
	http.HandleFunc("/api/comments", handleComments)
	http.HandleFunc("/api/companies", handleCompanies)
	http.Handle("/", http.FileServer(http.Dir("./public")))
	log.Println("Server started: http://localhost:" + port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
