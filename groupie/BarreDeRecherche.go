package groupie

// ici, on vas triée en fonction de la méthode choisie. Ce qui sera retourner sera le bon aurdre des ID.
func Trie(lotDeListe LotDeListe, méthode string, recherche string) []int {
	listeDesArtistes := lotDeListe.listeDesArtistes
	// listeDesLocations := lotDeListe.listeDesLocations
	listeDesDates := lotDeListe.listeDesDates
	// listeDesRelations := lotDeListe.listeDesRelations

	listeDesID := []int{}
	liste := []string{}
	listeN := []int{}
	listeID := []int{}
	switch méthode {
	case "Id":
		for i := 0; i < len(listeDesArtistes); i++ {
			listeN = append(listeN, listeDesArtistes[i].Id)
			listeID = append(listeID, listeDesArtistes[i].Id)
		}
		listeDesID = TrieParOdreCroissant(listeN, listeID)
	case "Name":
		for i := 0; i < len(listeDesArtistes); i++ {
			liste = append(liste, listeDesArtistes[i].Name)
			listeID = append(listeID, listeDesArtistes[i].Id)
		}
		listeDesID = TrieParOdreAlphabétique(liste, listeID)
	case "Members":
		for i := 0; i < len(listeDesArtistes); i++ {
			listeN = append(listeN, len(listeDesArtistes[i].Members))
			listeID = append(listeID, listeDesArtistes[i].Id)
		}
		listeDesID = TrieParOdreCroissant(listeN, listeID)
	case "CreationDate":
		for i := 0; i < len(listeDesArtistes); i++ {
			listeN = append(listeN, listeDesArtistes[i].CreationDate)
			listeID = append(listeID, listeDesArtistes[i].Id)
		}
		listeDesID = TrieParOdreCroissant(listeN, listeID)
	case "FirstAlbum":
		for i := 0; i < len(listeDesArtistes); i++ {
			liste = append(liste, listeDesArtistes[i].FirstAlbum)
			listeID = append(listeID, listeDesArtistes[i].Id)
		}
		listeDesID = TriéLesDates(liste, listeID)
	case "Locations":
		for i := 0; i < len(listeDesArtistes); i++ {
			liste = append(liste, listeDesArtistes[i].Locations)
			listeID = append(listeID, listeDesArtistes[i].Id)
		}
		listeDesID = TrieParOdreAlphabétique(liste, listeID)
	case "ConcertDates":
		for i := 0; i < len(listeDesDates); i++ {
			listeN = append(listeN, len(listeDesDates[i].Dates))
			listeID = append(listeID, listeDesDates[i].Id)
		}
		listeDesID = TrieParOdreCroissant(listeN, listeID)
	}
	return listeDesID
}

func Recherche(lotDeListe LotDeListe, catégorie string, recherche string) []int {
	Trie(lotDeListe, catégorie, recherche)
	return []int{}
}

func PeutÊtreVuAvecSeTermeDeRecherche(résultat string, recherche string) bool {
	return true
}

/*
Quand un utilisateur appelle la fonction Rechercher avec la barre de recherche on vas :
1 - appeller Trie avec la bonne méthode
2 - avent de mettre les élément dans la liste il appelle PeutÊtreVuAvecSeTermeDeRecherche()
3 - affiche tous se qui en ressort
*/
