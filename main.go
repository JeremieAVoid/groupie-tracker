package main

import (
	"fmt"
	groupie "groupie/src"
	"log"
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
		nombreAAfficherT := r.FormValue("nombreAAfficher")
		nombreAAfficher := groupie.TransformerEnNombre(nombreAAfficherT)
		listeID = groupie.Recherche(lotDeListe, r.FormValue("catégorie"), r.FormValue("recherche"), nombreAAfficher)
		texte := ""
		for i := 0; i < len(listeID); i++ {
			texte += strconv.Itoa(listeID[i]) + "\n"
		}
		groupie.PlacerLesRésultaDeRecherche(w, r, listeID, lotDeListe)
		// fmt.Fprintln(w, texte)
	})

	http.HandleFunc("/informationsAppelle", func(w http.ResponseWriter, r *http.Request) {
		idT := r.FormValue("Id")
		id, err := strconv.Atoi(idT)
		if err != nil {
			fmt.Println(idT)
			fmt.Println("Problème !")
			// panic(err)
		} else {
			groupie.ComplétéLaPageInformation(id-1, listeID, lotDeListe, "static/templates/Informations.html", w, r)
		}
		// w.Header().Set("Content-Type", "text/html; charset=utf-8")
		// http.ServeFile(w, r, "HTML/Informations.html")
	})

	// 2 - Les CSS :
	// http.HandleFunc("static/style/style.css", func(w http.ResponseWriter, r *http.Request) {
	// 	w.Header().Set("Content-Type", "text/css; charset=utf-8")
	// 	http.ServeFile(w, r, "static/style/style.css")
	// })
	// http.HandleFunc("static/style/styleBarreDeRecherche.css", func(w http.ResponseWriter, r *http.Request) {
	// 	w.Header().Set("Content-Type", "text/css; charset=utf-8")
	// 	http.ServeFile(w, r, "static/style/styleBarreDeRecherche.css")
	// })
	// http.HandleFunc("static/style/styleInformation.css", func(w http.ResponseWriter, r *http.Request) {
	// 	w.Header().Set("Content-Type", "text/css; charset=utf-8")
	// 	http.ServeFile(w, r, "static/style/styleInformation.css")
	// })
	// http.HandleFunc("static/style/styleTemplate.css", func(w http.ResponseWriter, r *http.Request) {
	// 	w.Header().Set("Content-Type", "text/css; charset=utf-8")
	// 	http.ServeFile(w, r, "static/style/styleTemplate.css")
	// })


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
		http.ServeFile(w, r, "static/templates/main.html")
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

// package main

// import (
//     "fmt"
//     "groupie/src" // maintenant que ton module s'appelle groupie
//     "log"
//     "net/http"
//     "os/exec"
//     "strconv"
// )

// func main() {
// 	fmt.Println("début")
// 	// 1 Les fonctions :
// 	lotDeListe := groupie.ChargerLesDonnées()
// 	listeID := []int{}
//     // Route HTML
//     http.HandleFunc("/static/templates/Informations.html", func(w http.ResponseWriter, r *http.Request) {
//         w.Header().Set("Content-Type", "text/html; charset=utf-8")
//         http.ServeFile(w, r, "HTML/Informations.html")
//     })
//     // Routes CSS
//     http.HandleFunc("/static/style/style.css", func(w http.ResponseWriter, r *http.Request) {
//         w.Header().Set("Content-Type", "text/css; charset=utf-8")
//         http.ServeFile(w, r, "CSS/style.css")
//     })
//     http.HandleFunc("/static/style/styleBarreDeRecherche.css", func(w http.ResponseWriter, r *http.Request) {
//         w.Header().Set("Content-Type", "text/css; charset=utf-8")
//         http.ServeFile(w, r, "CSS/styleBarreDeR     echerche.css")
//     })
//     http.HandleFunc("/static/style/styleInformation.css", func(w http.ResponseWriter, r *http.Request) {
//         w.Header().Set("Content-Type", "text/css; charset=utf-8")
//         http.ServeFile(w, r, "CSS/styleInformation.css")
//     })
//     http.HandleFunc("/static/style/styleTemplate.css", func(w http.ResponseWriter, r *http.Request) {
//         w.Header().Set("Content-Type", "text/css; charset=utf-8")
//         http.ServeFile(w, r, "CSS/styleTemplate.css")
//     })
//     // Route /open
//     // http.HandleFunc("/open", func(w http.ResponseWriter, r *http.Request) {
//         // w.Header().Set("Content-Type", "text/plain; charset=utf-8")
//         go func() {
//             _ = exec.Command("xdg-open", "http://localhost:8080/static/templates/homepage.html").Start()
//         }()
//         // w.Write([]byte("Attempted to open browser"))
//     // })
// 	http.HandleFunc("/Rechercher", func(w http.ResponseWriter, r *http.Request) {
// 		nombreAAfficherT := r.FormValue("nombreAAfficher")
// 		nombreAAfficher := groupie.TransformerEnNombre(nombreAAfficherT)
// 		listeID = groupie.Recherche(lotDeListe, r.FormValue("catégorie"), r.FormValue("recherche"), nombreAAfficher)
// 		texte := ""
// 		for i := 0; i < len(listeID); i++ {
// 			texte += strconv.Itoa(listeID[i]) + "\n"
// 		}
// 		groupie.PlacerLesRésultaDeRecherche(w, r, listeID, lotDeListe)
// 		// fmt.Fprintln(w, texte)
// 	})
// 	http.HandleFunc("/informationsAppelle", func(w http.ResponseWriter, r *http.Request) {
// 		idT := r.FormValue("Id")
// 		id, err := strconv.Atoi(idT)
// 		if err != nil {
// 			fmt.Println(idT)
// 			fmt.Println("Problème !")
// 			// panic(err)
// 		} else {
// 			groupie.ComplétéLaPageInformation(id-1, listeID, lotDeListe, "static/templates/Informations.html", w, r)
// 		}
// 		// w.Header().Set("Content-Type", "text/html; charset=utf-8")
// 		// http.ServeFile(w, r, "HTML/Informations.html")
// 	})

//     if err := http.ListenAndServe(":8080", nil); err != nil {
//         log.Fatal(err)
//     } else {
//         log.Println("Serveur lancé sur http://localhost:8080")
//     }
// }
