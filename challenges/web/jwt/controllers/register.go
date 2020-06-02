package controllers

import (
	"net/http"
	"time"

	"../models"
	"github.com/dgrijalva/jwt-go"
)

func (uc Controller) Register(w http.ResponseWriter, req *http.Request) {
	if alreadyLoggedIn(req) {
		http.Redirect(w, req, "https://localhost:443/", http.StatusSeeOther)
		return
	}
	if req.Method == http.MethodPost {
		un := req.FormValue("username")
		passwd := req.FormValue("password")
		fname := req.FormValue("firstname")
		lname := req.FormValue("lastname")

		//is username taken?
		if _, ok := models.DBUsers[un]; ok {
			http.Error(w, "Username already taken", http.StatusForbidden)
			return
		}

		//Create session with JWT token
		//Declare Expiration time for tokens
		expTime := time.Now().Add(5 * time.Minute)

		//Create Vulnerable jwt un purpose
		if un == "admin@cybercastors.com" {
			//Create JWT claim
			claims := &models.Claims{
				Username:       un,
				StandardClaims: jwt.StandardClaims{},
			}
			//Declare token
			token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
			//JWT token string
			tokenString, err := token.SignedString(models.JwtKey)
			if !handleError(w, err) {
				return
			}

			//Set Cookie
			c := &http.Cookie{
				Name:  "token",
				Value: tokenString,
			}

			http.SetCookie(w, c)
			models.DBSessions[c.Value] = un

		} else {
			//Create JWT claim
			claims := &models.Claims{
				Username: un,
				StandardClaims: jwt.StandardClaims{
					ExpiresAt: expTime.Unix(),
				},
			}
			//Declare token
			token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
			//JWT token string
			tokenString, err := token.SignedString(models.JwtKey)
			if !handleError(w, err) {
				return
			}

			//Set Cookie
			c := &http.Cookie{
				Name:    "token",
				Value:   tokenString,
				Expires: expTime,
			}
			http.SetCookie(w, c)
			models.DBSessions[c.Value] = un

		}

		//Store user in DB
		u := models.User{un, passwd, fname, lname}
		models.DBUsers[un] = u

		//redirect
		http.Redirect(w, req, "https://localhost:8443/", http.StatusSeeOther)
		return

	}
	models.Tpl.ExecuteTemplate(w, "signup.gohtml", nil)
}
