package controllers

import "net/http"

type GuideData struct {
	CurrentUser string
}

func GuidePage(w http.ResponseWriter, r *http.Request) {
	render(w, "guide.html", GuideData{CurrentUser: getCurrentUser(r)})
}
