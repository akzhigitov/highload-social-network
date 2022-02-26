package models

type Friend struct {
	Model
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}
