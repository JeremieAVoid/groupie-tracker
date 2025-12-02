package main

import (
    "fmt"
    "groupie" // maintenant que ton module s'appelle groupie
    "log"
    "net/http"
    "os/exec"
    "strconv"
    "html/template"
)

func main() {
<<<<<<< HEAD
    tmplHome := template.Must(template.ParseFiles("HTML/homepage.html"))
    tmplMain := template.Must(template.ParseFiles("HTML/main.html"))

    // Route /
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        tmplHome.Execute(w, nil)
    })

    // Route /main
    http.HandleFunc("/main", func(w http.ResponseWriter, r *http.Request) {
        tmplMain.Execute(w, nil)
=======
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
	// http.HandleFunc("/CSS/", func(w http.ResponseWriter, r *http.Request) {
	// 	w.Header().Set("Content-Type", "text/css; charset=utf-8")
	// 	http.ServeFile(w, r, "CSS/styleTemplate.css")
	// })
>>>>>>> Emerick

        fmt.Println("début")
        lotDeListe := groupie.ChargerLesDonnées()

        http.HandleFunc("/Rechercher", func(w http.ResponseWriter, r *http.Request) {
            nombreAAfficherT := r.FormValue("nombreAAfficher")
            nombreAAfficher := groupie.TransformerEnNombre(nombreAAfficherT)
            liste := groupie.Recherche(lotDeListe, r.FormValue("catégorie"), r.FormValue("recherche"), nombreAAfficher)

            texte := ""
            for i := 0; i < len(liste); i++ {
                texte += strconv.Itoa(liste[i]) + "\n"
            }
            fmt.Fprintln(w, texte)
        })
    })

    // Routes CSS
    http.HandleFunc("/CSS/style.css", func(w http.ResponseWriter, r *http.Request) {
        w.Header().Set("Content-Type", "text/css; charset=utf-8")
        http.ServeFile(w, r, "CSS/style.css")
    })
    http.HandleFunc("/CSS/styleBarreDeRecherche.css", func(w http.ResponseWriter, r *http.Request) {
        w.Header().Set("Content-Type", "text/css; charset=utf-8")
        http.ServeFile(w, r, "CSS/styleBarreDeRecherche.css")
    })

    // Route /open
    http.HandleFunc("/open", func(w http.ResponseWriter, r *http.Request) {
        w.Header().Set("Content-Type", "text/plain; charset=utf-8")
        go func() {
            _ = exec.Command("xdg-open", "http://localhost:8080/").Start()
        }()
        w.Write([]byte("Attempted to open browser"))
    })

    // 🚀 Lancer le serveur directement
    log.Println("Serveur lancé sur http://localhost:8080")
    if err := http.ListenAndServe(":8080", nil); err != nil {
        log.Fatal(err)
    }
}