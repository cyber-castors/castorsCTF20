package main

import (
	"net/http"
	"text/template"

	"./certificates"
	"./controllers"
	"./models"
)

func init() {
	models.Tpl = template.Must(template.ParseGlob("models/templates/*"))
	certificates.CreateCert()
}

func main() {
	uc := controllers.NewController()
	// mux := httprouter.New()
	http.HandleFunc("/dmv", uc.DMV)
	http.HandleFunc("/", uc.Index)
	http.HandleFunc("/logout", uc.Logout)
	http.HandleFunc("/signup", uc.Register)
	//mux.POST("/signup", uc.Register)
	//http.HandleFunc("/login", uc.Login)
	http.HandleFunc("/login", uc.Login)
	http.HandleFunc("/refresh", uc.Refresh)
	http.HandleFunc("/welcome", uc.Success)

	http.ListenAndServeTLS(":8443", "./certificates/cert.pem", "./certificates/key.pem", nil)
}
