package controllers

import (
	"log"
	"net/http"
	"strconv"

	"rolapp/models"
)

type CharactersData struct {
	Characters  []models.Character
	Message     string
	CurrentUser string
}

func CharactersPage(w http.ResponseWriter, r *http.Request) {
	currentUser := getCurrentUser(r)

	var characters []models.Character
	var err error

	if currentUser != "" {
		characters, err = models.GetCharactersByOwner(currentUser)
		if err != nil {
			log.Println("Error obteniendo personajes:", err)
		}
	}

	data := CharactersData{
		Characters:  characters,
		CurrentUser: currentUser,
	}

	render(w, "characters.html", data)
}

func CreateCharacter(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Redirect(w, r, "/personajes", http.StatusSeeOther)
		return
	}

	currentUser := getCurrentUser(r)
	if currentUser == "" {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}

	level, err := strconv.Atoi(r.FormValue("level"))
	if err != nil {
		log.Println("Nivel incorrecto:", err)
		level = 1
	}

	character := models.Character{
		Name:    r.FormValue("name"),
		Class:   r.FormValue("class"),
		Race:    r.FormValue("race"),
		Level:   level,
		History: r.FormValue("history"),
		Owner:   currentUser,
	}

	err = models.CreateCharacter(character)
	if err != nil {
		log.Println("Error creando personaje:", err)
	}

	http.Redirect(w, r, "/personajes", http.StatusSeeOther)
}

func DeleteCharacter(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Redirect(w, r, "/personajes", http.StatusSeeOther)
		return
	}

	currentUser := getCurrentUser(r)
	if currentUser == "" {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}

	characterID, err := strconv.Atoi(r.FormValue("character_id"))
	if err != nil {
		log.Println("ID personaje mal:", err)
		http.Redirect(w, r, "/personajes", http.StatusSeeOther)
		return
	}

	err = models.DeleteCharacter(characterID, currentUser)
	if err != nil {
		log.Println("Error eliminando personaje:", err)
	}

	http.Redirect(w, r, "/personajes", http.StatusSeeOther)
}
