package controllers

import "../models"

//CreateCars creates slice of cars to populate DB
func (cc Controller) CreateCars() []models.Cars {
	cars := []models.Cars{
		{
			Type:  `"Sport"`,
			Model: `"R8"`,
			Make:  `"Audi"`,
			Year:  2020,
		},
		{
			Type:  `"Sport"`,
			Model: `"Portofino"`,
			Make:  `"Ferrari"`,
			Year:  2018,
		},
		{
			Type:  `"Sedan"`,
			Model: `"Genesis"`,
			Make:  `"Hyundai"`,
			Year:  2019,
		},
		{
			Type:  `"Sedan"`,
			Model: `"Focus"`,
			Make:  `"Ford"`,
			Year:  2015,
		},
		{
			Type:  `"Sport"`,
			Model: `"Aventador"`,
			Make:  `"Lamborghini"`,
			Year:  2016,
		},
		{
			Type:  `"Minivan"`,
			Model: `"Odyssey"`,
			Make:  `"Honda"`,
			Year:  2020,
		},
		{
			Type:  `"Sport"`,
			Model: `"Chiron"`,
			Make:  `"Bugatti"`,
			Year:  2019,
		},
		{
			Type:  `"Sedan"`,
			Model: `"Accord"`,
			Make:  `"Honda"`,
			Year:  2020,
		},
		{
			Type:  `"Sedan"`,
			Model: `"Civic Si"`,
			Make:  `"Honda"`,
			Year:  2016,
		},
		{
			Type:  `"Sedan"`,
			Model: `"Yaris"`,
			Make:  `"Toyota"`,
			Year:  2020,
		},
		{
			Type:  `"Minivan"`,
			Model: `"Pacifica"`,
			Make:  `"Chrysler"`,
			Year:  2018,
		},
		{
			Type:  `"Minivan"`,
			Model: `"Sienna"`,
			Make:  `"Toyota"`,
			Year:  2017,
		},
		{
			Type:  `"Minivan"`,
			Model: `"Cargo"`,
			Make:  `"Nissan"`,
			Year:  2020,
		},
		{
			Type:  `"Minivan"`,
			Model: `"Sprinter 2500 Crew"`,
			Make:  `"Mercedes-Benz"`,
			Year:  2099,
		},
	}

	return cars

}
