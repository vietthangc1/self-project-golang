package mysql

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/thangpham4/self-project/entities"
	"github.com/thangpham4/self-project/pkg/randomx"
)

func TestCreateUser(t *testing.T) {
	createRandomUser(t)
}

func TestGetUserByID(t *testing.T) {
	newUser := createRandomUser(t)
	r := require.New(t)

	user, err := userAdmin.Get(context.Background(), newUser.ID)
	r.NoError(err)
	r.NotEmpty(user)

	r.Equal(newUser.Name, user.Name)
	r.Equal(newUser.Role, user.Role)
	r.Equal(newUser.Email, user.Email)

	r.NotEqual(newUser.Password, user.Password)
}

func TestGetUserByEmail(t *testing.T) {
	newUser := createRandomUser(t)
	r := require.New(t)

	user, err := userAdmin.GetByEmail(context.Background(), newUser.Email)
	r.NoError(err)
	r.NotEmpty(user)

	r.Equal(newUser.Name, user.Name)
	r.Equal(newUser.Role, user.Role)
	r.Equal(newUser.ID, user.ID)

	r.NotEqual(newUser.Password, user.Password)
}

func createRandomUser(t *testing.T) entities.UserAdmin {
	r := require.New(t)

	name := randomx.RandString(10)
	role := 2
	email := randomx.RandEmail()
	password := randomx.RandString(8)

	inputUser := entities.UserAdmin{
		Name:     name,
		Role:     int32(role),
		Email:    email,
		Password: password,
	}

	newUser, err := userAdmin.Create(context.Background(), inputUser)
	r.NoError(err)
	r.NotEmpty(newUser)

	r.Equal(name, newUser.Name)
	r.Equal(role, newUser.Role)
	r.Equal(email, newUser.Email)
	r.NotEmpty(newUser.ID)
	inputUser.ID = newUser.ID

	return inputUser
}
