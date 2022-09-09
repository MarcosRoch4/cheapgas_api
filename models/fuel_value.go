package models

type Fuel_Value struct {
	Id             uint    `json:"id"`
	Price          float64 `json:"price"`
	Fuel_Id        int     `json:"fuel_id"`
	Gas_Station_Id int     `json:"gas_station_id"`
}

var Fuel_Values []Fuel_Value
