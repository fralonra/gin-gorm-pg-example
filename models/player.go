package models

import (
	// "github.com/jinzhu/gorm"
)

const (
	PG = iota
	SG
	SF
	PF
	C
)

type Player struct {
	// gorm.Model
	ID       int      `gorm:"primary_key" form:"id" json:"id"`
	Name     string   `form:"name" json:"name" binding:"required"`
	Height   int      `form:"height" json:"height"`
	Weight   int      `form:"weight" json:"weight"`
	Number   int      `form:"number" json:"number"`
	Position int      `form:"position" json:"position"`
	Rating   *ratings `form:"ratings" json:"ratings"`
}

type ratings struct {
	// gorm.Model
	Dribble int `form:"dribble" json:"dribble"`
	Pass    int `form:"pass" json:"pass"`
}