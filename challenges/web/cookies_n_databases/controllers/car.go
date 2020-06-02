package controllers

import (
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"

	"../models"
)

//Controller allows main to access functions through polymorphism
type Controller struct{}

//NewController used to create instance over on main
func NewController() *Controller {
	return &Controller{}
}

//Welcome hadles requests to '/'
func (cc Controller) Welcome(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	cookie, err := req.Cookie("client")
	//If cookie is not found
	if err == http.ErrNoCookie {
		cookie = &http.Cookie{
			Name:  "client",
			Value: "100",
		}
		http.SetCookie(w, cookie)
		//Print out welcome page
		err := models.Tpl.ExecuteTemplate(w, "welcome.html", nil)
		if handleError(w, err) {
			return
		}
		return
	}

	count, err := strconv.Atoi(cookie.Value)

	if checkCookie(count-35, w) {
		//User can go to the car page
		err := models.Tpl.ExecuteTemplate(w, "car.gohtml", nil)
		if handleError(w, err) {
			return
		}
	} else {
		//Sorry page, "You must be the 3123213 visit,whereas you are the count-35 visitor...so close"
		count += 3
		cookie.Value = strconv.Itoa(count)
		http.SetCookie(w, cookie)
		err = models.Tpl.ExecuteTemplate(w, "sorry.html", count-35)
		if handleError(w, err) {
			return
		}
		return
	}
}

func handleError(w http.ResponseWriter, err error) bool {
	if err != nil {
		http.Error(w, err.Error(), http.StatusNoContent)
		return true
	}
	return false
}

func checkCookie(cookie int, res http.ResponseWriter) bool {
	if cookie == 3123213 {
		return true
	}
	return false
}
