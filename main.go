package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type Knowledge struct {
	ID         string `json:"id"`
	User       string `json:"user"`
	TimeMarker string `json:"timemarker"`
}

var knowledges = []Knowledge{}
var idKnowlenge int

func main() {
	r := mux.NewRouter()
	knowledgeR := r.PathPrefix("/knowledge").Subrouter()
	knowledgeR.Path("").Methods(http.MethodGet).HandlerFunc(getAllKnowledges)
	knowledgeR.Path("").Methods(http.MethodPost).HandlerFunc(createKnowledge)
	knowledgeR.Path("/{id}").Methods(http.MethodGet).HandlerFunc(getKnowledge)
	knowledgeR.Path("/{id}").Methods(http.MethodPut).HandlerFunc(changeKnowledge)
	knowledgeR.Path("/{id}").Methods(http.MethodDelete).HandlerFunc(deleteKnowledge)

	fmt.Println("Greetings! Start listening")
	fmt.Println(http.ListenAndServe(":8080", r))
}

func getAllKnowledges(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")

	if err := json.NewEncoder(w).Encode(knowledges); err != nil {
		fmt.Println(err)
		http.Error(w, "Error encoding response object", http.StatusInternalServerError)
	}
}

func createKnowledge(w http.ResponseWriter, r *http.Request) {

	k := Knowledge{}

	if err := json.NewDecoder(r.Body).Decode(&k); err != nil {
		fmt.Println(err)
		http.Error(w, "Error decoding responce object", http.StatusBadRequest)
		return
	}

	knowledges = append(knowledges, k)

	response, err := json.Marshal(&k)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Error encoding responce object", http.StatusInternalServerError)
		return
	}

	w.Header().Add("Context-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(response)
}

func indexById(knowledges []Knowledge, id string) int {
	for i := 0; i < len(knowledges); i++ {
		if knowledges[i].ID == id {
			return i
		}
	}

	return -1
}

func getKnowledge(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	index := indexById(knowledges, id)

	if index < 0 {
		http.Error(w, "Knowledge not found", http.StatusNotFound)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(knowledges[index]); err != nil {
		fmt.Println(err)
		http.Error(w, "Error encoding responce object", http.StatusInternalServerError)
		return
	}
}

func changeKnowledge(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	index := indexById(knowledges, id)

	if index < 0 {
		http.Error(w, "Knowledge not found", http.StatusNotFound)
		return
	}

	k := Knowledge{}
	if err := json.NewDecoder(r.Body).Decode(&k); err != nil {
		fmt.Println(err)
		http.Error(w, "Error decoding responce object", http.StatusBadRequest)
		return
	}

	knowledges[index] = k

	w.Header().Add("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(&k); err != nil {
		fmt.Println(err)
		http.Error(w, "Error encoding responce object", http.StatusInternalServerError)
		return
	}
}

func deleteKnowledge(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	index := indexById(knowledges, id)

	if index < 0 {
		http.Error(w, "Knowledge not found", http.StatusNotFound)
		return
	}

	knowledges = append(knowledges[:index], knowledges[index+1:]...)
	w.WriteHeader(http.StatusOK)
}
