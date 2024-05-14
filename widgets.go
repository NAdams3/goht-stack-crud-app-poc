package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func HandleWidget(w http.ResponseWriter, r *http.Request) {

	urlID := r.PathValue("id")
	fmt.Printf("urlID: %v \n", urlID)
	id, err := ValidateID(urlID)
	if err != nil {
		log.Fatal(err)
	}

	err = r.ParseForm()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("form data: %v \n", r.Form)

	var widget Widget
	widget.ID = id

	switch method := r.Method; method {
	case "POST":
		fmt.Println("method is post")

		err = widget.ValidateAndSet(r.Form)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("widget: %v \n", widget)

		err = widget.Create()
		if err != nil {
			log.Fatal(err)
		}

		w.Header().Set("HX-Location", fmt.Sprintf("/widget/edit/%d", widget.ID))

	case "PUT":
		fmt.Println("method is put")

		err = widget.ValidateAndSet(r.Form)
		if err != nil {
			log.Fatal(err)
		}

		err = widget.Update()
		if err != nil {
			log.Fatal(err)
		}

	case "DELETE":
		fmt.Println("method is delete")
		err = widget.Delete()
		if err != nil {
			log.Fatal(err)
		}

	default:
		fmt.Printf("method is: %v \n", method)
		log.Fatal("incorrect method")
	}

}

func WidgetCreate(w http.ResponseWriter, r *http.Request) {
	Render(w, templates["widgets-form"], nil)
}

func WidgetUpdate(w http.ResponseWriter, r *http.Request) {
	urlID := r.PathValue("id")
	id, err := strconv.Atoi(urlID)
	if err != nil {
		log.Fatal(err)
	}
	var widget Widget
	err = widget.GetByID(id)
	if err != nil {
		log.Fatal(err)
	}

	Render(w, templates["widgets-form"], widget)
}

func WidgetFilter(w http.ResponseWriter, r *http.Request) {

	if templates["widget-filter"] == nil {
		templates["widget-filter"] = getTemplate(templates["filter-adapter"], "views/parts/widgets-filter-form.html", "views/parts/widgets-table.html")
	}

	var widget Widget

	items, err := widget.Filter()
	if err != nil {
		log.Fatal(err)
	}

	Render(w, templates["widget-filter"], items)

}

func FilterWidgets(w http.ResponseWriter, r *http.Request) {
	fmt.Println("in filter widgets")
	err := r.ParseForm()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("filter results: %v \n", r.Form)

	var widget Widget
	err = widget.ValidateAndSet(r.Form)
	if err != nil {
		log.Fatal(err)
	}

	items, err := widget.Filter()
	if err != nil {
		log.Fatal(err)
	}

	err = templates["widget-filter"].ExecuteTemplate(w, "filter-display", items)
	if err != nil {
		log.Fatal(err)
	}

}
