// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type Medicine struct {
	ID     string `json:"id"`
	Name   string `json:"name"`
	Active string `json:"active"`
}

type NewMedicine struct {
	Name   string `json:"name"`
	Active string `json:"active"`
}
