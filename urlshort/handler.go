package main

import (
	"fmt"
	"net/http"
)

func (app *application) handleHome(w http.ResponseWriter, r *http.Request) {
	value, ok := app.getData(r.URL.Path)
	if !ok {
		fmt.Fprintf(w, "no short path for %s", r.URL.Path)
		return
	}
	http.Redirect(w, r, value, http.StatusPermanentRedirect)
}

// // MapHandler will return an http.HandlerFunc (which also
// // implements http.Handler) that will attempt to map any
// // paths (keys in the map) to their corresponding URL (values
// // that each key in the map points to, in string format).
// // If the path is not provided in the map, then the fallback
// // http.Handler will be called instead.
// func MapHandler(pathsToUrls map[string]string, app *application) http.HandlerFunc {
// 	return func(w http.ResponseWriter, r *http.Request) {
// 		value, ok := pathsToUrls[r.URL.Path]
// 		if !ok {
// 			app.handleErrorPage(w, r)
// 			return
// 		}
// 		http.Redirect(w, r, value, http.StatusPermanentRedirect)

// 	}
// }

// // YAMLHandler will parse the provided YAML and then return
// // an http.HandlerFunc (which also implements http.Handler)
// // that will attempt to map any paths to their corresponding
// // URL. If the path is not provided in the YAML, then the
// // fallback http.Handler will be called instead.
// func YAMLHandler(yml []byte, fallback http.Handler) (http.HandlerFunc, error) {
// 	return nil, nil
// }
