package adapters_test

import (
	"testing"

	"github.com/davidterranova/gorm-experiment/internal/adapters"
	"github.com/stretchr/testify/assert"
)

func TestFindByNameWithUsers(t *testing.T) {
	rolesRepo := adapters.NewRoleRepository(dbConn(t))

	role, err := rolesRepo.FindByNameWithUsers("user")
	assert.NoError(t, err)

	assert.NotEmpty(t, role.Users)
}
