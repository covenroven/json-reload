package main

import (
	"html/template"
	"log"
	"net/http"
)

const port = "3400"

func main() {
	http.HandleFunc("/", showJSON)

	log.Printf("Server running on port %s\n", port)
	http.ListenAndServe(":"+port, nil)
}

func showJSON(w http.ResponseWriter, r *http.Request) {
	status, err := ReadJSON()
	if err != nil {
		log.Fatal(err)
	}

	status.Randomize()

	if err := WriteJSON(status); err != nil {
		log.Fatal(err)
	}

	tmpl, err := template.ParseFiles("html/index.html")
	if err != nil {
		log.Fatal(err)
	}

	tmpl.Execute(w, status)
}
