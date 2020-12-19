package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strings"
)

func (app *application) handleHome(w http.ResponseWriter, r *http.Request) {
	files := []string{
		"ui/html/home.page.html",
		"ui/html/base.layout.html",
	}
	ts, err := template.ParseFiles(files...)
	if err != nil {
		log.Fatalln("error while parsing htmls: ", err)
	}
	if r.URL.Path == "/" {
		err = ts.Execute(w, app.data["intro"])
		if err != nil {
			log.Fatalln("error while executing the templates: ", err)
		}
		return
	}
	p := strings.TrimPrefix(r.URL.Path, "/")
	d, ok := app.data[p]
	if !ok {
		fmt.Fprintf(w, "the chapter %s is not found", p)
		return
	}
	err = ts.Execute(w, d)
	if err != nil {
		log.Fatalln("error while executing the templates: ", err)
	}
}
