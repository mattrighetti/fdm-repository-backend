package models

type FloodModel struct {
	ID           uint   `json:"id"`
	Name         string `json:"name"`
	Acronym      string `json:"acronym"`
	Version      string `json:"version"`
	Soa          string `json:"soa"`
	Cod          string `json:"cod"`
	Floodtypei   string `json:"floodtypei"`
	Floodtypeii  string `json:"floodtypeii"`
	Modeltypei   string `json:"modeltypei"`
	Modeltypeii  string `json:"modeltypeii"`
	Modeltypeiii string `json:"modeltypeiii"`
	Eis          string `json:"eis"`
	Description  string `json:"description"`
}

type FloodModelNoDescription struct {
	ID           uint   `json:"id"`
	Name         string `json:"name"`
	Acronym      string `json:"acronym"`
	Version      string `json:"version"`
	Soa          string `json:"soa"`
	Cod          string `json:"cod"`
	Floodtypei   string `json:"floodtypei"`
	Floodtypeii  string `json:"floodtypeii"`
	Modeltypei   string `json:"modeltypei"`
	Modeltypeii  string `json:"modeltypeii"`
	Modeltypeiii string `json:"modeltypeiii"`
	Eis          string `json:"eis"`
}

func (b *FloodModel) TableName() string {
	return "models"
}

func (b *FloodModelNoDescription) TableName() string {
	return "models"
}
