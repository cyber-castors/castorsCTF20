package controllers

import (
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"

	"../models"
)

func (uc Controller) Login(w http.ResponseWriter, req *http.Request) {
	if alreadyLoggedIn(req) {
		http.Redirect(w, req, "https://localhost:8443/", http.StatusSeeOther)
		return
	}
	if req.Method == http.MethodPost {
		un := req.FormValue("username")
		passwd := req.FormValue("password")

		u, ok := models.DBUsers[un]
		if !ok {
			http.Error(w, "Username/password don't match", http.StatusForbidden)
			return
		}
		if u.Password != passwd {
			http.Error(w, "Username/password don't match", http.StatusForbidden)
			return
		}

		//Create Session
		expTime := time.Now().Add(5 * time.Minute)

		claims := &models.Claims{
			Username: u.Username,
			StandardClaims: jwt.StandardClaims{
				ExpiresAt: expTime.Unix(),
			},
		}
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
		tokenString, err := token.SignedString(models.JwtKey)
		if !handleError(w, err) {
			return
		}
		c := &http.Cookie{
			Name:    "token",
			Value:   tokenString,
			Expires: expTime,
		}
		http.SetCookie(w, c)
		models.DBSessions[c.Value] = un
		http.Redirect(w, req, "https://localhost:8443/", http.StatusSeeOther)
	}
	models.Tpl.ExecuteTemplate(w, "login.gohtml", nil)

}
