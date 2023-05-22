package db

import (
	"context"
	"testing"

	"github.com/ngtrdai197/cobra-cmd/util"
	"github.com/stretchr/testify/require"
)

func createRandomUser(t *testing.T) User {
	args := CreateUserParams{
		Username:    util.RandomUser(6),
		FullName:    util.RandomUser(10),
		PhoneNumber: util.RandomPhoneNumber(),
	}
	user, err := testQueries.CreateUser(context.Background(), args)

	require.NoError(t, err)
	require.NotEmpty(t, user)

	require.Equal(t, args.Username, user.Username)
	require.Equal(t, args.FullName, user.FullName)
	require.Equal(t, args.PhoneNumber, user.PhoneNumber)

	require.NotZero(t, user.Username)
	require.NotZero(t, user.FullName)
	require.NotZero(t, user.PhoneNumber)
	require.NotZero(t, user.CreatedAt)

	return user
}

func TestCreateUser(t *testing.T) {
	createRandomUser(t)
}
