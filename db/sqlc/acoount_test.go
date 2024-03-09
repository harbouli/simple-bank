package simplebank

import (
	"context"
	"testing"
	"time"

	"github.com/harbouli/simplebank/utils"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/stretchr/testify/require"
)

func timestamptzToTime(tstz pgtype.Timestamptz) time.Time {
	// Assuming tstz has a Time field of type time.Time
	return tstz.Time
}

func CreateRandomAccount(t *testing.T) Account {
	arg := CreateAccountParams{
		Owner:    utils.RandomOwnerName(),
		Balance:  utils.RandomMoney(),
		Currency: utils.RandomCurrency(),
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
	CreateRandomAccount(t)

}

func TestGetAccount(t *testing.T) {
	account := CreateRandomAccount(t)
	account2, err := testQueries.GetAccount(context.Background(), account.ID)

	require.NoError(t, err)
	require.NotEmpty(t, account2)

	require.Equal(t, account2.Owner, account.Owner)
	require.Equal(t, account2.ID, account.ID)
	require.Equal(t, account2.Balance, account.Balance)
	require.Equal(t, account2.Currency, account.Currency)

	accountCreatedAtTime := timestamptzToTime(account.CreatedAt)
	account2CreatedAtTime := timestamptzToTime(account2.CreatedAt)

	require.WithinDuration(t, accountCreatedAtTime, account2CreatedAtTime, time.Second)

	require.NotZero(t, account2.ID)
	require.NotZero(t, account2.CreatedAt)

}
func TestUpdateAccount(t *testing.T) {
	newAccount := CreateRandomAccount(t)

	arg := UpdateAccountParams{
		ID:      newAccount.ID,
		Balance: 90,
	}
	account, err := testQueries.UpdateAccount(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, account)

	require.Equal(t, arg.ID, account.ID)
	require.Equal(t, arg.Balance, account.Balance)
	require.Equal(t, newAccount.Owner, account.Owner)
	require.Equal(t, newAccount.Currency, account.Currency)

	accountCreatedAtTime := timestamptzToTime(account.CreatedAt)
	account2CreatedAtTime := timestamptzToTime(newAccount.CreatedAt)

	require.WithinDuration(t, accountCreatedAtTime, account2CreatedAtTime, time.Second)

}

func TestDeleteAccount(t *testing.T) {
	newAccount := CreateRandomAccount(t)

	err := testQueries.DeleteAccount(context.Background(), newAccount.ID)
	require.NoError(t, err)

	account2, err := testQueries.GetAccount(context.Background(), newAccount.ID)

	require.Error(t, err)
	// require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, account2)
}

func TestListAccounts(t *testing.T) {
	for i := 0; i < 10; i++ {
		CreateRandomAccount(t)
	}

	arg := ListAccountsParams{
		Limit:  5,
		Offset: 5,
	}
	accounts, err := testQueries.ListAccounts(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, accounts)
	require.Len(t, accounts, 5)

	for _, account := range accounts {
		require.NotEmpty(t, account)

	}

}
