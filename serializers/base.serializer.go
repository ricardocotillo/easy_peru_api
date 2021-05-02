package serializers

import "cotillo.tech/inei_api/models"

type ProductAggregateSerializer struct {
	ID    uint    `json:"id"`
	Year  string  `json:"year"`
	Value float64 `json:"value"`
}

type ProductBySerializer struct {
	ID                 uint                    `json:"id"`
	Year               string                  `json:"year"`
	Value              float64                 `json:"value"`
	DepartmentID       uint                    `json:"-"`
	Department         models.Department       `json:"department"`
	EconomicActivityID uint                    `json:"-"`
	EconomicActivity   models.EconomicActivity `json:"economicActivity"`
}

type DepartmentSerializer struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

type EconomicActivitySerializer struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}
