package controllers

import (
	"log"
	"net/http"

	"rolapp/models"
)

type HomeData struct {
	Games       []models.Game
	CurrentUser string
}

func Home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	currentUser := getCurrentUser(r)
	if currentUser == "" {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}

	games, err := models.GetAllGames()
	if err != nil {
		log.Println("Error leyendo partidas:", err)
	}

	data := HomeData{
		Games:       games,
		CurrentUser: currentUser,
	}

	render(w, "home.html", data)
}
