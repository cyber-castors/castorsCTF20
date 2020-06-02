package controllers

import (
	"log"
	"net/http"

	"golang.org/x/crypto/bcrypt"

	"../models"
	"github.com/julienschmidt/httprouter"
)

//Dealer function to handle admin login
func (cc Controller) Dealer(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	if req.Method == http.MethodPost {
		un := req.FormValue("username")
		ps := req.FormValue("password")

		if un != "admin@cybercastors.com" {
			http.Error(w, "Username/password don't match", http.StatusForbidden)
			return
		}
		hash, err := bcrypt.GenerateFromPassword([]byte("naruto"), bcrypt.DefaultCost)
		if err != nil {
			log.Fatal(err)
		}

		if err := bcrypt.CompareHashAndPassword(hash, []byte(ps)); err != nil {
			http.Error(w, "Username/password don't match", http.StatusForbidden)
			return
		}
		err = models.Tpl.ExecuteTemplate(w, "home.gohtml", "castorCTF{daT4B_3n4m_1s_fuN_N_p0w3rfu7}")
		handleError(w, err)
		return

	}
	err := models.Tpl.ExecuteTemplate(w, "login.gohtml", nil)
	handleError(w, err)

}
