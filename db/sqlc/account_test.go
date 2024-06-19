package db

import (
	"context"
	"github.com/Rahugg/GolangWithCoolStuff/util"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func createRandomAccount(t *testing.T) Account {
	arg := CreateAccountParams{
		Owner:    util.RandomOwner(),
		Balance:  util.RandomMoney(),
		Currency: util.RandomCurrency(),
	}

	account, err := testQueries.CreateAccount(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, account)
	require.Equal(t, arg.Owner, account.Owner)
	require.Equal(t, arg.Balance, account.Balance)
	require.Equal(t, arg.Currency, account.Currency)

	require.NotZero(t, account.ID)
	require.NotZero(t, account.CreatedAt)

	return account
}

func TestCreateAccount(t *testing.T) {
	createRandomAccount(t)
}

func TestGetAccount(t *testing.T) {
	randomAccount := createRandomAccount(t)
	gotRandomAccount, err := testQueries.GetAccount(context.Background(), randomAccount.ID)
	require.NoError(t, err)
	require.NotEmpty(t, gotRandomAccount)

	require.Equal(t, randomAccount.ID, gotRandomAccount.ID)
	require.Equal(t, randomAccount.Owner, gotRandomAccount.Owner)
	require.Equal(t, randomAccount.Balance, gotRandomAccount.Balance)
	require.Equal(t, randomAccount.Currency, gotRandomAccount.Currency)
	require.WithinDuration(t, randomAccount.CreatedAt, gotRandomAccount.CreatedAt, time.Second)
}

func TestUpdateAccount(t *testing.T) {
	randomAccount := createRandomAccount(t)

	arg := UpdateAccountParams{
		ID:      randomAccount.ID,
		Balance: util.RandomMoney(),
	}

	updatedAccount, err := testQueries.UpdateAccount(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, updatedAccount)
	require.Equal(t, randomAccount.ID, updatedAccount.ID)
	require.Equal(t, randomAccount.Owner, updatedAccount.Owner)
	require.Equal(t, arg.Balance, updatedAccount.Balance)
	require.Equal(t, randomAccount.Currency, updatedAccount.Currency)
	require.WithinDuration(t, randomAccount.CreatedAt, updatedAccount.CreatedAt, time.Second)
}

func TestDeleteAccount(t *testing.T) {
	randomAccount := createRandomAccount(t)
	err := testQueries.DeleteAccount(context.Background(), randomAccount.ID)
	require.NoError(t, err)

	deletedAccount, err := testQueries.GetAccount(context.Background(), randomAccount.ID)
	require.Error(t, err)
	require.Empty(t, deletedAccount)
}

func TestListAccounts(t *testing.T) {
	for i := 0; i < 10; i++ {
		createRandomAccount(t)
	}

	arg := ListAccountsParams{
		Limit:  5,
		Offset: 5,
	}

	accounts, err := testQueries.ListAccounts(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, accounts, 5)

	for _, account := range accounts {
		require.NotEmpty(t, account)
	}

}
