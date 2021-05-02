package models

import "gorm.io/gorm"

type Department struct {
	gorm.Model
	Name     string    `json:"name"`
	Products []Product `json:"products"`
}

type EconomicActivity struct {
	gorm.Model `json:"-"`
	Name       string    `json:"name"`
	Products   []Product `json:"products"`
}

type Product struct {
	gorm.Model
	Year               string
	DepartmentID       uint
	Department         Department
	EconomicActivityID uint
	EconomicActivity   EconomicActivity
	Value              float64
	Structure          string
	ValueType          string
}
