package api

import (
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