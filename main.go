package main

import (
	"fmt"
	"groupie"
	"log"
	"net/http"
	"os/exec"
	"strconv"
)

func main() {
	fmt.Println("début")
	// 1 Les fonctions :
	lotDeListe := groupie.ChargerLesDonnées()

	http.HandleFunc("/Rechercher", func(w http.ResponseWriter, r *http.Request) {
		nombreAAfficherT := r.FormValue("nombreAAfficher")
		nombreAAfficher := groupie.TransformerEnNombre(nombreAAfficherT)
		liste := groupie.Recherche(lotDeListe, r.FormValue("catégorie"), r.FormValue("recherche"), nombreAAfficher)
		texte := ""
		for i := 0; i < len(liste); i++ {
			texte += strconv.Itoa(liste[i]) + "\n"
		}
		groupie.PlacerLesRésultaDeRecherche(w, r, liste, lotDeListe)
		// fmt.Fprintln(w, texte)
	})

	http.HandleFunc("/informations", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		http.ServeFile(w, r, "HTML/Informations.html")
	})

	// 2 - Les CSS :
	http.HandleFunc("/CSS/style.css", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/css; charset=utf-8")
		http.ServeFile(w, r, "CSS/style.css")
	})
	http.HandleFunc("/CSS/styleBarreDeRecherche.css", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/css; charset=utf-8")
		http.ServeFile(w, r, "CSS/styleBarreDeRecherche.css")
	})
	http.HandleFunc("/CSS/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/css; charset=utf-8")
		http.ServeFile(w, r, "CSS/styleInformation.css")
	})

	// 3 - Démarer le serveur :

	log.Println("Serveur lancé sur http://localhost:8080")
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		http.ServeFile(w, r, "HTML/main.html")
	})

	http.HandleFunc("/open", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
		go func() {
			_ = exec.Command("xdg-open", "http://localhost:8080/").Start()
		}()
		w.Write([]byte("Attempted to open browser"))
	})

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
