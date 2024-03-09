package simplebank

import (
	"context"
	"testing"
	"time"

	"github.com/harbouli/simplebank/utils"
	"github.com/stretchr/testify/require"
)

func CreateRandomTransfer(t *testing.T) Transfer {
	account1 := CreateRandomAccount(t)
	account2 := CreateRandomAccount(t)

	arg := CreateTransferParams{
		FromAccountID: account1.ID,
		ToAccountID:   account2.ID,
		Amount:        utils.RandomMoney(),
	}

	transfer, err := testQueries.CreateTransfer(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, transfer)

	require.Equal(t, arg.FromAccountID, transfer.FromAccountID)
	require.Equal(t, arg.ToAccountID, transfer.ToAccountID)
	require.Equal(t, arg.Amount, transfer.Amount)

	require.NotZero(t, transfer.ID)
	require.NotZero(t, transfer.CreatedAt)

	return transfer
}

func TestCreateTransfer(t *testing.T) {
	CreateRandomTransfer(t)
}

func TestGetTransfer(t *testing.T) {
	newTransfer := CreateRandomTransfer(t)

	transfer, err := testQueries.GetTransfer(context.Background(), newTransfer.ID)

	require.NoError(t, err)
	require.NotEmpty(t, transfer)

	require.Equal(t, transfer.Amount, newTransfer.Amount)
	require.Equal(t, transfer.FromAccountID, newTransfer.FromAccountID)
	require.Equal(t, transfer.ToAccountID, newTransfer.ToAccountID)

	entryCreatedAtTime := timestamptzToTime(transfer.CreatedAt)
	entry2CreatedAtTime := timestamptzToTime(transfer.CreatedAt)

	require.WithinDuration(t, entryCreatedAtTime, entry2CreatedAtTime, time.Second)

	require.NotZero(t, transfer.ID)
	require.NotZero(t, transfer.CreatedAt)

}

func TestUpdateTransfer(t *testing.T) {
	newTransfer := CreateRandomTransfer(t)

	arg := UpdateTransferParams{ID: newTransfer.ID, Amount: 90}

	transfer, err := testQueries.UpdateTransfer(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, transfer)

	require.Equal(t, arg.Amount, transfer.Amount)
	require.Equal(t, transfer.FromAccountID, newTransfer.FromAccountID)
	require.Equal(t, transfer.ToAccountID, newTransfer.ToAccountID)

	entryCreatedAtTime := timestamptzToTime(transfer.CreatedAt)
	entry2CreatedAtTime := timestamptzToTime(transfer.CreatedAt)

	require.WithinDuration(t, entryCreatedAtTime, entry2CreatedAtTime, time.Second)

	require.NotZero(t, transfer.ID)
	require.NotZero(t, transfer.CreatedAt)

}

func TestDeleteTransfer(t *testing.T) {
	newTransfer := CreateRandomTransfer(t)

	err := testQueries.DeleteEntries(context.Background(), newTransfer.ID)

	require.NoError(t, err)
	entry, err := testQueries.GetEntry(context.Background(), newTransfer.ID)

	require.Error(t, err)
	require.Empty(t, entry)

}

func TestListTransfers(t *testing.T) {

	for i := 0; i < 10; i++ {
		CreateRandomTransfer(t)
	}
	arg := ListTransfersParams{
		Limit:  5,
		Offset: 5,
	}
	entries, err := testQueries.ListTransfers(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, entries)
	require.Len(t, entries, 5)
}
