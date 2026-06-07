package controllers

import (
	"log"
	"net/http"
	"strconv"

	"rolapp/models"
)

type GamesData struct {
	Games       []models.Game
	CurrentUser string
}

type GameDetailData struct {
	Game          models.Game
	Characters    []models.Character
	Registrations []models.RegistrationView
	Message       string
	IsUserMaster  bool
	AcceptedCount int
	CurrentUser   string
}

func GamesPage(w http.ResponseWriter, r *http.Request) {
	games, err := models.GetAllGames()
	if err != nil {
		log.Println("Error obteniendo partidas:", err)
	}

	data := GamesData{
		Games:       games,
		CurrentUser: getCurrentUser(r),
	}

	render(w, "games.html", data)
}

func CreateGame(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Redirect(w, r, "/partidas", http.StatusSeeOther)
		return
	}

	currentUser := getCurrentUser(r)
	if currentUser == "" {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}

	maxPlayers, err := strconv.Atoi(r.FormValue("max_players"))
	if err != nil || maxPlayers < 1 {
		maxPlayers = 5
	}

	game := models.Game{
		Title:       r.FormValue("title"),
		Description: r.FormValue("description"),
		Date:        r.FormValue("date"),
		MasterName:  currentUser,
		MaxPlayers:  maxPlayers,
		Location:    r.FormValue("location"),
	}

	err = models.CreateGame(game)
	if err != nil {
		log.Println("Error creando partida:", err)
	}

	http.Redirect(w, r, "/partidas", http.StatusSeeOther)
}

func GameDetail(w http.ResponseWriter, r *http.Request) {
	gameID, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		log.Println("ID de partida incorrecto:", err)
		http.Redirect(w, r, "/partidas", http.StatusSeeOther)
		return
	}

	game, err := models.GetGameByID(gameID)
	if err != nil {
		log.Println("Error buscando partida:", err)
		http.Redirect(w, r, "/partidas", http.StatusSeeOther)
		return
	}

	currentUser := getCurrentUser(r)

	// Solo le muestro al usuario sus propios personajes en el desplegable de inscripción.
	var characters []models.Character
	if currentUser != "" {
		characters, err = models.GetCharactersByOwner(currentUser)
		if err != nil {
			log.Println("Error obteniendo personajes:", err)
		}
	}

	registrations, err := models.GetRegistrationsByGame(gameID)
	if err != nil {
		log.Println("Error obteniendo inscripciones:", err)
	}

	acceptedCount, err := models.CountAcceptedPlayers(gameID)
	if err != nil {
		log.Println("Error contando aceptados:", err)
	}

	data := GameDetailData{
		Game:          game,
		Characters:    characters,
		Registrations: registrations,
		IsUserMaster:  currentUser != "" && currentUser == game.MasterName,
		AcceptedCount: acceptedCount,
		CurrentUser:   currentUser,
	}

	render(w, "game_detail.html", data)
}

func CreateRegistration(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Redirect(w, r, "/partidas", http.StatusSeeOther)
		return
	}

	currentUser := getCurrentUser(r)
	if currentUser == "" {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}

	gameID, err := strconv.Atoi(r.FormValue("game_id"))
	if err != nil {
		log.Println("ID partida mal:", err)
		http.Redirect(w, r, "/partidas", http.StatusSeeOther)
		return
	}

	characterID, err := strconv.Atoi(r.FormValue("character_id"))
	if err != nil {
		log.Println("ID personaje mal:", err)
		http.Redirect(w, r, "/partidas/ver?id="+strconv.Itoa(gameID), http.StatusSeeOther)
		return
	}

	registration := models.Registration{
		GameID:      gameID,
		CharacterID: characterID,
		PlayerName:  currentUser, // el nombre viene de la cookie, no del formulario
		Status:      "pendiente",
	}

	err = models.CreateRegistration(registration)
	if err != nil {
		log.Println("Error creando inscripción:", err)
	}

	http.Redirect(w, r, "/partidas/ver?id="+strconv.Itoa(gameID), http.StatusSeeOther)
}

func AcceptRegistration(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Redirect(w, r, "/partidas", http.StatusSeeOther)
		return
	}

	registrationID, err := strconv.Atoi(r.FormValue("registration_id"))
	if err != nil {
		log.Println("ID inscripción mal:", err)
		http.Redirect(w, r, "/partidas", http.StatusSeeOther)
		return
	}

	gameID, err := strconv.Atoi(r.FormValue("game_id"))
	if err != nil {
		log.Println("ID partida mal:", err)
		http.Redirect(w, r, "/partidas", http.StatusSeeOther)
		return
	}

	game, err := models.GetGameByID(gameID)
	if err != nil {
		log.Println("Error buscando partida:", err)
		http.Redirect(w, r, "/partidas", http.StatusSeeOther)
		return
	}

	// Verificación server-side aunque los botones ya estén ocultos en la vista.
	currentUser := getCurrentUser(r)
	if currentUser != game.MasterName {
		http.Redirect(w, r, "/partidas/ver?id="+strconv.Itoa(gameID), http.StatusSeeOther)
		return
	}

	count, err := models.CountAcceptedPlayers(gameID)
	if err != nil {
		log.Println("Error contando aceptados:", err)
	}

	if count >= game.MaxPlayers {
		log.Println("Partida llena, no se puede aceptar más.")
		http.Redirect(w, r, "/partidas/ver?id="+strconv.Itoa(gameID), http.StatusSeeOther)
		return
	}

	err = models.UpdateRegistrationStatus(registrationID, "aceptado")
	if err != nil {
		log.Println("Error aceptando inscripción:", err)
	}

	http.Redirect(w, r, "/partidas/ver?id="+strconv.Itoa(gameID), http.StatusSeeOther)
}

func RejectRegistration(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Redirect(w, r, "/partidas", http.StatusSeeOther)
		return
	}

	registrationID, err := strconv.Atoi(r.FormValue("registration_id"))
	if err != nil {
		log.Println("ID inscripción mal:", err)
		http.Redirect(w, r, "/partidas", http.StatusSeeOther)
		return
	}

	gameID, err := strconv.Atoi(r.FormValue("game_id"))
	if err != nil {
		log.Println("ID partida mal:", err)
		http.Redirect(w, r, "/partidas", http.StatusSeeOther)
		return
	}

	game, err := models.GetGameByID(gameID)
	if err != nil {
		log.Println("Error buscando partida:", err)
		http.Redirect(w, r, "/partidas", http.StatusSeeOther)
		return
	}

	currentUser := getCurrentUser(r)
	if currentUser != game.MasterName {
		http.Redirect(w, r, "/partidas/ver?id="+strconv.Itoa(gameID), http.StatusSeeOther)
		return
	}

	err = models.UpdateRegistrationStatus(registrationID, "rechazado")
	if err != nil {
		log.Println("Error rechazando inscripción:", err)
	}

	http.Redirect(w, r, "/partidas/ver?id="+strconv.Itoa(gameID), http.StatusSeeOther)
}
