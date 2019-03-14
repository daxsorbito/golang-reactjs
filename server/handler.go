package server

import (
	"fmt"
	"net/http"
	"path"

	"github.com/alecthomas/template"
)

// Handler return http.Hander for server endpoint
func Handler(buildPath string) http.HandlerFunc {
	tmpl, err := template.ParseFiles(path.Join("templates", "index.html"))

	if err != nil {
		return func(res http.ResponseWriter, req *http.Request) {
			http.Error(res, err.Error(), http.StatusInternalServerError)
		}
	}

	data, err := NewViewData(buildPath)
	fmt.Println("data>>", data)

	if err != nil {
		return func(res http.ResponseWriter, req *http.Request) {
			http.Error(res, err.Error(), http.StatusInternalServerError)
		}
	}

	return func(res http.ResponseWriter, req *http.Request) {
		if err := tmpl.Execute(res, data); err != nil {
			http.Error(res, err.Error(), http.StatusInternalServerError)
		}
	}
}
