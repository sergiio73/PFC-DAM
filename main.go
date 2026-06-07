package main

import (
	"log"
	"net/http"

	"rolapp/controllers"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", controllers.Home)

	mux.HandleFunc("/login", controllers.LoginPage)
	mux.HandleFunc("/logout", controllers.Logout)

	mux.HandleFunc("/personajes", controllers.CharactersPage)
	mux.HandleFunc("/personajes/crear", controllers.CreateCharacter)
	mux.HandleFunc("/personajes/eliminar", controllers.DeleteCharacter)

	mux.HandleFunc("/partidas", controllers.GamesPage)
	mux.HandleFunc("/partidas/crear", controllers.CreateGame)
	mux.HandleFunc("/partidas/ver", controllers.GameDetail)

	mux.HandleFunc("/inscripciones/crear", controllers.CreateRegistration)
	mux.HandleFunc("/inscripciones/aceptar", controllers.AcceptRegistration)
	mux.HandleFunc("/inscripciones/rechazar", controllers.RejectRegistration)

	mux.HandleFunc("/guia", controllers.GuidePage)

	log.Println("Servidor iniciado en http://localhost:8080")

	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		log.Fatal("Error arrancando servidor:", err)
	}
}
