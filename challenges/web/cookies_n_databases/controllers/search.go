package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"../models"
	"github.com/julienschmidt/httprouter"
)

type data struct {
	Id    string
	Type  string
	Model string
	Make  string
	Year  string
}

//Search database for given value
func (cc Controller) Search(w http.ResponseWriter, req *http.Request, ps httprouter.Params) {
	c, _ := req.Cookie("client")
	count, _ := strconv.Atoi(c.Value)
	if checkCookie(count-35, w) {
		id := req.FormValue("id")
		fmt.Println(id)
		_, err := models.Database.Exec("use cars;")
		query := "select * from Cars where Id = " + id + ";"
		fmt.Println(query)
		rows, err := models.Database.Query(query)
		if err != nil {
			http.Error(w, err.Error(), 404)
			return
		}
		result := make([]struct {
			Id    string
			Type  string
			Model string
			Make  string
			Year  string
		}, 0)

		for rows.Next() {
			var data data
			rows.Scan(&data.Id, &data.Type, &data.Model, &data.Make, &data.Year)
			result = append(result, data)
		}
		err = models.Tpl.ExecuteTemplate(w, "show.gohtml", result)
		if handleError(w, err) {
			return
		}
	} else {
		//Sorry page, "You must be the 3123213 visit,whereas you are the count-35 visitor...so close"
		err := models.Tpl.ExecuteTemplate(w, "sorry.html", count-35)
		if handleError(w, err) {
			return
		}
		return
	}

}
