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
	err := s.db.AutoMigrate(&models.User{})
	if err != nil {
		return err
	}
	err = s.db.AutoMigrate(&models.Problem{})
	if err != nil {
		return err
	}
	return s.db.AutoMigrate(&models.Wall{})
}

func (s *SQLite) UpsertUser(u models.User) error {
	result := s.db.Clauses(clause.OnConflict{
		UpdateAll: true,
	}).Create(&u)
	return result.Error
}

func (s *SQLite) DeleteUser(models.User) error {
	// TODO
	return nil
}

func (s *SQLite) GetUsers() ([]models.User, error) {
	var users []models.User
	result := s.db.Find(&users)
	return users, result.Error
}

func (s *SQLite) GetUser(name string) (models.User, error) {
	var user models.User
	result := s.db.Where("name=?", name).First(&user)
	return user, result.Error
}

func (s *SQLite) UpsertWall(w models.Wall) error {
	result := s.db.Clauses(clause.OnConflict{
		UpdateAll: true,
	}).Create(&w)
	return result.Error
}

func (s *SQLite) DeleteWall(w models.Wall) error {
	// TODO
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

func (s *SQLite) UpsertProblem(p models.Problem) error {
	result := s.db.Clauses(clause.OnConflict{
		UpdateAll: true,
	}).Create(&p)
	return result.Error
}

func (s *SQLite) DeleteProblem(w models.Problem) error {
	// TODO
	return nil
}

func (s *SQLite) GetProblems() ([]models.Problem, error) {
	var problems []models.Problem
	result := s.db.Find(&problems)
	return problems, result.Error
}

func (s *SQLite) GetProblem(uid string) (models.Problem, error) {
	var problem models.Problem
	result := s.db.Where("uid=?", uid).First(&problem)
	return problem, result.Error
}

func (s *SQLite) GetWallProblems(wallid string) (models.Problem, error) {
	var problem models.Problem
	result := s.db.Where("wallid=?", wallid).First(&problem)
	return problem, result.Error
}
