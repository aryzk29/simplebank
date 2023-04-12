package db

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestTransferTx(t *testing.T) {
	store := NewStore(testDB)

	account1 := createRandomAccount(t)
	account2 := createRandomAccount(t)
	fmt.Println(">> Before:", account1.Balance, account2.Balance)

	// run n concurrent transfer transaction
	n := 10
	amount := int64(10)

	errs := make(chan error)

	for i := 0; i < n ; i++ {

		//add deadlock test case
		fromAccountID := account1.ID
		toAccountID := account2.ID

		if i % 2 == 1 {
			fromAccountID = account2.ID
			toAccountID = account1.ID
		}

		go func ()  {
			_, err := store.TransferTx(context.Background(), TransferTxParams{
				FromAccountID: fromAccountID,
				ToAccountID: toAccountID,
				Amount: amount,
			})

			errs <- err
		} ()
	}

	//check results
	for i := 0; i < n; i++ {
		err := <-errs
		require.NoError(t,err)
	}
	
	//check final updated balance
	updatedAccount1, err := testQueries.GetAccount(context.Background(), account1.ID)
	require.NoError(t, err)
	
	updatedAccount2, err := testQueries.GetAccount(context.Background(), account2.ID)
	require.NoError(t, err)

	fmt.Println(">> After:", updatedAccount1.Balance, updatedAccount2.Balance)
	
		require.Equal(t, account1.Balance, updatedAccount1.Balance)
		require.Equal(t, account2.Balance, updatedAccount2.Balance)
}