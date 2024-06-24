package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type Knowledge struct {
	ID         string `json:"id"`
	User       string `json:"user"`
	TimeMarker string `json:"timemarker"`
}

var knowledges = []Knowlenge{}
var idKnowlenge int

func main() {
	r := mux.NewRouter()
	knowledgeR := r.PathPrefix("/knowledge").Subrouter()
	knowledgeR.Path("").Methods(http.MethodGet).HandlerFunc(getAllKnowledges)
	knowledgeR.Path("").Methods(http.MethodPost).HandlerFunc(createKnowledge)
	knowledgeR.Path("/{id}").Methods(http.MethodGet).HandlerFunc(getKnowledge)
	knowledgeR.Path("/{id}").Methods(http.MethodDelete).HandlerFunc(deleteKnowledge)

	fmt.Println("Greetings! Start listening")
	fmt.Println(http.ListenAndServe(":8080", r))
}

func getAllKnowledges(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Nothing")
}

func createKnowledge(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Nothing")
}

func getKnowledge(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Nothing")
}

func deleteKnowledge(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Nothing")
}
