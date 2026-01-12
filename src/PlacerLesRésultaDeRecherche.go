package groupie

import (
	"net/http"
	"strconv"
)

func PlacerLesRésultaDeRecherche(w http.ResponseWriter, r *http.Request, listeID []int, lotDeListe LotDeListe) {
	//bloc principale :
	data := PageData{
		Prénom:                "",
		Image:                 "",
		DateDeCréation:        "",
		PremierAlbum:          "",
		VisiblePrénom:         "",
		VisibleImage:          "invisible",
		VisibleDateDeCréation: "",
		VisibleMembres:        "",
		VisiblePremierAlbum:   "",
	}
	PlacerUnePage(w, r, data, "static/templates/recherche.html")

	//page 1 :
	for i := 0; i < len(listeID); i++ {
		blocArtiste := TrouverUnElementParID_ArtisteS(listeID[i], lotDeListe.listeDesArtistes)

		data2 := PageData{
			Prénom:           blocArtiste.Name,
			Image:            blocArtiste.Image,
			DateDeCréation:   strconv.Itoa(blocArtiste.CreationDate),
			PremierAlbum:     blocArtiste.FirstAlbum,
			Id:               strconv.Itoa(i + 1),
			TexteListeMembre: CrééLeTexteListeMembre(blocArtiste),
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
			data2.TexteListeMembre = ""
			data2.VisibleMembres = "invisible"
		}
		if r.FormValue("FirstAlbum") != "on" {
			data2.PremierAlbum = ""
			data2.VisiblePremierAlbum = "invisible"
		}

		PlacerUnePage(w, r, data2, "static/templates/templateBlocSimple.html")
	}
}
