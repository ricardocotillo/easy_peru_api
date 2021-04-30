package models

import "gorm.io/gorm"

type Department struct {
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
	Year               string           `json:"year"`
	DepartmentID       uint             `json:"-"`
	Department         Department       `json:"department"`
	EconomicActivityID uint             `json:"-"`
	EconomicActivity   EconomicActivity `json:"economicActivity"`
	Value              float64          `json:"value"`
	Structure          string           `json:"structure"`
	ValueType          string           `json:"valueType"`
}
