package models

import "time"

type Note struct {
	ID        uint   `json:"id" gorm:"primaryKey"`
	Title     string `json:"title"`
	Content   string `json:"content"`
	Archived  bool   `json:"archived"`
	Favorite  bool   `json:"favorite"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (Note) TableName() string {
	return "notes"
}
