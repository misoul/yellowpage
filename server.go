package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/misoul/yellowpage/dal/mysql"
	"github.com/misoul/yellowpage/dal"
)

var dbUrl = "root:yellowpage@tcp(192.168.99.100:3306)/testdb1?parseTime=True"
var commentService, _ = mysql.InitComment(dbUrl)
var companyService, _ = mysql.InitCompany(dbUrl)

// Handle comments
func handleComments(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		// Add a new comment to the in memory slice of comments
		_, err := commentService.Update(dal.Comment{Author: r.FormValue("author"), Text: r.FormValue("text")})
		if err != nil {
			http.Error(w, fmt.Sprintf("Unable to add comment: %s", err), http.StatusInternalServerError)
			return
		}

		// Marshal the comments to indented json.
		comments, _ :=  commentService.Search(nil)
		commentData, err := json.Marshal(comments)
		if err != nil {
			http.Error(w, fmt.Sprintf("Unable to marshal comments to json: %s", err), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Cache-Control", "no-cache")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		io.Copy(w, bytes.NewReader(commentData))

	case "GET":
		resultComments, _ := commentService.Search(nil)
		resultData, err := json.Marshal(resultComments)
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


// Handle companies
func handleCompanies(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		keywords := r.URL.Query()["keywords"]
		resultCompanies, _ := companyService.Search(keywords)
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Cache-Control", "no-cache")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		// stream the contents of the file to the response
		resultData, err := json.Marshal(resultCompanies)
		if err != nil {
			http.Error(w, fmt.Sprintf("Failed to marshaling json: %s", err), http.StatusInternalServerError)
		}
		io.Copy(w, bytes.NewReader(resultData))
	default:
		// Don't know the method, so error
		http.Error(w, fmt.Sprintf("Unsupported method: %s", r.Method), http.StatusMethodNotAllowed)
	}
}

func Log(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s %s %s", r.RemoteAddr, r.Method, r.URL)
		handler.ServeHTTP(w, r)
	})
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
	log.Fatal(http.ListenAndServe(":"+port, Log(http.DefaultServeMux)))

	commentService.Finalize()
	companyService.Finalize()
}
