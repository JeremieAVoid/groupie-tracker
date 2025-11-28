package groupie

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
)

type ArtisteS struct {
	Id           int      `json:"id"`
	Image        string   `json:"image"`
	Name         string   `json:"name"`
	Members      []string `json:"members"`
	CreationDate int      `json:"creationDate"`
	FirstAlbum   string   `json:"firstAlbum"`
	Locations    string   `json:"locations"`
	ConcertDates string   `json:"concertDates"`
	Relations    string   `json:"relations"`
}

type LocationsS struct {
	Id        int      `json:"id"`
	Locations []string `json:"locations"`
	Dates     string   `json:"dates"`
}

type DatesS struct {
	Id    int      `json:"id"`
	Dates []string `json:"dates"`
}

type RelationS struct {
	Id             int                 `json:"id"`
	DatesLocations map[string][]string `json:"datesLocations"`
}

type LotDeListe struct {
	listeDesArtistes  []ArtisteS
	listeDesLocations []LocationsS
	listeDesDates     []DatesS
	listeDesRelations []RelationS
}

func ChargerLesDonnées() LotDeListe {
	lotDeListe := LotDeListe{}
	bloque := "https://groupietrackers.herokuapp.com/api/"
	lotDeListe.listeDesArtistes = ChargerLesArtistes(bloque + "artists")
	lotDeListe.listeDesLocations = ChargerLesLocations(bloque + "locations")
	lotDeListe.listeDesDates = ChargerLesDates(bloque + "dates")
	lotDeListe.listeDesRelations = ChargerLesRelation(bloque + "relation")

	// fmt.Println(listeDesArtistes[45-1].Name)
	// fmt.Println(len(listeDesArtistes))
	// fmt.Println(len(listeDesLocations))
	// fmt.Println(len(listeDesDates))
	// fmt.Println(len(listeDesRelations))
	return lotDeListe
}

func Ressource(url string) string {
	req, _ := http.NewRequest("GET", url, nil)
	res, _ := http.DefaultClient.Do(req)
	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)
	valeur := string(body)
	if valeur[0] != '[' {
		total := ""
		if url == "https://groupietrackers.herokuapp.com/api/locations" {
			for i := 9; i < len(valeur)-2; i++ {
				total += string(valeur[i])
			}
		} else if url == "https://groupietrackers.herokuapp.com/api/dates" {
			for i := 9; i < len(valeur)-2; i++ {
				total += string(valeur[i])
			}
		} else if url == "https://groupietrackers.herokuapp.com/api/relation" {
			for i := 9; i < len(valeur)-2; i++ {
				total += string(valeur[i])
			}
		}
		valeur = total
	}
	return valeur
}

func ChargerLesArtistes(url string) []ArtisteS {
	data := Ressource(url)
	var artiste []ArtisteS
	err := json.Unmarshal([]byte(data), &artiste)
	if err != nil {
		fmt.Println("Erreur :", err)
		return []ArtisteS{}
	}
	return artiste
}
func ChargerLesLocations(url string) []LocationsS {
	data := Ressource(url)
	var locations []LocationsS
	err := json.Unmarshal([]byte(data), &locations)
	if err != nil {
		fmt.Println("Erreur :", err)
		return []LocationsS{}
	}
	return locations
}
func ChargerLesDates(url string) []DatesS {
	data := Ressource(url)
	var dates []DatesS
	err := json.Unmarshal([]byte(data), &dates)
	if err != nil {
		fmt.Println("Erreur :", err)
		return []DatesS{}
	}
	return dates
}
func ChargerLesRelation(url string) []RelationS {
	data := Ressource(url)
	var relation []RelationS
	err := json.Unmarshal([]byte(data), &relation)
	if err != nil {
		fmt.Println("Erreur :", err)
		return []RelationS{}
	}
	return relation
}

// ici, on vas triée en fonction de la méthode choisie. Ce qui sera retourner sera le bon aurdre des ID.
func Trie(lotDeListe LotDeListe, méthode string) []int {
	listeDesArtistes := lotDeListe.listeDesArtistes
	// listeDesLocations := lotDeListe.listeDesLocations
	listeDesDates := lotDeListe.listeDesDates
	// listeDesRelations := lotDeListe.listeDesRelations

	listeDesID := []int{}
	liste := []string{}
	listeID := []int{}
	switch méthode {
	case "Id":
		for i := 0; i < len(listeDesArtistes); i++ {
			liste = append(liste, strconv.Itoa(listeDesArtistes[i].Id))
			listeID = append(listeID, listeDesArtistes[i].Id)
		}
	case "Name":
		for i := 0; i < len(listeDesArtistes); i++ {
			liste = append(liste, listeDesArtistes[i].Name)
			listeID = append(listeID, listeDesArtistes[i].Id)
		}
	case "Members":
		for i := 0; i < len(listeDesArtistes); i++ {
			liste = append(liste, strconv.Itoa(len(listeDesArtistes[i].Members)))
			listeID = append(listeID, listeDesArtistes[i].Id)
		}
	case "CreationDate":
		for i := 0; i < len(listeDesArtistes); i++ {
			liste = append(liste, strconv.Itoa(listeDesArtistes[i].CreationDate))
			listeID = append(listeID, listeDesArtistes[i].Id)
		}
	case "FirstAlbum":
		for i := 0; i < len(listeDesArtistes); i++ {
			liste = append(liste, listeDesArtistes[i].FirstAlbum)
			listeID = append(listeID, listeDesArtistes[i].Id)
		}
	case "Locations":
		for i := 0; i < len(listeDesArtistes); i++ {
			liste = append(liste, listeDesArtistes[i].Locations)
			listeID = append(listeID, listeDesArtistes[i].Id)
		}
	case "ConcertDates":
		for i := 0; i < len(listeDesDates); i++ {
			liste = append(liste, strconv.Itoa(len(listeDesDates[i].Dates)))
			listeID = append(listeID, listeDesDates[i].Id)
		}
		// case "EndConcert":
		// 	for i := 0; i < len(listeDesDates); i++ {
		// 		liste = append(liste, strconv.Itoa(listeDesDates[i].Dates[len(listeDesDates[i].Dates)-1]))
		// 		listeID = append(listeID, listeDesDates[i].Id)
		// 	}
	}
	listeDesID = TrieParOdreAlphabétique(liste, listeID)
	return listeDesID
}

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
	for i := 0; i < len(liste); i++ {
		println(liste[i] + "  " + strconv.Itoa(listeID[i]))
	}
	return listeID
}
