package models

type Person struct {
	ID   int    `gorm:"primary_key" form:"id" json:"id"`
	Name string `form:"name" json:"name" binding:"required"`
}
