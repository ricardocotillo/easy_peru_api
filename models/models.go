package models

type Cuadro struct {
	Titulo        string         `json:"titulo"`
	TipoDeValores string         `json:"tipo_de_valores"`
	Estructura    string         `json:"estructura"`
	Departamentos []Departamento `json:"departamentos"`
}

type Departamento struct {
	Nombre string `json:"nombre"`
	Años   []Año  `json:"años"`
}
type Año struct {
	Año                string      `json:"año"`
	VapasAgregadoBruto float32     `json:"vapas_agregado_bruto"`
	Actividades        []Actividad `json:"actividades"`
}

type Actividad struct {
	Nombre string  `json:"nombre"`
	Valor  float64 `json:"valor"`
}
