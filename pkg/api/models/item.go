package models

// Item represents an item with a name and description
type Item struct {
	Name        string `json:"name" binding:"required"`
	Description string `json:"description" binding:"required"`
}
