package models

// MissingModel is a struct that represents MissingModel
type MissingModel struct {
	ID           uint   `json:"id"`
	Name         string `json:"name"`
	Bibliography string `json:"bibliography"`
}

// TableName returns the name of the table where MissingModels are stored
// in the database
func (b *MissingModel) TableName() string { return "missing_models" }
