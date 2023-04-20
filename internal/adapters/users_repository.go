package adapters

import (
	"fmt"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db}
}

func (r *UserRepository) FindAll() ([]User, error) {
	var users []User

	err := r.db.Preload("Roles").Find(&users).Error

	return users, err
}

func (r *UserRepository) FindById(id uuid.UUID) (*User, error) {
	var user User

	err := r.db.Preload("Roles").Where("id = ?", id).First(&user).Error

	return &user, err
}

func (r *UserRepository) Create(user *User) (*User, error) {
	err := r.db.Create(&user).Error

	return user, err
}

func (r *UserRepository) Delete(user *User) error {
	return r.db.Select("Roles").Delete(&user).Error
}

type User struct {
	Id uuid.UUID `gorm:"type:uuid;primaryKey"`

	CreatedAt time.Time `gorm:"column:created_at"`

	FirstName string `gorm:"column:first_name"`
	LastName  string `gorm:"column:last_name"`
	Email     string `gorm:"column:email"`

	Roles []Role `gorm:"many2many:users_roles;"`
}

func (u User) String() string {
	return fmt.Sprintf(
		"User{Id: %s, CreatedAt: %s, FirstName: %s, LastName: %s, Email: %s, Roles: %s}",
		u.Id, u.CreatedAt, u.FirstName, u.LastName, u.Email, u.Roles,
	)
}

func (User) TableName() string {
	return "users"
}

type Role struct {
	Id uuid.UUID `gorm:"type:uuid;primaryKey"`

	CreatedAt time.Time `gorm:"column:created_at"`

	Name string `gorm:"column:name"`
}

func (Role) TableName() string {
	return "roles"
}

type UserRole struct {
	UserId uuid.UUID `gorm:"type:uuid;primaryKey"`
	RoleId uuid.UUID `gorm:"type:uuid;primaryKey"`

	CreatedAt time.Time `gorm:"column:created_at"`
}

func (r Role) String() string {
	return fmt.Sprintf(
		"Role{Name: %s}",
		r.Name,
	)
}

func (UserRole) TableName() string {
	return "users_roles"
}
