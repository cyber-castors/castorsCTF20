package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"net/url"
	"os"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

var local *template.Template
var probs *template.Template

func init() {
	local = template.Must(template.ParseGlob("./files/*"))
	probs = template.Must(template.ParseGlob("./problems/*"))
}

func main() {
	mux := httprouter.New()

	mux.GET("/", index)
	mux.GET("/test/:directory/:theme/:whynot", super)
	mux.GET("/problems/math", math)
	mux.POST("/problems/math", mathCheck)

	//Remember to Delete
	mux.GET("/.backup/", backup)

	//Serve File with Directory listing
	http.ListenAndServe(":8080", mux)
}
func index(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	err := local.ExecuteTemplate(w, "start.html", nil)
	handleError(w, err)
}

func backup(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	http.ServeFile(w, req, "main.go")
}

func mathCheck(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	err := req.ParseForm()
	handleError(w, err)
	check(w, req.Form)
}
func math(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	game(w)
}
func file(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	http.FileServer(http.Dir("."))
}

func check(w http.ResponseWriter, form url.Values) {
	answers, err := os.Open("problems/answers.csv")
	handleError(w, err)

	data, _ := csv.NewReader(answers).ReadAll()

	index, err := strconv.Atoi(form["index"][0])
	handleError(w, err)
	value := form["var"][0]

	f_answers := make(map[int]string)

	for i, v := range data {
		f_answers[i+1] = v[0]
	}

	if f_answers[index] == value {
		last := struct {
			Header string
			SorC   string
		}{
			"correct!!",
			"Congrats!",
		}

		err := probs.ExecuteTemplate(w, "mathCheck.gohtml", last)
		handleError(w, err)
	} else {
		last := struct {
			Header string
			SorC   string
		}{
			"incorrect.",
			"Sorry...",
		}

		err := probs.ExecuteTemplate(w, "mathCheck.gohtml", last)
		handleError(w, err)

	}

}
func super(w http.ResponseWriter, req *http.Request, ps httprouter.Params) {
	var file string = "/" + ps.ByName("directory") + "/" + ps.ByName("theme") + "/" + ps.ByName("whynot")
	test, err := os.Open(file)
	handleError(w, err)
	defer test.Close()

	scanner := bufio.NewScanner(test)
	var content string
	for scanner.Scan() {
		content = scanner.Text()
	}

	fmt.Fprintf(w, "Directories: %s/%s\n", ps.ByName("directory"), ps.ByName("theme"))
	fmt.Fprintf(w, "File: %s\n", ps.ByName("whynot"))
	fmt.Fprintf(w, "Contents: %s\n", content)
}

func handleError(w http.ResponseWriter, err error) {
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Fatalln(err)
	}
}

func game(w http.ResponseWriter) {
	problems, err := os.Open("problems/problems.csv")
	if err != nil {
		fmt.Println(err)
	}

	data, err := csv.NewReader(problems).ReadAll()

	//Create empty struct to contain questions and their indexes
	questions := struct {
		Index    int
		Question string
	}{}
	ques := make([]struct {
		Index    int
		Question string
	}, 0)
	for i, v := range data {
		questions.Index = i + 1
		questions.Question = v[0]
		ques = append(ques, questions)
	}

	err = probs.ExecuteTemplate(w, "math.gohtml", ques)
	handleError(w, err)
}
