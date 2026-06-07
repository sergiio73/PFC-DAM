package controllers

import (
	"html/template"
	"log"
	"net/http"
)

// Función auxiliar para leer el usuario de la cookie en cualquier handler.
func getCurrentUser(r *http.Request) string {
	cookie, err := r.Cookie("rolapp_user")
	if err != nil {
		return ""
	}
	return cookie.Value
}

func render(w http.ResponseWriter, page string, data interface{}) {
	files := []string{
		"views/layout.html",
		"views/" + page,
	}

	tmpl, err := template.ParseFiles(files...)
	if err != nil {
		log.Println("Error cargando plantilla:", err)
		http.Error(w, "Error cargando la página", http.StatusInternalServerError)
		return
	}

	err = tmpl.ExecuteTemplate(w, "layout", data)
	if err != nil {
		log.Println("Error ejecutando plantilla:", err)
		http.Error(w, "Error mostrando la página", http.StatusInternalServerError)
		return
	}
}
