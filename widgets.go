package main

import (
	"fmt"
	"log"
	"net/http"
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

		w.Header().Set("HX-Location", fmt.Sprintf("/widget/update/%d", widget.ID))

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
