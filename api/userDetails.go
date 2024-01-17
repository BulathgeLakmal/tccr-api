package api

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
	db "github.com/it21152832/Learning-Backend/db/sqlc"
	"github.com/lib/pq"
)

type createUserDetailsRequest struct {
	UserID       int64  `json:"user_id" binding:"required"`
	Phone        int32  `json:"phone" binding:"required"`
	AddressLine1 string `json:"address_line1" binding:"required"`
	AddressLine2 string `json:"address_line2" binding:"required"`
}

func (server *Server) createUserDetails(ctx *gin.Context) {
	var req createUserDetailsRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.CreateUserDetailsParams{
		UserID:       req.UserID,
		Phone:        req.Phone,
		AddressLine1: req.AddressLine1,
		AddressLine2: req.AddressLine2,
	}

	userDetails, err := server.store.CreateUserDetails(ctx, arg)
	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok {
			switch pqErr.Code.Name() {
			case "foreign_key_violation", "unique_violation":
				ctx.JSON(http.StatusForbidden, errorResponse(err))
				return
			}
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, userDetails)
}

type getUserDetailsRequest struct {
	UserDetailsID int64 `uri:"user_details_id" binding:"required,min=1"`
}

func (server *Server) getUserDetails(ctx *gin.Context) {
	var req getUserDetailsRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	userDetails, err := server.store.GetUser(ctx, req.UserDetailsID)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, userDetails)
}
