package main

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
	"text/template"

	"./controllers"
	"./models"
)

func init() {

	//Do something
	models.TPL = template.Must(template.ParseGlob("./templates/*"))

	//Open flag
	file, err := os.Open("flag.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	scnr := bufio.NewScanner(file)
	for scnr.Scan() {
		models.Flag = scnr.Text()
	}
}

func main() {
	cc := controllers.NewController()

	http.HandleFunc("/", cc.Index)
	http.HandleFunc("/rubix", cc.Rubix)
	http.HandleFunc("/.flagkindsir", cc.Flag)

	http.ListenAndServe(":8080", nil)
}
