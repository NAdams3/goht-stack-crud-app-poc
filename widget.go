package main

import (
	"errors"
	"fmt"
	"net/url"
)

type Widget struct {
	ID   int
	Name string
	Type string
}

func (widget *Widget) ValidateAndSet(form url.Values) error {
	var err error

	if form["name"] == nil || len(form["name"]) == 0 || form["name"][0] == "" {
		err = errors.Join(err, errors.New("Name is invalid"))
	} else {
		widget.Name = form["name"][0]
	}

	if form["type"] == nil || len(form["type"]) == 0 || form["type"][0] == "" {
		err = errors.Join(err, errors.New("Type is invalid"))
	} else {
		widget.Type = form["type"][0]
	}

	fmt.Printf("widget in val and set: %v \n", widget)

	return err
}

func (widget *Widget) Create() error {
	query := `INSERT INTO widgets (name, type) VALUES (?, ?)`

	result, err := POOL.Exec(query, widget.Name, widget.Type)
	if err != nil {
		return err
	}

	lastID, err := result.LastInsertId()
	if err != nil {
		return err
	}

	widget.ID = int(lastID)

	return nil
}

func (widget Widget) Update() error {
	fmt.Printf("updating widget: %v \n", widget)
	return nil
}

func (widget Widget) Delete() error {
	fmt.Printf("deleting widget: %v \n", widget)
	return nil
}
