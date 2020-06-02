package controllers

import (
	"log"
	"strconv"
	"strings"

	"../models"
)

//PopulateDB create users and cars tables
func (cc Controller) PopulateDB(car []models.Cars) {
	//Check if table Users is empty
	if empty("Users") {
		query := `create table if not exists Users (
					Username varchar(25),
					Password varchar(50));`
		_, err := models.Database.Exec(query)
		if err != nil {
			log.Fatal(err)
			return
		}
		defer models.Database.Close()

		//Insert into Table Users
		for key, value := range models.DBUsers {
			query := `insert into Users values (` + key + `,` + value + `)`
			_, err := models.Database.Exec("use cars;")
			_, err = models.Database.Exec(query)
			if err != nil {
				log.Fatal(err)
				return
			}
			defer models.Database.Close()
		}
	}
	//Check if table Cars is emtpy
	if empty("Cars") {
		//Create Table Cars
		query := `create table if not exists Cars (
		Id int,
		Type varchar(25),
		Model varchar(25),
		Make varchar(25),
		Year int);`
		_, err := models.Database.Exec("use cars;")
		_, err = models.Database.Exec(query)
		if err != nil {
			log.Fatal(err)
			return
		}
		defer models.Database.Close()

		//Insert into Table Cars
		for _, v := range car {
			id := 0
			if strings.Contains(v.Type, "Sport") {
				id = 1
			}
			if strings.Contains(v.Type, "Sedan") {
				id = 2
			}
			if strings.Contains(v.Type, "Minivan") {
				id = 3
			}
			query := `insert into Cars values (` + strconv.Itoa(id) + `,` + v.Type + `,` + v.Model + `,` + v.Make + `,` + strconv.Itoa(v.Year) + `)`
			_, err = models.Database.Exec("use cars;")
			_, err := models.Database.Query(query)
			if err != nil {
				log.Fatal(err)
				return
			}
			defer models.Database.Close()
		}
		return
	}
}

func empty(object string) bool {
	query := "select * from " + object + ";"

	rows, err := models.Database.Query(query)
	if err != nil {
		if strings.Contains(err.Error(), "Table 'cars.Users' doesn't exist") {
			return true
		}
		if strings.Contains(err.Error(), "Table 'cars.Cars' doesn't exist") {
			return true
		}
	}
	defer rows.Close()
	for rows.Next() {
		return false
	}
	return true
}
