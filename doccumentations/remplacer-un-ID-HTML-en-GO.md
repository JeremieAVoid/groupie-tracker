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

