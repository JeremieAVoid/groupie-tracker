package groupie

import (
	"net/http"
	"strconv"
)

func CrééLeTexteListeMembre(blocArtiste ArtisteS) string {
	texteListe := ""
	for i := 0; i < len(blocArtiste.Members); i++ {
		if i != 0 {
			texteListe += "\n"
		}
		texteListe += "• " + blocArtiste.Members[i]
	}
	return texteListe
}

func CrééLeTexteListeConcert(date DatesS, localisation LocationsS) string {
	texteListe := ""

	for i := 0; i < len(localisation.Locations); i++ {
		if len(date.Dates) > i {
			if i != 0 {
				texteListe += "\n"
			}
			texteListe += "• " + localisation.Locations[i] + " : " + Remplacer(date.Dates[i], '*', 0)
		}
	}
	return texteListe
}

func ComplétéLaPageInformation(id int, listeID []int, lotDeListe LotDeListe, nomPage string, w http.ResponseWriter, r *http.Request) {
	if len(listeID) == 0 {
		for i := 0; i < len(lotDeListe.listeDesArtistes); i++ {
			listeID = append(listeID, i)
		}
	}
	if id < 0 {
		id = 0
	}
	if id >= len(listeID) {
		id = len(listeID) - 1
	}
	if id >= len(lotDeListe.listeDesArtistes) {
		return
	}
	blocArtiste := TrouverUnElementParID_ArtisteS(listeID[id], lotDeListe.listeDesArtistes)
	blocLocation := TrouverUnElementParID_LocationsS(listeID[id], lotDeListe.listeDesLocations)
	blocDate := TrouverUnElementParID_DatesS(listeID[id], lotDeListe.listeDesDates)

	data2 := PageData{
		Prénom:           blocArtiste.Name,
		Image:            blocArtiste.Image,
		DateDeCréation:   strconv.Itoa(blocArtiste.CreationDate),
		PremierAlbum:     blocArtiste.FirstAlbum,
		Id:               strconv.Itoa(blocArtiste.Id),
		TexteListeMembre: CrééLeTexteListeMembre(blocArtiste),
		ListeConcert:     CrééLeTexteListeConcert(blocDate, blocLocation),
	}
	PlacerUnePage(w, r, data2, nomPage)
}
