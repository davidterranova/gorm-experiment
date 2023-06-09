package adapters

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/google/uuid"
)

type User struct {
	Id uuid.UUID `gorm:"type:uuid;primaryKey"`

	CreatedAt time.Time `gorm:"column:created_at"`

	FirstName string `gorm:"column:first_name"`
	LastName  string `gorm:"column:last_name"`

	Roles      []*Role     `gorm:"many2many:users_roles;"`
	UsersRoles []*UserRole `gorm:"foreignKey:UserId;references:Id"`

	Emails []*Email `gorm:"foreignKey:UserId;references:Id"`
}

func (u User) String() string {
	// return fmt.Sprintf(
	// 	"User{Id: %s, CreatedAt: %s, FirstName: %s, LastName: %s, Email: %s, Roles: %s, UsersRoles: %s}",
	// 	u.Id, u.CreatedAt, u.FirstName, u.LastName, u.Email, u.Roles, u.UsersRoles,
	// )

	data, _ := json.MarshalIndent(u, "", "  ")
	return string(data)
}

func (User) TableName() string {
	return "users"
}

type Email struct {
	Id     uuid.UUID `gorm:"type:uuid;primaryKey"`
	UserId uuid.UUID `gorm:"type:uuid"`

	CreatedAt time.Time `gorm:"column:created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at"`

	Email     string `gorm:"column:email"`
	Verified  bool   `gorm:"column:verified"`
	Principal bool   `gorm:"column:principal"`
}

func (Email) TableName() string {
	return "emails"
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

func (r Role) String() string {
	return fmt.Sprintf(
		"Role{Name: %s}",
		r.Name,
	)
}

type UserRole struct {
	UserId uuid.UUID `gorm:"type:uuid;primaryKey"`
	RoleId uuid.UUID `gorm:"type:uuid;primaryKey"`

	User *User `gorm:"foreignKey:UserId;references:Id"`
	Role *Role `gorm:"foreignKey:RoleId;references:Id"`

	Notes     string    `gorm:"column:notes"`
	CreatedAt time.Time `gorm:"column:created_at"`
}

func (UserRole) TableName() string {
	return "users_roles"
}

func (ur UserRole) String() string {
	return fmt.Sprintf(
		"UserRole{UserId: %s, RoleId: %s, CreatedAt: %s, Notes: %s}",
		ur.UserId, ur.RoleId, ur.CreatedAt, ur.Notes,
	)
}
