package main

import (
	"flag"
	"log"
	"net/http"
)

type storage interface {
	getData(string) (string, bool)
}

// application holds dependencies
type application struct {
	router *http.ServeMux
	storage
}

// pass the execution to the router
func (app *application) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	app.router.ServeHTTP(w, r)
}

func main() {
	addr := flag.String("addr", ":4000", "the port to listen")
	stype := flag.String("stype", "go-map", "storage type")
	flag.Parse()

	app := newApplication()
	var err error
	switch *stype {
	case "go-map":
		app.storage, err = newMap()
		if err != nil {
			log.Fatalln("case go-map: ", err)
		}
	case "go-yaml":
		app.storage, err = newYaml()
		if err != nil {
			log.Fatalln("case go-yaml: ", err)
		}
	case "go-mysql":
		app.storage, err = newMysql()
		if err != nil {
			log.Fatalln(err)
		}
	}

	srv := &http.Server{
		Addr:    *addr,
		Handler: app,
	}
	log.Println("starting server on:", *addr)
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
