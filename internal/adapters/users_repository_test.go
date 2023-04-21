package adapters_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/davidterranova/gorm-experiment/internal/adapters"
	"github.com/davidterranova/gorm-experiment/pkg/pg"
	"github.com/google/uuid"
	"github.com/kelseyhightower/envconfig"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"gorm.io/gorm"
)

func TestFindById(t *testing.T) {
	repo := adapters.NewUserRepository(dbConn(t))
	user, err := repo.FindById(uuid.MustParse("00000000-0000-0000-0000-000000000001"))
	assert.NoError(t, err)
	assert.NotEmpty(t, user)
	assert.NotEmpty(t, user.Roles)
	assert.NotNil(t, user.UsersRoles)
	fmt.Println(user)
}

func TestFindAll(t *testing.T) {
	repo := adapters.NewUserRepository(dbConn(t))
	users, err := repo.FindAll()
	assert.NoError(t, err)
	assert.NotEmpty(t, users)

	for _, user := range users {
		assert.NotEmpty(t, user.Roles)
		assert.NotNil(t, user.UsersRoles)
		fmt.Println(user)
	}
}

func TestCreate(t *testing.T) {
	usersRepo := adapters.NewUserRepository(dbConn(t))
	rolesRepo := adapters.NewRoleRepository(dbConn(t))

	t.Run("without custom association data", func(t *testing.T) {
		userRole, err := rolesRepo.FindByName("user")
		assert.NoError(t, err)

		jdoe := &adapters.User{
			Id:        uuid.New(),
			CreatedAt: time.Now().UTC().Round(time.Second),
			FirstName: "John",
			LastName:  "Doe",
			Email:     "jdoe@test.local",
			Roles:     []*adapters.Role{userRole},
		}

		_, err = usersRepo.Create(jdoe)
		assert.NoError(t, err)

		loadedJdoe, err := usersRepo.FindById(jdoe.Id)
		assert.NoError(t, err)
		assert.Equal(t, jdoe.Id, loadedJdoe.Id)

		err = usersRepo.Delete(jdoe)
		assert.NoError(t, err)
	})

	t.Run("with custom association data set manually", func(t *testing.T) {
		userRole, err := rolesRepo.FindByName("user")
		assert.NoError(t, err)

		jdoe := &adapters.User{
			Id:         uuid.New(),
			CreatedAt:  time.Now().UTC().Round(time.Second),
			FirstName:  "John",
			LastName:   "Doe",
			Email:      "jdoe@test.local",
			UsersRoles: []*adapters.UserRole{},
		}

		notes := "default role"
		jdoe.UsersRoles = append(jdoe.UsersRoles, &adapters.UserRole{
			UserId: jdoe.Id,
			RoleId: userRole.Id,
			Notes:  notes,
		})

		_, err = usersRepo.Create(jdoe)
		assert.NoError(t, err)

		loadedJdoe, err := usersRepo.FindById(jdoe.Id)
		assert.NoError(t, err)
		assert.Equal(t, jdoe.Id, loadedJdoe.Id)
		assert.Equal(t, "default role", loadedJdoe.UsersRoles[0].Notes)

		err = usersRepo.Delete(jdoe)
		assert.NoError(t, err)
	})

	t.Run("with custom association data set automatically", func(t *testing.T) {
		userRole, err := rolesRepo.FindByName("user")
		assert.NoError(t, err)

		jdoe := &adapters.User{
			Id:         uuid.New(),
			CreatedAt:  time.Now().UTC().Round(time.Second),
			FirstName:  "John",
			LastName:   "Doe",
			Email:      "jdoe@test.local",
			UsersRoles: []*adapters.UserRole{},
		}

		notes := "default role"
		jdoe.UsersRoles = append(jdoe.UsersRoles, &adapters.UserRole{
			User:  jdoe,
			Role:  userRole,
			Notes: notes,
		})

		_, err = usersRepo.Create(jdoe)
		assert.NoError(t, err)

		loadedJdoe, err := usersRepo.FindById(jdoe.Id)
		assert.NoError(t, err)
		assert.Equal(t, jdoe.Id, loadedJdoe.Id)
		assert.Equal(t, "default role", loadedJdoe.UsersRoles[0].Notes)

		err = usersRepo.Delete(jdoe)
		assert.NoError(t, err)
	})
}

func dbConn(t *testing.T) *gorm.DB {
	t.Helper()

	cfg := pg.DBConfig{}
	err := envconfig.Process("GORM_EXPERIMENT_DB", &cfg)
	require.NoError(t, err)

	db, err := pg.Open(cfg)
	require.NoError(t, err)

	return db
}
