package serializers

type ProductAggregateSerializer struct {
	ID    uint    `json:"id"`
	Year  string  `json:"year"`
	Value float64 `json:"value"`
}

type ProductBySerializer struct {
	ID                 uint                       `json:"id"`
	Year               string                     `json:"year"`
	Value              float64                    `json:"value"`
	DepartmentID       uint                       `json:"-"`
	Department         DepartmentSerializer       `json:"department"`
	EconomicActivityID uint                       `json:"-"`
	EconomicActivity   EconomicActivitySerializer `json:"economicActivity"`
}

type DepartmentSerializer struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

type EconomicActivitySerializer struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}
