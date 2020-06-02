package main

import (
	"database/sql"
	"log"
	"net/http"
	"text/template"

	"./controllers"
	"./models"
	_ "github.com/go-sql-driver/mysql"
	"github.com/julienschmidt/httprouter"
)

func init() {
	//Initialize templates
	models.Tpl = template.Must(template.ParseGlob("models/templates/*"))
	//Intialize database
	models.Database, _ = sql.Open("mysql", "root:passw0rd@tcp(db:3306)/")
	//Create Database
	_, err := models.Database.Exec("create database if not exists cars;")
	if err != nil {
		log.Fatalln(err)
	}
	_, err = models.Database.Exec("use cars;")
	if err != nil {
		log.Fatalln(err)
		return
	}
}

func main() {
	mux := httprouter.New()
	cc := controllers.NewController()
	//Populate
	cars := cc.CreateCars()
	cc.PopulateDB(cars)

	mux.POST("/search", cc.Search)
	mux.GET("/no", cc.No)
	mux.GET("/yes", cc.Yes)
	mux.GET("/dealer", cc.Dealer)
	mux.POST("/dealer", cc.Dealer)
	mux.GET("/", cc.Welcome)
	http.Handle("/favicon.ico", http.NotFoundHandler())

	http.ListenAndServe(":8080", mux)
}
