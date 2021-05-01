package models

import "gorm.io/gorm"

type Department struct {
	gorm.Model `json:"-"`
	Name       string    `json:"name"`
	Products   []Product `json:"products"`
}

type EconomicActivity struct {
	gorm.Model `json:"-"`
	Name       string    `json:"name"`
	Products   []Product `json:"products"`
}

type Product struct {
	gorm.Model         `json:"-"`
	Year               string           `json:"year"`
	DepartmentID       uint             `json:"-"`
	Department         Department       `json:"-"`
	EconomicActivityID uint             `json:"-"`
	EconomicActivity   EconomicActivity `json:"-"`
	Value              float64          `json:"value"`
	Structure          string           `json:"-"`
	ValueType          string           `json:"-"`
}
