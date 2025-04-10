package models

import (
	"gorm.io/gorm"
)

func ProblemFields() []string {
	return []string{"uid", "wallid", "name", "description", "grade", "rating", "sends", "holds"}
}

type Problem struct {
	gorm.Model
	ID          uint   `gorm:"primaryKey"`
	Uid         string `json:"uid" gorm:"unique"`
	WallID      string `json:"wallid"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Grade       int    `json:"grade"`
	Rating      int    `json:"rating"`
	Sends       int    `json:"sends"`
	Holds       string `json:"holds"`
}
