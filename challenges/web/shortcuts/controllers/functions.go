package controllers

import (
	"bytes"
	"encoding/csv"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"

	"../models"
	"github.com/julienschmidt/httprouter"
	uuid "github.com/satori/go.uuid"
)

type Short struct{}

func New() *Short {
	return &Short{}
}

func (cc *Short) Index(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	cookie, err := req.Cookie("session")
	if err != nil {
		id, err := uuid.NewV4()
		if err != nil {
			log.Fatalln(err)
		}
		cookie = &http.Cookie{
			Name:     "session",
			Value:    id.String(),
			HttpOnly: true,
		}
		//600 sec == 10 mintues
		cookie.MaxAge = 600
		http.SetCookie(w, cookie)
		models.DBSessions[id.String()] = models.Session{id.String(), time.Now()}

		models.Lists = models.Original
		shortct := "shortcuts" + id.String()
		cmd := exec.Command("cp", "shortcuts/", shortct, "-r")
		err = cmd.Run()
		models.Dir = shortct
		fmt.Println(models.Dir)
		if err != nil {
			fmt.Println(err)
		}
	} else {
		models.Dir = "shortcut" + cookie.Value
	}
	err = models.Shortcuts.ExecuteTemplate(w, "index.html", nil)
	if err != nil {
		http.Redirect(w, req, "/list", http.StatusTemporaryRedirect)
	}

}

func (cc *Short) Upload(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	c, err := req.Cookie("session")
	if err != nil {
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}
	if req.Method == http.MethodPost {
		//Open
		f, h, err := req.FormFile("file")
		handleError(w, err)
		defer f.Close()

		//Read bytes
		bs, err := ioutil.ReadAll(f)
		handleError(w, err)

		//Write to file
		name := c.Value
		path := "shortcuts" + name + "/"
		dst, err := os.Create(filepath.Join(path, h.Filename))
		handleError(w, err)
		defer dst.Close()

		_, err = dst.Write(bs)
		handleError(w, err)
		models.Lists = append(models.Lists, h.Filename)
		http.Redirect(w, req, "/list", http.StatusSeeOther)
	}
	w.Header().Set("Content-Type", "text/html;charset=utf-8")
	io.WriteString(w, `
	<form method="POST" enctype="multipart/form-data" action="/upload">
	<input type="file" name="file">
	<input type="submit">
	</form>
	<br>`)

}

func csvWriter(data []string) {
	f, err := os.OpenFile(".list.csv", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
	w := csv.NewWriter(f)
	w.Write(data)
	w.Flush()
}

//Short for shortcuts
func (cc *Short) Short(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	c, err := req.Cookie("session")
	if err != nil {
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}
	name := req.FormValue("alias")
	dir := models.Dir
	if checkName(name, models.Lists) {
		dir = c.Value
		alias := "shortcuts" + dir + "/" + name + ".go"
		fmt.Println(alias)
		output := exec.Command("go", "run", alias)
		var out bytes.Buffer
		output.Stdout = &out
		err := output.Run()
		if err != nil {
			fmt.Fprint(w, "Sorry there doesn't seem to be a "+name+".go file")
			return
		}

		data := struct {
			Name   string
			Output string
		}{
			req.FormValue("alias"),
			replace(out.String()),
		}
		err = models.Shortcuts.ExecuteTemplate(w, "shortcut.gohtml", data)
		handleError(w, err)
	} else {
		fmt.Print(dir)
		http.Error(w, err.Error(), 404)
		handleError(w, err)
	}

}

func (cc *Short) Show(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	_, err := req.Cookie("session")
	if err != nil {
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}
	err = models.Shortcuts.ExecuteTemplate(w, "list.gohtml", models.Lists)
	handleError(w, err)
}

func handleError(w http.ResponseWriter, err error) {
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func checkName(name string, lines []string) bool {
	for _, v := range lines {
		if strings.Compare(name, v) == 0 {
			fmt.Println("true")
			return true
		}
	}
	return false
}
func replace(s string) string {
	s = strings.ReplaceAll(s, "\n", "<br>")
	s = strings.ReplaceAll(s, `\n`, "<br>")
	return s
}
