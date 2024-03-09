package simplebank

// import (
// 	"context"
// 	"fmt"

// 	"github.com/jackc/pgx/v5"
// )

// type Store struct {
// 	*Queries
// 	db *pgx.Conn
// }

// func NewStore(db *pgx.Conn) *Store {

// 	return &Store{
// 		db:      db,
// 		Queries: New(db),
// 	}
// }

// // execT executes a function within a database transaction
// func (store *Store) execTx(ctx context.Context, fn func(*Queries) error) error {
// 	tx, err := store.db.BeginTx(ctx, pgx.TxOptions{})
// 	if err != nil {
// 		return err
// 	}

// 	q := New(tx)
// 	err = fn(q)

// 	if err != nil {
// 		if rbErr := tx.Rollback(ctx); rbErr != nil {
// 			fmt.Errorf("tx error %v , rb erro %v", err, rbErr)
// 		}
// 		return err
// 	}
// 	return tx.Commit(ctx)
// }

// // TransferTxParams contains the input parameters of the transfer transaction
// type TransfeTxParams struct {
// 	FromAccountID int64 `json:"form_account_id"`
// 	ToAccountID   int64 `json:"to_account_id"`
// 	Amount        int64 `json:"amount"`
// }

// // TransferTxResult is the result of the transfer transaction
// type TransferTxResult struct {
// 	Transfer    Transfer `json:"transfer"`
// 	FromAccount Account  `json:"from_account"`
// 	ToAccount   Account  `json:"to_account"`
// 	FromEntry   Entry    `json:" from_entry"`
// 	ToEntry     Entry    `json:" to_entry"`
// }

// // TransferTx performs a money transfer from one account to the other.
// // It creates a transfer record, add account entries, and update accounts'balance within a single database transaction
// // func (store *Store) TransferTx(ctx context.Context, arg TransfeTxParams) (TransferTxResult, error) {
// // 	var result TransferTxResult
// // 	err := store.execTx(ctx, func(q *Queries) error {
// // 		result.Transfer, err = q.
// // 		return nil
// // 	})
// // }
