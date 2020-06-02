package models

import (
	"database/sql"
	"text/template"
)

//Cars struct to hold value from database
type Cars struct {
	Type  string `json:"type"` //Sedan or two door or bus
	Model string `json:"model"`
	Make  string `json:"make"`
	Year  int    `json:"year"`
}

//Tpl template to parse html files
var Tpl *template.Template

//Database variable to accesable from anywhere
var Database *sql.DB
