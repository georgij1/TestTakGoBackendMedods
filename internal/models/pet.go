package models

// Pet represents a pet in the system
type Pet struct {
	ID   int    `json:"id" example:"1"`
	Name string `json:"name" example:"Fluffy"`
	Type string `json:"type,omitempty" example:"cat"`
}