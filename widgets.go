package main

import (
	"fmt"
	"log"
	"net/http"
)

func HandleWidget(w http.ResponseWriter, r *http.Request) {

	var urlID string
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

		widget.Create()

	case "PUT":
		fmt.Println("method is put")

		err = widget.ValidateAndSet(r.Form)
		if err != nil {
			log.Fatal(err)
		}

		widget.Update()

	case "DELETE":
		fmt.Println("method is delete")
		widget.Delete()

	default:
		fmt.Printf("method is: %v \n", method)
		log.Fatal("incorrect method")
	}

}
