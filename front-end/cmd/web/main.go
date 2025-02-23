package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
)

const webPort = "3000"

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		render(w, "test.page.gohtml")
	})

	fmt.Printf("Starting front end service on port %s\n", webPort)
	err := http.ListenAndServe(fmt.Sprintf(":%s", webPort), nil)
	if err != nil {
		log.Panic(err)
	}
}

func render(w http.ResponseWriter, t string) {

	basePath, _ := filepath.Abs("./cmd/web/templates/")

	partials := []string{
		filepath.Join(basePath, "base.layout.gohtml"),
		filepath.Join(basePath, "header.partial.gohtml"),
		filepath.Join(basePath, "footer.partial.gohtml"),
	}
	templateSlice := []string{
		filepath.Join(basePath, t),
	}
	templateSlice = append(templateSlice, partials...)

	tmpl, err := template.ParseFiles(templateSlice...)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := tmpl.Execute(w, nil); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
