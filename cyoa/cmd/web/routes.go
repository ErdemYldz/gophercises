package main

func (app *application) routes() {
	app.router.HandleFunc("/", app.handleHome)

}
