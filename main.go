package main

import (
	"fmt"
	groupie "groupie/src"
	"log"
	"math/rand"
	"net/http"
	"os/exec"
	"strconv"
)

func main() {
	fmt.Println("début")
	// 1 Les fonctions :
	lotDeListe := groupie.ChargerLesDonnées()
	// lotDeListe := groupie.LotDeListe{}
	listeID := []int{}

	http.HandleFunc("/Rechercher", func(w http.ResponseWriter, r *http.Request) {
		listeID = FonctionRecherche(w, r, listeID, lotDeListe)
	})

	http.HandleFunc("/informationsAppelle", func(w http.ResponseWriter, r *http.Request) {
		FonctionInformationsAppelle(w, r, listeID, lotDeListe)
	})

	http.HandleFunc("/CliqueBoutonDeNavigation", func(w http.ResponseWriter, r *http.Request) {
		FonctionCliqueBoutonDeNavigation(w, r, listeID, lotDeListe)
	})

	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.HandleFunc("static/templates/Informations.html", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		http.ServeFile(w, r, "static/templates/Informations.html")
	})

	// 3 - Démarer le serveur :

	log.Println("Serveur lancé sur http://localhost:8080")
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		http.ServeFile(w, r, "static/templates/homepage.html")
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

func FonctionRecherche(w http.ResponseWriter, r *http.Request, listeID []int, lotDeListe groupie.LotDeListe) []int {
	if r.FormValue("pageAcceuil") == "1" {
		listeID = groupie.Recherche(lotDeListe, "Name", "", 50)
		groupie.PlacerLesRésultaDeRecherche(w, r, listeID, lotDeListe)
	} else {
		nombreAAfficherT := r.FormValue("nombreAAfficher")
		nombreAAfficher := groupie.TransformerEnNombre(nombreAAfficherT)
		listeID = groupie.Recherche(lotDeListe, r.FormValue("catégorie"), r.FormValue("recherche"), nombreAAfficher)
		texte := ""
		for i := 0; i < len(listeID); i++ {
			texte += strconv.Itoa(listeID[i]) + "\n"
		}
		groupie.PlacerLesRésultaDeRecherche(w, r, listeID, lotDeListe)
	}
	return listeID
}

func FonctionCliqueBoutonDeNavigation(w http.ResponseWriter, r *http.Request, listeID []int, lotDeListe groupie.LotDeListe) {
	nom := "main.html"
	switch r.FormValue("idBouton") {
	case "0":
		nom = "homepage.html"
	case "1":
		FonctionRecherche(w, r, listeID, lotDeListe)
		nom = "main.html"
		return
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	http.ServeFile(w, r, "static/templates/"+nom)
}

func FonctionInformationsAppelle(w http.ResponseWriter, r *http.Request, listeID []int, lotDeListe groupie.LotDeListe) {
	id := 0
	erreur := false
	if r.FormValue("idBoutonAléatoire") == "2" {
		id = rand.Intn(groupie.NombreLotDeListe(lotDeListe))
	} else {
		idT := r.FormValue("Id")
		id2, err := strconv.Atoi(idT)
		if err != nil {
			erreur = true
			fmt.Println(idT)
			fmt.Println("Problème !")
			// panic(err)
		} else {
			id = id2
		}
	}
	if !erreur {
		groupie.ComplétéLaPageInformation(id-1, listeID, lotDeListe, "static/templates/Informations.html", w, r)
	}
	// w.Header().Set("Content-Type", "text/html; charset=utf-8")
	// http.ServeFile(w, r, "HTML/Informations.html")
}
