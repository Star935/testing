package models

// Estructura que representa a una materia
type Subject struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}