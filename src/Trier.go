package groupie

import (
	"fmt"
	"strconv"
)

func TrieParOdreAlphabétique(liste []string, listeID []int) []int {
	//pour triée :
	//pour tous les éléments de la liste je sélectionne le plus petit et je le place au début
	// liste := []string{"ab", "aa", "b"}

	// for i := 0; i < len(liste); i++ {
	// 	println(liste[i] + "  " + strconv.Itoa(listeID[i]))
	// }
	for i := 0; i < len(liste)-1; i++ {
		plusPetit := i
		for j := i + 1; j < len(liste); j++ {
			if liste[plusPetit] > liste[j] {
				plusPetit = j
			}
		}
		if plusPetit != i {
			temps := liste[i]
			tempsID := listeID[i]
			liste[i] = liste[plusPetit]
			listeID[i] = listeID[plusPetit]
			liste[plusPetit] = temps
			listeID[plusPetit] = tempsID
		}
	}
	// for i := 0; i < len(liste); i++ {
	// 	println(liste[i] + "  " + strconv.Itoa(listeID[i]))
	// }
	return listeID
}

func TrieParOdreCroissant(liste []int, listeID []int) []int {
	for i := 0; i < len(liste)-1; i++ {
		plusPetit := i
		for j := i + 1; j < len(liste); j++ {
			if liste[plusPetit] > liste[j] {
				plusPetit = j
			}
		}
		if plusPetit != i {
			temps := liste[i]
			tempsID := listeID[i]
			liste[i] = liste[plusPetit]
			listeID[i] = listeID[plusPetit]
			liste[plusPetit] = temps
			listeID[plusPetit] = tempsID
		}
	}
	// for i := 0; i < len(liste); i++ {
	// 	println(strconv.Itoa(liste[i]) + "  " + strconv.Itoa(listeID[i]))
	// }
	return listeID
}

func TriéLesDates(liste []string, listeID []int) []int {
	//je vais tous les parcourire comme pour les autres fonction sauf que la, je vais vérifier, la date, mois, jours.
	// le seul format valide est : jj-mm-aaaa
	for i := 0; i < len(liste)-1; i++ {

		if len(liste[i]) != 10 {
			println(len(liste[i]))
			fmt.Println("Le format accepter pour les dates est '23-04-2006'. Merci de le RESPECTER !")
			continue
		}

		plusPetit := i
		jours_i := TransformerEnNombre(string(liste[plusPetit][0]) + string(liste[plusPetit][1]))
		mois_i := TransformerEnNombre(string(liste[plusPetit][3]) + string(liste[plusPetit][4]))
		année_i := TransformerEnNombre(string(liste[plusPetit][6]) + string(liste[plusPetit][7]) + string(liste[plusPetit][8]) + string(liste[plusPetit][9]))

		for j := i + 1; j < len(liste); j++ {
			jours_j := TransformerEnNombre(string(liste[j][0]) + string(liste[j][1]))
			mois_j := TransformerEnNombre(string(liste[j][3]) + string(liste[j][4]))
			année_j := TransformerEnNombre(string(liste[j][6]) + string(liste[j][7]) + string(liste[j][8]) + string(liste[j][9]))

			if (année_i > année_j) || (année_i == année_j && mois_i > mois_j) || (année_i == année_j && mois_i == mois_j && jours_i > jours_j) {
				plusPetit = j
				jours_i = TransformerEnNombre(string(liste[plusPetit][0]) + string(liste[plusPetit][1]))
				mois_i = TransformerEnNombre(string(liste[plusPetit][3]) + string(liste[plusPetit][4]))
				année_i = TransformerEnNombre(string(liste[plusPetit][6]) + string(liste[plusPetit][7]) + string(liste[plusPetit][8]) + string(liste[plusPetit][9]))
			}
		}
		if plusPetit != i {
			temps := liste[i]
			tempsID := listeID[i]
			liste[i] = liste[plusPetit]
			listeID[i] = listeID[plusPetit]
			liste[plusPetit] = temps
			listeID[plusPetit] = tempsID
		}
	}

	// for i := 0; i < len(liste); i++ {
	// 	println(liste[i] + "  " + strconv.Itoa(listeID[i]))
	// }
	return listeID
}

func TransformerEnNombre(texte string) int {
	valeur, err := strconv.Atoi(texte)
	if err != nil {
		fmt.Println("Prombèles rencontrer dans la conversion de nombre :")
		fmt.Println(err)
		// panic(err)
		return 0
	}
	return valeur
}

//Comment je vais calculer la pertinance ?
/*
 	recherche : 'qui'
 	résultat : 'qui mange', 'avec qui il est', 'qui'

Moins il y a de caractère autres que ceux rechercher, plus la pertinance est élevé.

Comme ils contiennent tous la chosse rechercher, inutile de le revérifier
*/
func TrierParPetinance(liste []string, listeID []int) []int {
	for i := 0; i < len(liste)-1; i++ {
		plusPetit := i
		for j := i + 1; j < len(liste); j++ {
			if len(liste[plusPetit]) > len(liste[j]) {
				plusPetit = j
			}
		}
		if plusPetit != i {
			temps := liste[i]
			tempsID := listeID[i]
			liste[i] = liste[plusPetit]
			listeID[i] = listeID[plusPetit]
			liste[plusPetit] = temps
			listeID[plusPetit] = tempsID
		}
	}
	return listeID
}
