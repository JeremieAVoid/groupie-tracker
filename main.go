package main

import (
	"fmt"
	"groupie"
)

func main() {
	fmt.Println("début")
	lotDeListe := groupie.ChargerLesDonnées()
	groupie.Trie(lotDeListe, "Name")
}
