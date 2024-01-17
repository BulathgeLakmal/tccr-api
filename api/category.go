package api

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
	db "github.com/it21152832/Learning-Backend/db/sqlc"
	"github.com/lib/pq"
)

type createCategoryRequest struct {
	CategoryName string `json:"category_name"  binding:"required"`
	CategoryDesc string `json:"category_desc"  binding:"required"`
}

func (server *Server) createCategory(ctx *gin.Context) {
	var req createCategoryRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.CreateCategoryParams{
		CategoryName:    req.CategoryName,
		CategoryDesc: req.CategoryDesc,
	}

	category, err := server.store.CreateCategory(ctx, arg)
	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok{
			switch pqErr.Code.Name(){
			case "foreign_key_violation","unique_violation":
				ctx.JSON(http.StatusForbidden, errorResponse(err))
				return
			}
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, category)
}


type getCategoryRequest struct {
	CategoryID int64 `uri:"category_id" binding:"required,min=1"`
}

func (server *Server) getCategory(ctx *gin.Context) {
	var req getCategoryRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	category, err := server.store.GetCategory(ctx, req.CategoryID)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, category)
}

// type ListAccountRequest struct {
// 	pageID   int32 `form:"page_id" binding:"required,min=1"`
// 	pageSize int32 `form:"page_id" binding:"required,min=5,max=10"`
// }

// func (server *Server) ListAccount(ctx *gin.Context) {
// 	var req ListAccountRequest
// 	if err := ctx.ShouldBindQuery(&req); err != nil {
// 		ctx.JSON(http.StatusBadRequest, errorResponse(err))
// 		return
// 	}
// 	arg := db.ListAccountsParams{
// 		Limit:  req.pageSize,
// 		Offset: (req.pageID - 1) * req.pageSize,
// 	}

// 	accounts, err := server.store.ListAccounts(ctx, arg)
// 	if err != nil {
// 		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
// 		return

// 	}

// 	ctx.JSON(http.StatusOK, accounts)
// }
