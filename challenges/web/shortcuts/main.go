package main

import (
	"encoding/csv"
	"html/template"
	"log"
	"net/http"
	"os"

	"./controllers"
	"./models"
	"github.com/julienschmidt/httprouter"
)

func init() {
	models.Shortcuts = template.Must(template.ParseGlob("templates/*"))

	lst, err := os.Open("controllers/.list.csv")
	if err != nil {
		log.Fatalln(err)
	}
	data, err := csv.NewReader(lst).ReadAll()

	for _, v := range data {
		models.Lists = append(models.Lists, v[0])
	}
	models.Original = models.Lists
	go controllers.CleanSession()

}

func main() {
	mux := httprouter.New()
	sc := controllers.New()

	mux.GET("/", sc.Index)
	mux.GET("/list", sc.Show)
	mux.GET("/shortcuts", sc.Short)
	mux.POST("/upload", sc.Upload)
	mux.GET("/upload", sc.Upload)
	http.ListenAndServe(":8080", mux)
}
