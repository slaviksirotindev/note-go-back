package domain

type Note struct {
	ID          int    `json:"id" gorm:"primaryKey"`
	Name        string `json:"name"`
	Description string `json:"description"`
}
