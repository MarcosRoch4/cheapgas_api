package models

type Fuel struct {
	Id   uint   `json:"id"`
	Name string `json:"name"`
}

var Fuels []Fuel
