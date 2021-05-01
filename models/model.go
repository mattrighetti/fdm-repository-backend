package models

import (
	_ "github.com/go-sql-driver/mysql"
)

type FloodModel struct {
	ID           uint   `json:"id"`
	Name         string `json:"name"`
	Acronym      string `json:"acronym"`
	Version      string `json:"version"`
	Cod          string `json:"cod"`
	Soa          string `json:"soa"`
	Floodtypei   string `json:"floodtypei"`
	Floodtypeii  string `json:"floodtypeii"`
	Modeltypei   string `json:"modeltypei"`
	Modeltypeii  string `json:"modeltypeii"`
	Modeltypeiii string `json:"modeltypeiii"`
	Eis          string `json:"eis"`
	Description  string `json:"description,omitempty"`
}

func (b *FloodModel) TableName() string {
	return "models"
}