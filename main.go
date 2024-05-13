package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"html/template"
	"log/slog"
	"net/http"
)

var templates map[string]*template.Template

func main() {

	err := initDB()
	if err != nil {
		panic(err)
	}

	templateDefs := []string{"home", "widgets-form", "filter-adapter"}
	templates = make(map[string]*template.Template)

	// parse html
	for _, name := range templateDefs {
		if templates[name] == nil {
			templates[name] = template.New(name)
			templates[name] = getTemplate(templates[name], "views/parts/page.html", "views/"+name+".html")
		}
	}

	// routes
	http.HandleFunc("/", Home)
	http.HandleFunc("/widget/create", WidgetCreate)
	http.HandleFunc("/widget/update/", WidgetUpdate)
	http.HandleFunc("/widgets", WidgetFilter)

	// api enpoints

	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("views/assets/"))))

	fmt.Println("********** Server Ready **********")

	err = http.ListenAndServe(":3000", nil)

	if err != nil {
		slog.Debug("error starting server", err)
		panic(err)
	}

}

func Render(w http.ResponseWriter, pageTemplate *template.Template, data any) {
	err := pageTemplate.ExecuteTemplate(w, "page", data)
	if err != nil {
		slog.Debug("error executing template", err)
		panic(err)
	}
}

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Println("in home")

	Render(w, templates["home"], nil)
}

func WidgetCreate(w http.ResponseWriter, r *http.Request) {
	Render(w, templates["widgets-form"], nil)
}

func WidgetUpdate(w http.ResponseWriter, r *http.Request) {
	// get data
	Render(w, templates["widgets-form"], nil)
}

func WidgetFilter(w http.ResponseWriter, r *http.Request) {

	template := getTemplate(templates["filter-adapter"], "views/parts/widgets-filter-form.html", "views/parts/widgets-table.html")
	Render(w, template, nil)

}

func getTemplate(t *template.Template, filePaths ...string) *template.Template {
	templatePart, err := t.ParseFiles(filePaths...)
	if err != nil {
		slog.Debug("error parsing template file", err)
		panic(err)
	}

	return templatePart
}
