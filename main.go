package main

import (
	"log"
	"net/http"
)

//DB tool

func main() {
	//init()
	InitDatabase()

	router := NewRouter()

	log.Fatal(http.ListenAndServe(":8080", router))

	defer i.DB.Close()
}
