package controllers

import (
	"fmt"
	"io"
	"net/http"

	"../models"
)

type Controller struct{}

func NewController() *Controller {
	return &Controller{}
}

func (cc Controller) Index(w http.ResponseWriter, req *http.Request) {
	err := models.TPL.ExecuteTemplate(w, "index.html", nil)
	if err != nil {
		io.WriteString(w, err.Error())
	}

	http.Redirect(w, req, "/rubix", http.StatusTemporaryRedirect)

}

func (cc Controller) Rubix(w http.ResponseWriter, req *http.Request) {
	err := models.TPL.ExecuteTemplate(w, "rubics.html", nil)
	if err != nil {
		io.WriteString(w, err.Error())
	}
}

func (cc Controller) Flag(w http.ResponseWriter, req *http.Request) {
	if req.Method == http.MethodPost {
		one := req.FormValue("cookies")
		two := req.FormValue("puppies")

		fmt.Println(req)

		decide(one, w, req)
		decide(two, w, req)
		return
	}
	models.TPL.ExecuteTemplate(w, ".flagkindsir.html", nil)
}

func decide(str string, w http.ResponseWriter, req *http.Request) {
	if len(str) == 0 {
		return
	}
	switch str {
	case "flag":
		getFlag(w)
	case "cookies":
		getCookie(w, req)
	case "puppies":
		getPuppies(w, req)

	}
}

func getFlag(w http.ResponseWriter) {
	fmt.Print(models.Flag)
	err := models.TPL.ExecuteTemplate(w, ".flagkindsir.html", models.Flag)
	if err != nil {
		io.WriteString(w, err.Error())
	}
}

func getCookie(w http.ResponseWriter, req *http.Request) {
	err := models.TPL.ExecuteTemplate(w, ".flagkindsir.html", "https://bit.ly/2AjkbEZ")
	if err != nil {
		io.WriteString(w, err.Error())
	}
}

func getPuppies(w http.ResponseWriter, req *http.Request) {
	err := models.TPL.ExecuteTemplate(w, ".flagkindsir.html", "https://bit.ly/3ckaEuw")
	if err != nil {
		io.WriteString(w, err.Error())
	}
}
