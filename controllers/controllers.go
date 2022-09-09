package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/MarcosRoch4/api-rest-go/database"
	"github.com/MarcosRoch4/api-rest-go/models"
	"github.com/gorilla/mux"
)

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Home Page")
}

// Postos de Combutíveis Disponíveis

func TodosPosto(w http.ResponseWriter, r *http.Request) {
	var p []models.Gas_Station
	database.DB.Order("id").Table("gas_station").Find(&p)
	json.NewEncoder(w).Encode(p)
}

func RetornaUmPosto(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var posto models.Gas_Station
	database.DB.Table("gas_station").First(&posto, id)
	json.NewEncoder(w).Encode(posto)

}

func NovoPosto(w http.ResponseWriter, r *http.Request) {
	var posto models.Gas_Station
	json.NewDecoder(r.Body).Decode(&posto)
	database.DB.Table("gas_station").Create(&posto)
	json.NewEncoder(w).Encode(&posto)
}

func DeletaPosto(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var posto models.Gas_Station
	database.DB.Table("gas_station").Delete(&posto, id)
	json.NewEncoder(w).Encode(posto)
}

func EditaPosto(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var posto models.Gas_Station
	database.DB.Table("gas_station").First(&posto, id)
	json.NewDecoder(r.Body).Decode(&posto)
	database.DB.Table("gas_station").Save(&posto)
	json.NewEncoder(w).Encode(posto)
}

// Combutíveis Disponíveis

func TodosCombustiveis(w http.ResponseWriter, r *http.Request) {
	var p []models.Fuel
	database.DB.Order("id").Table("fuel").Find(&p)
	json.NewEncoder(w).Encode(p)
}

func RetornaUmCombustivel(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var combustivel models.Fuel
	database.DB.Table("fuel").First(&combustivel, id)
	json.NewEncoder(w).Encode(combustivel)

}

func NovoCombustivel(w http.ResponseWriter, r *http.Request) {
	var combustivel models.Fuel
	json.NewDecoder(r.Body).Decode(&combustivel)
	database.DB.Table("fuel").Create(&combustivel)
	json.NewEncoder(w).Encode(&combustivel)
}

func DeletaCombustivel(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var combustivel models.Fuel
	database.DB.Table("combustiveis").Delete(&combustivel, id)
	json.NewEncoder(w).Encode(combustivel)
}

func EditaCombustivel(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var combustivel models.Fuel
	database.DB.Table("fuel").First(&combustivel, id)
	json.NewDecoder(r.Body).Decode(&combustivel)
	database.DB.Table("fuel").Save(&combustivel)
	json.NewEncoder(w).Encode(combustivel)
}

// Valores dos Combustíveis

func TodosValores(w http.ResponseWriter, r *http.Request) {
	var p []models.Fuel_Value
	database.DB.Order("id").Table("fuel_value").Find(&p)
	json.NewEncoder(w).Encode(p)
}

func RetornaUmValor(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var valor models.Fuel_Value
	database.DB.Table("fuel_value").First(&valor, id)
	json.NewEncoder(w).Encode(valor)

}

func NovoValor(w http.ResponseWriter, r *http.Request) {
	var valor models.Fuel_Value
	json.NewDecoder(r.Body).Decode(&valor)
	database.DB.Table("fuel_value").Create(&valor)
	json.NewEncoder(w).Encode(&valor)
}

func DeletaValor(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var valor models.Fuel_Value
	database.DB.Table("fuel_value").Delete(&valor, id)
	json.NewEncoder(w).Encode(valor)
}

func EditaValor(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var valor models.Fuel_Value
	database.DB.Table("fuel_value").First(&valor, id)
	json.NewDecoder(r.Body).Decode(&valor)
	database.DB.Table("fuel_value").Save(&valor)
	json.NewEncoder(w).Encode(valor)
}

// Categoria dos Postos de Combustíveis

func TodasCategorias(w http.ResponseWriter, r *http.Request) {
	var p []models.Category
	database.DB.Order("id").Table("category").Find(&p)
	json.NewEncoder(w).Encode(p)
}

func RetornaUmaCategoria(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var categoria models.Category
	database.DB.Table("category").First(&categoria, id)
	json.NewEncoder(w).Encode(categoria)

}

func NovaCategoria(w http.ResponseWriter, r *http.Request) {
	var categoria models.Category
	json.NewDecoder(r.Body).Decode(&categoria)
	database.DB.Table("category").Create(&categoria)
	json.NewEncoder(w).Encode(&categoria)
}

func DeletaCategoria(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var categoria models.Category
	database.DB.Table("category").Delete(&categoria, id)
	json.NewEncoder(w).Encode(categoria)
}

func EditaCategoria(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var categoria models.Category
	database.DB.Table("category").First(&categoria, id)
	json.NewDecoder(r.Body).Decode(&categoria)
	database.DB.Table("category").Save(&categoria)
	json.NewEncoder(w).Encode(categoria)
}
