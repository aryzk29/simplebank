package api

import (
	"database/sql"
	"fmt"
	"net/http"

	db "github.com/aryzk29/simplebankcp/db/sqlc"
	"github.com/gin-gonic/gin"
)

type transferRequest struct {
	fromAccount int64  `json:"from_account" binding:"required,min=1"`
	toAccount   int64  `json:"to_account" binding:"required,min=1"`
	amount      int64  `json:amount binding:"required,min=1"`
	currency    string `json:currency binding:"required,currency"`
}

func (server *Server) createTransfer(ctx *gin.Context) {
	var request transferRequest

	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	if !server.validAccount(ctx, request.fromAccount, request.currency) {
		return
	}

	if !server.validAccount(ctx, request.toAccount, request.currency) {
		return
	}

	arg := db.TransferTxParams{
		FromAccountID: request.fromAccount,
		ToAccountID:   request.toAccount,
		Amount:        request.amount,
	}

	result, err := server.store.TransferTx(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, result)
}

func (server *Server) validAccount(ctx *gin.Context, accountID int64, currency string) bool {
	account, err := server.store.GetAccount(ctx, accountID)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return false
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return false
	}

	if account.Currency != currency {
		err := fmt.Errorf("account [%d] currency mismatch: %s and %s", account.ID, account.Currency, currency)
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return false
	}
	return true
}
