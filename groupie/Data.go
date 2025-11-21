package groupie

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
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

func ChargerLesDonnées() {
	bloque := "https://groupietrackers.herokuapp.com/api/"
	listeDesArtiste := ChargerLesArtistes(bloque + "artists")
	listeDesLocations := ChargerLesLocations(bloque + "locations")
	listeDesDates := ChargerLesDates(bloque + "dates")
	listeDesRelation := ChargerLesRelation(bloque + "relation")

	fmt.Println(listeDesArtiste[45-1].Name)
	fmt.Println(len(listeDesArtiste))
	fmt.Println(len(listeDesLocations))
	fmt.Println(len(listeDesDates))
	fmt.Println(len(listeDesRelation))
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
