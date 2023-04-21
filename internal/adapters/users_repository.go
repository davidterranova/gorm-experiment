package adapters

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db}
}

func (r *UserRepository) FindAll() ([]*User, error) {
	var users []*User

	err := r.db.Preload("Roles").Preload("UsersRoles").Find(&users).Error

	return users, err
}

func (r *UserRepository) FindById(id uuid.UUID) (*User, error) {
	var user User

	// err := r.db.Preload("Roles").Preload("UsersRoles").Where("id = ?", id).First(&user).Error
	err := r.db.Preload("UsersRoles.Role").Preload(clause.Associations).Where("id = ?", id).First(&user).Error

	return &user, err
}

func (r *UserRepository) Create(user *User) (*User, error) {
	err := r.db.Create(&user).Error

	return user, err
}

func (r *UserRepository) Delete(user *User) error {
	return r.db.Select("Roles").Delete(&user).Error
}
