package database

import (
	"fmt"

	"github.com/ladecadence/PhotonAPI/pkg/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/logger"
)

type SQLite struct {
	db *gorm.DB
}

func (s *SQLite) Open(fileName string) (*gorm.DB, error) {
	database, err := gorm.Open(sqlite.Open(fileName), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
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

func (s *SQLite) GetWalls(fields []string) ([]models.Wall, error) {
	var walls []models.Wall
	var result *gorm.DB
	if fields != nil && len(fields) > 0 {
		result = s.db.Select(fields).Find(&walls)
	} else {
		result = s.db.Find(&walls)
	}
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

func (s *SQLite) GetProblems(page int, page_size int, filter models.ProblemFilter) ([]models.Problem, error) {
	var result *gorm.DB

	var problems []models.Problem
	var tx *gorm.DB

	if filter.Active {
		if filter.OrderBy != models.FilterOrderByNothing {
			order := ""
			switch filter.OrderBy {
			case models.FilterOrderByGrade:
				order += "grade"
			case models.FilterOrderByName:
				order += "name"
			case models.FilterOrderBySends:
				order += "sends"
			}

			switch filter.OrderDir {
			case models.FilterOrderAsc:
				order += " asc"
			case models.FilterOrderDesc:
				order += " desc"
			default:
				order += " asc"
			}
			tx = s.db.Order(order)
		}
		if filter.WallID != "" {
			if tx == nil {
				tx = s.db.Where("wall_id", filter.WallID)
			} else {
				tx = tx.Where("wall_id", filter.WallID)
			}
		}

		if filter.GradeRange != nil {
			if tx == nil {
				tx = s.db.Where("grade >= " + fmt.Sprintf("%v", filter.GradeRange[0]) + " AND grade <= " + fmt.Sprintf("%v", filter.GradeRange[1]))
			} else {
				tx = tx.Where("grade >= " + fmt.Sprintf("%v", filter.GradeRange[0]) + " AND grade <= " + fmt.Sprintf("%v", filter.GradeRange[1]))
			}
		}

	} else {
		tx = s.db.Order("id asc")
	}

	if page >= 0 && page_size > 0 {
		result = tx.Limit(page_size).Offset(page * page_size).Find(&problems)
	} else {
		result = tx.Find(&problems)
	}
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
