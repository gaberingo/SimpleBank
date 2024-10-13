package api

import (
	"net/http"

	db "github.com/gaberingo/SimpleBank/db/sqlc"
	"github.com/gin-gonic/gin"
)

type getlistTransferParams struct {
	FromID    int64 `form:"from_account" binding:"required,min=1"`
	ToID      int64 `form:"to_account" binding:"required,min=1"`
	TxLimit   int32 `form:"tx_limit" binding:"required,min=5,max=10"`
	TxOffsite int32 `form:"tx_offsite" binding:"required,min=1"`
}

func (server *Server) getListTransfer(ctx *gin.Context) {
	var req getlistTransferParams

	err := ctx.ShouldBindQuery(&req)

	if req.FromID == req.ToID {
		ctx.JSON(http.StatusBadRequest, "From and To account cannot be same")
	}
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.ListTransfersParams{
		FromAccountID: req.FromID,
		ToAccountID:   req.ToID,
		Limit:         req.TxLimit,
		Offset:        (req.TxOffsite - 1) * req.TxLimit,
	}

	listTransfer, err := server.store.ListTransfers(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, listTransfer)
}
