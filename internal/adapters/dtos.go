package adapters

import (
	"fmt"
	"time"

	"github.com/google/uuid"
)

type User struct {
	Id uuid.UUID `gorm:"type:uuid;primaryKey"`

	CreatedAt time.Time `gorm:"column:created_at"`

	FirstName string `gorm:"column:first_name"`
	LastName  string `gorm:"column:last_name"`
	Email     string `gorm:"column:email"`

	Roles []*Role `gorm:"many2many:users_roles;"`
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

	Users []*User `gorm:"many2many:users_roles;"`
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