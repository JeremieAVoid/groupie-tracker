# Projet Groupy Tracker.
## • Architecture Technique.
### 1 - Structure du Projet.
```
┌─src
├─┬─BarreDeRecherche.go 			# Script de la barre de recherche.
├─┼─ComplétéLaPageInformation.go 	# Script qui écrit les informations dans la page information.
├─┼─Data.go 						# Script qui récupère les données de l'API et les structure dans des structs dédiés.
├─┼─PlacerLesRésultaDeRecherche.go 	# Place les résultat de la recherche sur la page recherche.html suite à une recherche.
├─┼─Remplacer.go 					# Contient le script qui permet de remplacer un caractère par un autre dans une chaine de caractère.
├─┼─Trier.go 						# Contient les script utilisé pour triées les données suite à une recherche.
├─┴─TrouverUnFichierParID.go 		# Permet de retrouver un éléments dans un strct choisie en fonction de son identifiant.
│
├─static
├─┬─font 							# Ces un dossier contenant plusieurs polices d'écriture.
├─┼─img 							# Ces un dossier contenant les images utilisé dans le projet.
├─┼─script
├─┼───Yeux.js 						# Script permettant d'annimé les yeux du bouton de la page d'aceuil.
├─┼─style
├─┼─┬─boutonAvecYeux.css 			# Effet CSS utilisé pour le boutons avec des yeux utilisé sur la page d'aceuil.
├─┼─┼─logo.css 						# CSS du logo du site.
├─┼─┼─stylePageDeRecherche.css 		# CSS de la page de recherche.
├─┼─┼─styleBarreDeRecherche.css	 	# CSS de la barre de recherche.
├─┼─┼─styleGénéral.css 				# CSS général du site.
├─┼─┼─styleHautDePage.css 			# CSS du <header> des pages HTML.
├─┼─┼─styleHomePage.css 			# CSS de la page d'aceuil.
├─┼─┼─styleInformation.css 			# CSS de la page d'information.
├─┼─┴─styleTemplate.css 			# CSS des bloques de groupe placer sur la page recherche suite à une recherche.
├─┼─templates
├─┼─┬─homepage.html 				# Page d'aceuil.
├─┼─┼─Informations.html 			# Page d'information.
├─┼─┼─Logo.html 					# Page de test du logo.
├─┼─┼─recherche.httml 				# Page de recherche.
├─┴─┴─templateBlocSimple.html 		# Bloques de groupe placer sur la page recherche suite à une recherche.
│
├─go.mod 							# Module groupie utilisé dans les script GO.
└─README.md 						# Le document que vous êtes en train de lire.
```

### 2 - Langage Utilisées.
- HTML : Pour les pages du sites.
- CSS : Pour le styles des pages.
- JavaScript : Pour une annimation sur le site.
- GO : Pour l'intéraction utilisateurs et l'affichages des éléments sur la page.

## • Scripts Disponibles
BarreDeRecherche.go :
- Recherche() : prend en paramètre des paramètres de recherche et une structure de liste d'éléments et appelle Trie() puis renvoie les éléments qui peuvent apparètres dans les résultats de recherches.
- Trie() : prend en paramètre des paramètres de recherche et une structure de liste d'éléments et renvoie la liste des identifiants des éléments qui peuvent aparaitre avec les paramètre de recherches choisies.
- PeutÊtreVuAvecSeTermeDeRecherchet() : prend en paramètre deux chaines de carractère une pour ceux qui est chercher et l'autre pour le résultat. Il renvoie un booléen pour dire si avec ce qui est chercher se résulta peut être afficher dans les résultats de recherches.
- ToUpper() : Permet de passer une chaine de caractère en majuscule.

ComplétéLaPageInformation.go : 
- CrééLeTexteListeMembre() : prend en paramètre une structure ArtisteS et renvoie un texte contenant le nom de chaque m'embre de se groupe. 
- CrééLeTexteListeConcert() : prend en paramètre la structure DatesS et LocationS et renvoie un texte contenant la liste des concerts réalisés.
- ComplétéLaPageInformation() : prend en paramètre l'identifiant de la page sur la quel ont veut avoir des information ainsi que les listes contenant les données avec les information puis réalise la page en fonction de ces données.

Data.go :
- NombreLotDeListe() : Permet de dire le nombre de groupe présant dans la liste.
- ChargerLesDonnées() : Il s'agit de la fonction qui vas réaliser un appel à l'API pour remplirent les données locals.
- Ressource() : Prend en paramètre une URL et retourne ceux qu'elle contenait.
- ChargerLesArtistes() : Fonction appeller par ChargerLesDonnées() pour charger la partie des artistes.
- ChargerLesLocations() : Fonction appeller par ChargerLesDonnées() pour charger la partie des localisations.
- ChargerLesDates() : Fonction appeller par ChargerLesDonnées() pour charger la partie des dates.
- ChargerLesRelation() : Fonction appeller par ChargerLesDonnées() pour charger la partie des relations.

PlacerLesRésultaDeRecherche.go :
- PlacerLesRésultaDeRecherche() : Vas écrire ajouter les résultats de la recherche sur la pages. C'est cette fonction qui affichera ou non l'image du groupe, le nom, la date, [...] en fonction de se que l'utilisateur à choisie.

Remplacer.go :
- Remplacer() : Permet de remplacer un caractère par un autre dans une phrase.

Trier.go :
- TrieParOdreAlphabétique() : Cette fonction retourne une liste triée dans l'ordre alphabétique.
- TrieParOdreCroissant() : Cette fonction retourne une liste triée dans l'ordre croisant.
- TriéLesDates() : Cette fonction retourne une liste triée dans l'ordre croisant des date qui ont la forme jj-mm-aaaa.
- TransformerEnNombre() : Permet de transformer un texte en un nombre et de renvoyer une erreur si se n'est pas possible.
- TrierParPetinance() : Permet de triée la liste par pertinance en mettant en premier ce qui est le plus pertinant celon la recherche (si ont recherche 'pomme', il est plus pertinant de montrer 'pommes' que 'pommier' puisque le mots 'pommes' à moins de différance avec se qui à été rechercher que le mot 'pommier').

TrouverUnFichierParID.go :
- TrouverUnElementParID_ArtisteS() : Renvoie la structure ArtisteS qui possède le bon ID.
- TrouverUnElementParID_LocationsS() : Renvoie la structure LocationsS qui possède le bon ID.
- TrouverUnElementParID_DatesS() : Renvoie la structure DateS qui possède le bon ID.
- TrouverUnElementParID_RelationS() : Renvoie la structure RelationS qui possède le bon ID.
- TrouverUnElementParID_ArtisteS() : Renvoie la structure artiste qui possède le bon ID.
- PlacerUnePage() : Il s'agit de la fonction appeller à chaque fois que l'on veut changer la page du site (il s'agit toujours de la dernière fonction appeler dans l'opération de changement/raffraichisement de page.)

main.go :
- main() : Il s'agit de la fonction qui allume le serveur local du site. C'est cette fonction qui vas charger les fichier CSS et géré l'appuie sur les bouton des pages HTML.
- FonctionRecherche() : Cette fonction est appeller quand un utilisateur réalise une recherche, elle vas déclancher tout le processus de recherche.
- FonctionCliqueBoutonDeNavigation() : Cette fonction vas rediriger l'utilisateur vers la page d'aceuil ou la page d'information en fonction de se qu'il a choisie dans la barre de navigation.
- FonctionInformationsAppelle() : Cette fonction est appeller quand un utilisateur shouaite accédé à la page d'information d'un groupe en particulier.

## • Utilisation.
1 - Ouvrez un éditeur de code.
2 - Allez dans le dossier `groupie-tracker`.
3 - Écrire la commande `go run .`.

## • Réaliser par :
- Jérémie.
- Émerick.
- Paul.
