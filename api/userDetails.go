// package api

// import (
// 	"database/sql"
// 	"net/http"

// 	"github.com/gin-gonic/gin"
// 	db "github.com/it21152832/Learning-Backend/db/sqlc"
// 	"github.com/lib/pq"
// )

// type createUserDetailsRequest struct {
// 	UserID       int64 `json:"user_id" binding:"required"`
// 	Phone        int32         `json:"phone" binding:"required"`
// 	AddressLine1 string        `json:"address_line1" binding:"required"`
// 	AddressLine2 string        `json:"address_line2" binding:"required"`
// }

// func (server *Server) createUserDetails(ctx *gin.Context) {
// 	var req createUserDetailsRequest
// 	if err := ctx.ShouldBindJSON(&req); err != nil {
// 		ctx.JSON(http.StatusBadRequest, errorResponse(err))
// 		return
// 	}

// 	arg := db.CreateUserDetailsParams{
// 		UserID:   req.UserID,
// 		Phone: req.Phone,
// 		AddressLine1: req.AddressLine1,
// 		AddressLine2:req.AddressLine2,
// 	}

// 	userDetails, err := server.store.CreateUserDetails(ctx, arg)
// 	if err != nil {
// 		if pqErr, ok := err.(*pq.Error); ok {
// 			switch pqErr.Code.Name() {
// 			case "foreign_key_violation", "unique_violation":
// 				ctx.JSON(http.StatusForbidden, errorResponse(err))
// 				return
// 			}
// 		}
// 		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
// 		return
// 	}

// 	ctx.JSON(http.StatusOK, userDetails)
// }

// type getUserDetailsRequest struct {
// 	UserDetailsID int64 `uri:"user_details_id" binding:"required,min=1"`
// }

// func (server *Server) getUserDetails(ctx *gin.Context) {
// 	var req getUserDetailsRequest
// 	if err := ctx.ShouldBindUri(&req); err != nil {
// 		ctx.JSON(http.StatusBadRequest, errorResponse(err))
// 		return
// 	}

// 	userDetails, err := server.store.GetUserDetails(ctx, req.UserDetailsID)
// 	if err != nil {
// 		if err == sql.ErrNoRows {
// 			ctx.JSON(http.StatusNotFound, errorResponse(err))
// 			return
// 		}
// 		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
// 		return
// 	}

// 	ctx.JSON(http.StatusOK, userDetails)
// }

// type deleteUserDetailsRequest struct {
// 	UserDetailsID int64 `uri:"user_details_id" binding:"required,min=1"`
// }

// func (server *Server) deleteUserDetails(ctx *gin.Context) {
// 	var req deleteUserDetailsRequest
// 	if err := ctx.ShouldBindUri(&req); err != nil {
// 		ctx.JSON(http.StatusBadRequest, errorResponse(err))
// 		return
// 	}

// 	userDetails, err := server.store.GetUserDetails(ctx, req.UserDetailsID)
// 	if err != nil {
// 		if err == sql.ErrNoRows {
// 			ctx.JSON(http.StatusNotFound, errorResponse(err))
// 			return
// 		}
// 		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
// 		return
// 	}

// 	ctx.JSON(http.StatusOK, userDetails)
// }

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

type userDetailsResponse struct {
	Phone        int32  `json:"phone"`
	AddressLine1 string `json:"address_line1"`
	AddressLine2 string `json:"address_line2"`
}

func newUserDetailsResponse(userDetail db.UserDetail) userDetailsResponse {
	return userDetailsResponse{
		Phone:        userDetail.Phone,
		AddressLine1: userDetail.AddressLine1,
		AddressLine2: userDetail.AddressLine2,
	}
}

type detailsRequest struct {
	UserID int64 `json:"user_id" binding:"required"`
}

type detailsResponse struct {
	UserDetail userDetailsResponse `json:"userDetails"`
}

func (server *Server) createUserDetails(ctx *gin.Context) {
	var req createUserDetailsRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		handleError(ctx, http.StatusBadRequest, err)
		return
	}

	arg := db.CreateUserDetailsParams{
		UserID:       req.UserID,
		Phone:        req.Phone,
		AddressLine1: req.AddressLine1,
		AddressLine2: req.AddressLine2,
	}

	userDetail, err := server.store.CreateUserDetails(ctx, arg)
	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok {
			switch pqErr.Code.Name() {
			case "foreign_key_violation", "unique_violation":
				handleError(ctx, http.StatusForbidden, err)
				return
			}
		}
		handleError(ctx, http.StatusInternalServerError, err)
		return
	}

	rsp := detailsResponse{
		UserDetail: newUserDetailsResponse(userDetail),
	}
	ctx.JSON(http.StatusOK, rsp)
}

func (server *Server) detailsUser(ctx *gin.Context) {
	var req detailsRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		handleError(ctx, http.StatusBadRequest, err)
		return
	}

	userDetail, err := server.store.GetUserDetails(ctx, req.UserID)
	if err != nil {
		if err == sql.ErrNoRows {
			handleError(ctx, http.StatusNotFound, err)
			return
		}
		handleError(ctx, http.StatusInternalServerError, err)
		return
	}

	rsp := detailsResponse{
		UserDetail: newUserDetailsResponse(userDetail),
	}
	ctx.JSON(http.StatusOK, rsp)
}

// type getUserDetailsRequest struct {
// 	UserDetailsID int64 `uri:"user_details_id" binding:"required,min=1"`
// }

// func (server *Server) getUserDetails(ctx *gin.Context) {
// 	var req getUserDetailsRequest
// 	if err := ctx.ShouldBindUri(&req); err != nil {
// 		ctx.JSON(http.StatusBadRequest, errorResponse(err))
// 		return
// 	}

// 	userDetails, err := server.store.GetUserDetails(ctx, req.UserDetailsID)
// 	if err != nil {
// 		if err == sql.ErrNoRows {
// 			ctx.JSON(http.StatusNotFound, errorResponse(err))
// 			return
// 		}
// 		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
// 		return
// 	}

// 	ctx.JSON(http.StatusOK, userDetails)
// }

// type deleteUserDetailsRequest struct {
// 	UserDetailsID int64 `uri:"user_details_id" binding:"required,min=1"`
// }

// func (server *Server) deleteUserDetails(ctx *gin.Context) {
// 	var req deleteUserDetailsRequest
// 	if err := ctx.ShouldBindUri(&req); err != nil {
// 		ctx.JSON(http.StatusBadRequest, errorResponse(err))
// 		return
// 	}

// 	userDetails, err := server.store.GetUserDetails(ctx, req.UserDetailsID)
// 	if err != nil {
// 		if err == sql.ErrNoRows {
// 			ctx.JSON(http.StatusNotFound, errorResponse(err))
// 			return
// 		}
// 		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
// 		return
// 	}

// 	ctx.JSON(http.StatusOK, userDetails)
// }
