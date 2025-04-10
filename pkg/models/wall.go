package models

import (
	"gorm.io/gorm"
)

func WallFields() []string {
	return []string{"uid", "name", "description", "adjustable", "deg_min", "deg_max", "image", "img_w", "img_h", "holds"}
}

type Wall struct {
	gorm.Model
	ID          uint    `gorm:"primaryKey"`
	Uid         string  `json:"uid" gorm:"unique"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Adjustable  bool    `json:"adjustable"`
	DegMin      float64 `json:"deg_min"`
	DegMax      float64 `json:"deg_max"`
	Image       []byte  `json:"image"`
	ImgW        int     `json:"img_w"`
	ImgH        int     `json:"img_h"`
	Holds       string  `json:"holds"`
}
