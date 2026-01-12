Structure du projet :
```
├─src
├─┬─BarreDeRecherche.go #Script de la barre de recherche.
├─┼─ComplétéLaPageInformation.go #Script qui écrit les informations dans la page information.
├─┼─Data.go #Script qui récupère les données de l'API et les structure dans des structs dédiés.
├─┼─PlacerLesRésultaDeRecherche.go #Place les résultat de la recherche sur la page recherche.html suite à une recherche.
├─┼─Remplacer.go #Contient le script qui permet de remplacer un caractère par un autre dans une chaine de caractère.
├─┼─Trier.go #Contient les script utilisé pour triées les données suite à une recherche.
├─┼─TrouverUnFichierParID.go #Permet de retrouver un éléments dans un strct choisie en fonction de son identifiant.
│
├─static
├─┼─font #Ces un dossier contenant plusieurs polices d'écriture.
├─┼─img #Ces un dossier contenant les images utilisé dans le projet.
├─┼─script
├─┼─┼─Yeux.js #Script permettant d'annimé les yeux du bouton de la page d'aceuil.
│
├─┼─style
├─┼─┼─boutonAvecYeux.css #Effet CSS utilisé pour le boutons avec des yeux utilisé sur la page d'aceuil.
├─┼─┼─logo.css #style CSS du logo du site.
├─┼─┼─stylePageDeRecherche.css #CSS de la page de recherche.
├─┼─┼─styleBarreDeRecherche.css #CSS de la barre de recherche.
├─┼─┼─styleGénéral.css #CSS général du site.
├─┼─┼─styleHautDePage.css #CSS du <header> des pages HTML.
├─┼─┼─styleHomePage.css #CSS de la page d'aceuil.
├─┼─┼─styleInformation.css #CSS de la page d'information.
├─┼─┼─styleTemplate.css #CSS des bloques de groupe placer sur la page recherche suite à une recherche.
│
├─┼─templates
├─┼─┼─homepage.html #Page d'aceuil
├─┼─┼─Informations.html #Page d'information
├─┼─┼─Logo.html #Page de test du logo.
├─┼─┼─recherche.httml #Page de recherche
├─┼─┼─templateBlocSimple.html #Bloques de groupe placer sur la page recherche suite à une recherche.
│
├─go.mod #Module groupie utilisé dans les script GO.
├─README.md #Le document que vous êtes en train de lire.
```





Pour le test :

# Comment faire pour remplacer un ID HTML en GO ?

Les scriptes :
``` go

type PageData struct {
	Prénom string
}

//---------------
data := PageData{
	Prénom: "nom mis",   
}
lienDuHTML := "static/templates/recherche.html"

tmpl, err := template.ParseFiles(lienDuHTML)
if err != nil {
	http.Error(w, "Erreur de chargement du template: "+err.Error(), http.StatusInternalServerError)
	return
}
w.Header().Set("Content-Type", "text/html; charset=utf-8")
err = tmpl.Execute(w, data)
if err != nil {
	log.Println("Erreur d'exécution du template:", err)
}
```

Et en HTML :
``` html
<span id="prénom">{{.Prénom}}</span>
```

