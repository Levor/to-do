package repositories

import (
	"github.com/jinzhu/gorm"
	"github.com/levor/to-do/internal/models"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) Create(m *models.User) error {
	if err := r.db.Create(m).Error; err != nil {
		return err
	}
	return nil
}

func (r *UserRepository) GetById(id int64) (models.User, error) {
	var user models.User
	if err := r.db.Where("id = ?", id).First(&user).Error; err != nil {
		return models.User{}, err
	}
	return user, nil
}
func (r *UserRepository) FindByLogin(login string) (models.User, error) {
	var user models.User
	if err := r.db.Where("login = ?", login).First(&user).Error; err != nil {
		return models.User{}, err
	}
	return user, nil
}
func (r *UserRepository) GetAll() ([]models.User, error) {
	var users []models.User
	if err := r.db.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}
