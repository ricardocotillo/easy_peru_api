package models

import "gorm.io/gorm"

type Departament struct {
	gorm.Model
	Name     string `json:"name"`
	Products []Product
}

type EconomicActivity struct {
	gorm.Model
	Name     string `json:"name"`
	Products []Product
}

type Product struct {
	gorm.Model
	Year               string `json:"year"`
	DepartamentID      uint
	EconomicActivityID uint
	Value              float64 `json:"value"`
	Structure          string  `json:"structure"`
	ValueType          string  `json:"valueType"`
}
