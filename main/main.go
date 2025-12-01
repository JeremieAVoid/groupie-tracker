package main

import (
	"net/http"
	"log"
	"html/template"
	
)

func main() {

tmplHome := template.Must(template.ParseFiles("html/homepage.html"))
tmplGroupie := template.Must(template.ParseFiles("html/Groupie.html"))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tmplHome.Execute(w, nil)
	})

	http.HandleFunc("/Groupie", func(w http.ResponseWriter, r *http.Request) {
		tmplGroupie.Execute(w, nil)
	})

			http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
			log.Println("Serveur lancé sur http://localhost:8080")
			log.Fatal(http.ListenAndServe(":8080", nil))
	}

