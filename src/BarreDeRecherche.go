package groupie

import "strconv"

// ici, on vas triée en fonction de la méthode choisie. Ce qui sera retourner sera le bon aurdre des ID.
func Trie(lotDeListe LotDeListe, méthode string, recherche string, nombreMaximum int) []int {
	// println("-------")
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
			if PeutÊtreVuAvecSeTermeDeRecherche(strconv.Itoa(listeDesArtistes[i].Id), recherche) {
				liste = append(liste, strconv.Itoa(listeDesArtistes[i].Id))
				listeN = append(listeN, listeDesArtistes[i].Id)
				listeID = append(listeID, listeDesArtistes[i].Id)
			}
		}
		listeDesID = TrieParOdreCroissant(listeN, listeID)
	case "Name":
		for i := 0; i < len(listeDesArtistes); i++ {
			if PeutÊtreVuAvecSeTermeDeRecherche(listeDesArtistes[i].Name, recherche) {
				liste = append(liste, listeDesArtistes[i].Name)
				listeID = append(listeID, listeDesArtistes[i].Id)
			}
		}
		listeDesID = TrieParOdreAlphabétique(liste, listeID)
	case "Members":
		for i := 0; i < len(listeDesArtistes); i++ {
			if PeutÊtreVuAvecSeTermeDeRecherche(strconv.Itoa(len(listeDesArtistes[i].Members)), recherche) {
				liste = append(liste, strconv.Itoa(listeDesArtistes[i].Id))
				listeN = append(listeN, len(listeDesArtistes[i].Members))
				listeID = append(listeID, listeDesArtistes[i].Id)
			}
		}
		listeDesID = TrieParOdreCroissant(listeN, listeID)
	case "CreationDate":
		for i := 0; i < len(listeDesArtistes); i++ {
			if PeutÊtreVuAvecSeTermeDeRecherche(strconv.Itoa(listeDesArtistes[i].CreationDate), recherche) {
				liste = append(liste, strconv.Itoa(listeDesArtistes[i].Id))
				listeN = append(listeN, listeDesArtistes[i].CreationDate)
				listeID = append(listeID, listeDesArtistes[i].Id)
			}
		}
		listeDesID = TrieParOdreCroissant(listeN, listeID)
	case "FirstAlbum":
		for i := 0; i < len(listeDesArtistes); i++ {
			if PeutÊtreVuAvecSeTermeDeRecherche(listeDesArtistes[i].FirstAlbum, recherche) {
				liste = append(liste, listeDesArtistes[i].FirstAlbum)
				listeID = append(listeID, listeDesArtistes[i].Id)
			}
		}
		listeDesID = TriéLesDates(liste, listeID)
	case "Locations":
		for i := 0; i < len(listeDesArtistes); i++ {
			if PeutÊtreVuAvecSeTermeDeRecherche(listeDesArtistes[i].Locations, recherche) {
				liste = append(liste, listeDesArtistes[i].Locations)
				listeID = append(listeID, listeDesArtistes[i].Id)
			}
		}
		listeDesID = TrieParOdreAlphabétique(liste, listeID)
	case "ConcertDates":
		for i := 0; i < len(listeDesDates); i++ {
			if PeutÊtreVuAvecSeTermeDeRecherche(strconv.Itoa(len(listeDesDates[i].Dates)), recherche) {
				liste = append(liste, strconv.Itoa(listeDesArtistes[i].Id))
				listeN = append(listeN, len(listeDesDates[i].Dates))
				listeID = append(listeID, listeDesDates[i].Id)
			}
		}
		listeDesID = TrieParOdreCroissant(listeN, listeID)
	}
	if len(listeDesID) > nombreMaximum && nombreMaximum > 0 {
		listeFini := []int{}
		for i := 0; i < nombreMaximum; i++ {
			listeFini = append(listeFini, listeDesID[i])
		}
		listeDesID = listeFini
	}
	if len(recherche) > 0 {
		listeDesID = TrierParPetinance(liste, listeDesID)
	}
	// println("--Début DEBUT trieParPetinance--")
	// for i := 0; i < len(listeDesID); i++ {
	// 	println(liste[i] + "  " + strconv.Itoa(listeDesID[i]))
	// }
	// println("--Fin--")
	return listeDesID
}

func Recherche(lotDeListe LotDeListe, catégorie string, recherche string, nombreMaximum int) []int {
	résultat := Trie(lotDeListe, catégorie, recherche, nombreMaximum)
	return résultat
}

/*
Quand un utilisateur appelle la fonction Rechercher avec la barre de recherche on vas :
1 - appeller Trie avec la bonne méthode
2 - avent de mettre les élément dans la liste il appelle PeutÊtreVuAvecSeTermeDeRecherche()
3 - affiche tous se qui en ressort
*/

/*
Que doit faire la barre de recherche ?
1 - donner se qui se rapprocher à 3 lettre prait du mots chercher
cherche : poisson
trouve : poison, poisson, poisons, poissonier
2 - Donne se qui contient la chose chercher à l'intérieur :
cherche : qui
trouve : quimange, avecquiilest, qui
*/

func PeutÊtreVuAvecSeTermeDeRecherche(résultat string, recherche string) bool {
	// 	recherche : 'qui'
	// 	résultat : 'qui mange', 'avec qui il est', 'qui'
	if résultat == recherche || len(recherche) == 0 {
		return true
	}
	max := len(résultat)
	compte := 0
	if len(recherche) <= max {
		for i := 0; i < len(résultat); i++ {
			if compte < len(recherche) {
				if résultat[i] == recherche[compte] {
					compte++
					if compte >= len(recherche) {
						return true
					}
				} else {
					compte = 0
				}
			}
		}
	}
	return false
}
