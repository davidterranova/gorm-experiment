package adapters

import "gorm.io/gorm"

type RoleRepository struct {
	db *gorm.DB
}

func NewRoleRepository(db *gorm.DB) *RoleRepository {
	return &RoleRepository{db}
}

func (r *RoleRepository) FindByName(name string) (*Role, error) {
	var role Role

	err := r.db.Where("name = ?", name).First(&role).Error

	return &role, err
}

func (r *RoleRepository) FindByNameWithUsers(name string) (*Role, error) {
	var role Role

	err := r.db.Preload("Users").Where("name = ?", name).First(&role).Error

	return &role, err
}
