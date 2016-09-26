package main

import (
	"encoding/json"
	"fmt"
	"html"
	"io"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
)

// Index page
//
func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
}

//TodoIndex page
//
func (i *Impl) TodoIndex(w http.ResponseWriter, r *http.Request) {
	todos := []Todo{}

	i.DB.Find(&todos)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(&todos); err != nil {
		fmt.Fprintf(w, "error enconding todos")
	}

}

//TodoShow handler
//
func TodoShow(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	todoID := vars["todoId"]
	fmt.Fprintln(w, "Todo show:", todoID)

	todo := Todo{}
	if DB.First(&todo, todoID).Error != nil {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "todo not found")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(&todo); err != nil {
		fmt.Fprintf(w, "error enconding todos")
	}

}

//TodoCreate handler
func TodoCreate(w http.ResponseWriter, r *http.Request) {

	todo := Todo{}

	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		panic(err)
	}
	if err := r.Body.Close(); err != nil {
		panic(err)
	}
	if err := json.Unmarshal(body, &todo); err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(422) // unprocessable entity
		if err := json.NewEncoder(w).Encode(err); err != nil {
			panic(err)
		}
	}

	if err := DB.Save(&todo).Error; err != nil {
		fmt.Fprintf(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(&todo); err != nil {
		fmt.Fprintf(w, "error enconding todos")
	}

	// if err := DB.Save(&todo).Error; err != nil {
	// 	fmt.Fprintf(w, err.Error(), http.StatusInternalServerError)
	// 	return
	// }

}
