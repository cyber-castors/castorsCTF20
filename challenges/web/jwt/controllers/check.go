package controllers

import (
	"fmt"
	"net/http"

	"../models"
)

func alreadyLoggedIn(req *http.Request) bool {
	c, err := req.Cookie("token")
	if err != nil {
		return false
	}
	un := models.DBSessions[c.Value]
	fmt.Println(un, c.Value)
	_, ok := models.DBUsers[un]
	return ok
}

func getUser(req *http.Request) models.User {
	var u models.User
	c, err := req.Cookie("token")
	if err != nil {
		return u
	}

	un, ok := models.DBSessions[c.Value]
	if ok {
		u = models.DBUsers[un]
	}
	return u
}

func handleError(w http.ResponseWriter, err error) bool {
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return false
	}
	return true
}
