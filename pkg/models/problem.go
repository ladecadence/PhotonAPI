package models

import (
	"gorm.io/gorm"
)

const (
	FilterOrderByNothing = iota
	FilterOrderByName
	FilterOrderByGrade
	FilterOrderBySends

	FilterOrderAsc  = 0
	FilterOrderDesc = 1
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

type ProblemFilter struct {
	Active     bool
	WallID     string
	GradeRange []int
	OrderDir   int
	OrderBy    int
}

func (f *ProblemFilter) Clear() {
	f.Active = false
	f.WallID = ""
	f.GradeRange = nil
	f.OrderBy = FilterOrderByNothing
	f.OrderDir = FilterOrderAsc
}

func (f *ProblemFilter) SetWallID(wallid string) {
	f.Active = true
	f.WallID = wallid
}

func (f *ProblemFilter) SetGradeRange(gmin, gmax int) {
	f.Active = true
	f.GradeRange = nil
	f.GradeRange = append(f.GradeRange, gmin)
	f.GradeRange = append(f.GradeRange, gmax)
}

func (f *ProblemFilter) SetOrderDir(dir int) {
	f.Active = true
	if dir >= FilterOrderAsc && dir <= FilterOrderDesc {
		f.OrderDir = dir
	}
}

func (f *ProblemFilter) SetOrderBy(order int) {
	f.Active = true
	if order >= FilterOrderByNothing && order <= FilterOrderBySends {
		f.OrderBy = order
	}
}
