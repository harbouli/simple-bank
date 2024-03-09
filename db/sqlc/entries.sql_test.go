package simplebank

import (
	"context"
	"testing"
	"time"

	"github.com/harbouli/simplebank/utils"
	"github.com/stretchr/testify/require"
)

func CreateRandomEntry(t *testing.T) Entry {

	account := CreateRandomAccount(t)

	arg := CreateEntryParams{AccountID: account.ID, Amount: utils.RandomMoney()}

	entry, err := testQueries.CreateEntry(context.Background(), arg)
	require.NoError(t, err)

	require.NotEmpty(t, entry)

	require.Equal(t, arg.Amount, entry.Amount)
	require.Equal(t, arg.AccountID, entry.AccountID)
	require.Equal(t, account.ID, entry.AccountID)

	require.NotZero(t, entry.ID)
	require.NotZero(t, entry.CreatedAt)
	return entry
}

func TestCreateEntry(t *testing.T) {
	CreateRandomEntry(t)
}

func TestGetEntry(t *testing.T) {
	newEntry := CreateRandomEntry(t)

	entry, err := testQueries.GetEntry(context.Background(), newEntry.ID)

	require.NoError(t, err)
	require.NotEmpty(t, entry)

	require.Equal(t, entry.ID, newEntry.ID)
	require.Equal(t, entry.Amount, newEntry.Amount)
	require.Equal(t, entry.AccountID, newEntry.AccountID)

	entryCreatedAtTime := timestamptzToTime(entry.CreatedAt)
	entry2CreatedAtTime := timestamptzToTime(entry.CreatedAt)

	require.WithinDuration(t, entryCreatedAtTime, entry2CreatedAtTime, time.Second)

	require.NotZero(t, entry.ID)
	require.NotZero(t, entry.CreatedAt)

}
func TestUpdateEntry(t *testing.T) {
	newEntry := CreateRandomEntry(t)

	arg := UpdatEentryParams{
		ID:     newEntry.ID,
		Amount: 190,
	}

	entry, err := testQueries.UpdatEentry(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, entry)

	require.Equal(t, entry.ID, newEntry.ID)
	require.Equal(t, arg.Amount, entry.Amount)
	require.Equal(t, entry.AccountID, newEntry.AccountID)

	entryCreatedAtTime := timestamptzToTime(entry.CreatedAt)
	entry2CreatedAtTime := timestamptzToTime(entry.CreatedAt)

	require.WithinDuration(t, entryCreatedAtTime, entry2CreatedAtTime, time.Second)

	require.NotZero(t, entry.ID)
	require.NotZero(t, entry.CreatedAt)

}

func TestDeleteEntry(t *testing.T) {
	newEntry := CreateRandomEntry(t)

	err := testQueries.DeleteEntries(context.Background(), newEntry.ID)

	require.NoError(t, err)
	entry, err := testQueries.GetEntry(context.Background(), newEntry.ID)

	require.Error(t, err)
	// require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, entry)

}

func TestListEntries(t *testing.T) {
	for i := 0; i < 10; i++ {
		CreateRandomEntry(t)
	}
	arg := ListEntriesParams{
		Limit:  5,
		Offset: 5,
	}
	entries, err := testQueries.ListEntries(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, entries)
	require.Len(t, entries, 5)
}
