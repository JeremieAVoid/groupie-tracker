package groupie

import (
	"net/http"
	"strconv"
)

func ComplétéLaPageInformation(id int, listeID []int, lotDeListe LotDeListe, nomPage string, w http.ResponseWriter, r *http.Request) {
	blocArtiste := TrouverUnElementParID_ArtisteS(listeID[id], lotDeListe.listeDesArtistes)
	texteListeMembre := ""
	for i := 0; i < len(blocArtiste.Members); i++ {
		if i != 0 {
			texteListeMembre += "\n"
		}
		texteListeMembre += "• " + blocArtiste.Members[i]
	}
	data2 := PageData{
		Prénom:           blocArtiste.Name,
		Image:            blocArtiste.Image,
		DateDeCréation:   strconv.Itoa(blocArtiste.CreationDate),
		PremierAlbum:     blocArtiste.FirstAlbum,
		Id:               strconv.Itoa(blocArtiste.Id),
		TexteListeMembre: texteListeMembre,
	}
	PlacerUnePage(w, r, data2, nomPage)
}
