package database

import (
	"github.com/ladecadence/PhotonAPI/pkg/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type SQLite struct {
	db *gorm.DB
}

func (s *SQLite) Open(fileName string) (*gorm.DB, error) {
	database, err := gorm.Open(sqlite.Open(fileName), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	s.db = database
	return s.db, nil
}

func (s *SQLite) Init() error {
	return s.db.AutoMigrate(&models.Wall{})
}

func (s *SQLite) UpsertWall(w models.Wall) error {
	result := s.db.Clauses(clause.OnConflict{
		UpdateAll: true,
	}).Create(&w)
	return result.Error
}

func (s *SQLite) DeleteWall(w models.Wall) error {
	return nil
}

func (s *SQLite) GetWalls() ([]models.Wall, error) {
	var walls []models.Wall
	result := s.db.Find(&walls)
	return walls, result.Error
}

func (s *SQLite) GetWall(uid string) (models.Wall, error) {
	var wall models.Wall
	result := s.db.Where("uid=?", uid).First(&wall)
	return wall, result.Error
}
