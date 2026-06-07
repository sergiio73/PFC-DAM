package controllers

import (
	"net/http"
	"strings"
)

type LoginData struct {
	CurrentUser string
}

func LoginPage(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		username := strings.TrimSpace(r.FormValue("username"))
		if username == "" {
			http.Redirect(w, r, "/login", http.StatusSeeOther)
			return
		}

		http.SetCookie(w, &http.Cookie{
			Name:  "rolapp_user",
			Value: username,
			Path:  "/",
		})
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	render(w, "login.html", LoginData{})
}

func Logout(w http.ResponseWriter, r *http.Request) {
	http.SetCookie(w, &http.Cookie{
		Name:   "rolapp_user",
		Value:  "",
		Path:   "/",
		MaxAge: -1,
	})
	http.Redirect(w, r, "/", http.StatusSeeOther)
}
