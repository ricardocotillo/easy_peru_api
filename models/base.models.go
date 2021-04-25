package models

import "gorm.io/gorm"

type Departament struct {
	gorm.Model
	Name string `json:"name"`
}

type EconomicActivity struct {
	gorm.Model
	Name string `json:"name"`
}

type PBI struct {
	gorm.Model
	Year               string `json:"year"`
	DepartamentID      int
	EconomicActivityID int
	Value              int    `json:"value"`
	Structure          string `json:"structure"`
	ValueType          string `json:"valueType"`
}
