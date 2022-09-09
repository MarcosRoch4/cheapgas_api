package models

type Category struct {
	Id   uint   `json:"id"`
	Name string `json:"name"`
}

var Categories []Category