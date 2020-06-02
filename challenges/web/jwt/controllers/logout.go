package controllers

import (
	"net/http"

	"../models"
)

func (uc Controller) Logout(w http.ResponseWriter, req *http.Request) {
	if !alreadyLoggedIn(req) {
		http.Redirect(w, req, "https://localhost:8443/", http.StatusSeeOther)
		return
	}
	c, _ := req.Cookie("token")

	//Delete token
	delete(models.DBSessions, c.Value)

	//Remove cookie
	c = &http.Cookie{
		Name:   "token",
		Value:  "",
		MaxAge: -1, //Age = -1 just deletes the cookie i think
	}
	http.SetCookie(w, c)
	http.Redirect(w, req, "/login", http.StatusSeeOther)
}
