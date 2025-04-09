package database

import (
	"github.com/ladecadence/PhotonAPI/pkg/models"
	"gorm.io/gorm"
)

type Database interface {
	Open(string) (*gorm.DB, error)
	Init() error
	UpsertWall(models.Wall) error
	DeleteWall(models.Wall) error
	GetWalls() ([]models.Wall, error)
	GetWall(string) (models.Wall, error)
}
