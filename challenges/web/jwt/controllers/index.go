package controllers

import (
	"net/http"

	"../models"
)

func (uc Controller) Index(w http.ResponseWriter, req *http.Request) {
	u := getUser(req)
	if !alreadyLoggedIn(req) {
		http.Redirect(w, req, "https://localhost:8443/login", http.StatusSeeOther)
	}
	models.Tpl.ExecuteTemplate(w, "index.gohtml", u)
}

func (uc Controller) DMV(w http.ResponseWriter, req *http.Request) {
	u := getUser(req)
	models.Tpl.ExecuteTemplate(w, "lounge.gothml", u)
}
