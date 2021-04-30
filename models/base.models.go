package models

import "gorm.io/gorm"

type Departament struct {
	gorm.Model
	Name     string    `json:"name"`
	Products []Product `json:"products"`
}

type EconomicActivity struct {
	gorm.Model
	Name     string    `json:"name"`
	Products []Product `json:"products"`
}

type Product struct {
	gorm.Model
	Year               string `json:"year"`
	DepartamentID      uint
	Departament        Departament `json:"department"`
	EconomicActivityID uint
	EconomicActivity   EconomicActivity `json:"economicActivity"`
	Value              float64          `json:"value"`
	Structure          string           `json:"structure"`
	ValueType          string           `json:"valueType"`
}
