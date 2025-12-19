package groupie

import (
	"html/template"
	"log"
	"net/http"
	"strconv"
)

// Cette fonction doit être appeller uniquement si la liste est d'une taille supérieur à 0
func TrouverUnElementParID_ArtisteS(ID int, liste []ArtisteS) ArtisteS {
	if len(liste) == 0 {
		println("Les TrouverUnElementParID_... ne prennent pas en compte les liste vide ! Veuillez vérifier la taille avant de l'appeller s'il vous plait.")
	}
	for i := 0; i < len(liste); i++ {
		if ID == liste[i].Id {
			return liste[i]
		}
	}
	println("Rien n'a été trouvé")
	return liste[0]
}

// Cette fonction doit être appeller uniquement si la liste est d'une taille supérieur à 0
func TrouverUnElementParID_LocationsS(ID int, liste []LocationsS) LocationsS {
	if len(liste) == 0 {
		println("Les TrouverUnElementParID_... ne prennent pas en compte les liste vide ! Veuillez vérifier la taille avant de l'appeller s'il vous plait.")
	}
	for i := 0; i < len(liste); i++ {
		if ID == liste[i].Id {
			return liste[i]
		}
	}
	println("Rien n'a été trouvé")
	return liste[0]
}

// Cette fonction doit être appeller uniquement si la liste est d'une taille supérieur à 0
func TrouverUnElementParID_DatesS(ID int, liste []DatesS) DatesS {
	if len(liste) == 0 {
		println("Les TrouverUnElementParID_... ne prennent pas en compte les liste vide ! Veuillez vérifier la taille avant de l'appeller s'il vous plait.")
	}
	for i := 0; i < len(liste); i++ {
		if ID == liste[i].Id {
			return liste[i]
		}
	}
	println("Rien n'a été trouvé")
	return liste[0]
}

// Cette fonction doit être appeller uniquement si la liste est d'une taille supérieur à 0
func TrouverUnElementParID_RelationS(ID int, liste []RelationS) RelationS {
	if len(liste) == 0 {
		println("Les TrouverUnElementParID_... ne prennent pas en compte les liste vide ! Veuillez vérifier la taille avant de l'appeller s'il vous plait.")
	}
	for i := 0; i < len(liste); i++ {
		if ID == liste[i].Id {
			return liste[i]
		}
	}
	println("Rien n'a été trouvé")
	return liste[0]
}

type PageData struct {
	Prénom                string
	Image                 string
	DateDeCréation        string
	Membres               string
	PremierAlbum          string
	VisiblePrénom         string
	VisibleImage          string
	VisibleDateDeCréation string
	VisibleMembres        string
	VisiblePremierAlbum   string
	Id                    string
}

func PlacerLesRésultaDeRecherche(w http.ResponseWriter, r *http.Request, listeID []int, lotDeListe LotDeListe) {
	//bloc principale :
	data := PageData{
		Prénom:                "",
		Image:                 "",
		DateDeCréation:        "",
		Membres:               "",
		PremierAlbum:          "",
		VisiblePrénom:         "",
		VisibleImage:          "invisible",
		VisibleDateDeCréation: "",
		VisibleMembres:        "",
		VisiblePremierAlbum:   "",
	}
	PlacerUnePage(w, r, data, "HTML/main.html")

	//page 1 :
	for i := 0; i < len(listeID); i++ {
		blocArtiste := TrouverUnElementParID_ArtisteS(listeID[i], lotDeListe.listeDesArtistes)
		data2 := PageData{
			Prénom:         blocArtiste.Name,
			Image:          blocArtiste.Image,
			DateDeCréation: strconv.Itoa(blocArtiste.CreationDate),
			Membres:        blocArtiste.Members[0],
			PremierAlbum:   blocArtiste.FirstAlbum,
			Id:             strconv.Itoa(i + 1),
		}
		if r.FormValue("Image") != "on" {
			data2.Image = ""
			data2.VisibleImage = "invisible"
		}
		if r.FormValue("Name") != "on" {
			data2.Prénom = ""
			data2.VisiblePrénom = "invisible"
		}
		if r.FormValue("CreationDate") != "on" {
			data2.DateDeCréation = ""
			data2.VisibleDateDeCréation = "invisible"
		}
		if r.FormValue("Members") != "on" {
			data2.Membres = ""
			data2.VisibleMembres = "invisible"
		}
		if r.FormValue("FirstAlbum") != "on" {
			data2.PremierAlbum = ""
			data2.VisiblePremierAlbum = "invisible"
		}

		PlacerUnePage(w, r, data2, "HTML/templateBlocSimple.html")
	}
}

func PlacerUnePage(w http.ResponseWriter, r *http.Request, data PageData, lienPage string) {
	tmpl, err := template.ParseFiles(lienPage)
	if err != nil {
		http.Error(w, "Erreur de chargement du template: "+err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	err = tmpl.Execute(w, data)
	if err != nil {
		log.Println("Erreur d'exécution du template:", err)
	}
}

func ComplétéLaPageInformation(id int, listeID []int, lotDeListe LotDeListe, nomPage string, w http.ResponseWriter, r *http.Request) {
	blocArtiste := TrouverUnElementParID_ArtisteS(listeID[id], lotDeListe.listeDesArtistes)

	data2 := PageData{
		Prénom:         blocArtiste.Name,
		Image:          blocArtiste.Image,
		DateDeCréation: strconv.Itoa(blocArtiste.CreationDate),
		Membres:        blocArtiste.Members[0],
		PremierAlbum:   blocArtiste.FirstAlbum,
		Id:             strconv.Itoa(blocArtiste.Id),
	}
	PlacerUnePage(w, r, data2, nomPage)
}
