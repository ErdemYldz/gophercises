package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type application struct {
	router *http.ServeMux
	data   chapters
}

func (app *application) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	app.router.ServeHTTP(w, r)
}

func main() {
	addr := flag.String("addr", ":4040", "the port address")
	fn := flag.String("fn", "gopher.json", "the story file to open")
	flag.Parse()

	chaps, err := loadChapters(*fn)
	if err != nil {
		log.Fatalln(err)
	}

	app := newApplication()
	app.data = chaps

	srv := &http.Server{
		Addr:    *addr,
		Handler: app,
	}

	log.Printf("starting server on %s\n", *addr)
	err = srv.ListenAndServe()
	log.Fatalln(err)

}

func newApplication() *application {
	app := &application{
		router: http.NewServeMux(),
	}
	app.routes()
	return app
}
func loadChapters(filename string) (chapters, error) {
	f, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("error while opening the file: %s ", err)
	}
	var chaps chapters
	err = json.Unmarshal(f, &chaps)
	if err != nil {
		return nil, fmt.Errorf("error while unmarshaling: %s", err)
	}
	return chaps, nil
}
