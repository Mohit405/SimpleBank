package db

import (
	"context"
	"testing"
	"time"

	"github.com/mohit405/simplebank/util"
	"github.com/stretchr/testify/require"
)

func createRandomEntery(t *testing.T, account Account) Entery {
	args := CreateEnteryParams{
		AccountID: account.ID,
		Amount:    util.RandomMoney(),
	}

	entery, err := testQueries.CreateEntery(context.Background(), args)

	require.NoError(t, err)
	require.NotEmpty(t, entery)

	require.Equal(t, args.AccountID, entery.AccountID)
	require.Equal(t, args.Amount, entery.Amount)

	require.NotZero(t, entery.ID)
	require.NotZero(t, entery.CreatedAt)

	return entery
}

func TestCreateEntery(t *testing.T) {
	account := createRandomAccount(t)
	createRandomEntery(t, account)
}

func TestGetEntery(t *testing.T) {
	account := createRandomAccount(t)
	entery1 := createRandomEntery(t, account)
	entery2, err := testQueries.GetEntry(context.Background(), entery1.ID)

	require.NoError(t, err)
	require.NotEmpty(t, entery2)

	require.Equal(t, entery1.AccountID, entery2.AccountID)
	require.Equal(t, entery1.ID, entery2.ID)
	require.Equal(t, entery1.Amount, entery2.Amount)
	require.WithinDuration(t, entery1.CreatedAt, entery2.CreatedAt, time.Second)
}
