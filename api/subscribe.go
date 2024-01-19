package api

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
	db "github.com/it21152832/Learning-Backend/db/sqlc"
	"github.com/lib/pq"
)

type createSubscribeRequest struct {
	UserID   int64 `json:"user_id" binding:"required"`
	CourseID int64 `json:"course_id" binding:"required"`
}

type subscribeResponse struct {
	CourseID int64 `json:"course_id"`
}

func newSubscribeResponse(subscribe db.Subscribe) subscribeResponse {
	return subscribeResponse{
		CourseID: subscribe.CourseID,
	}
}

type subscribedUserRequest struct {
	UserID int64 `json:"user_id" binding:"required"`
}

type subscribedUserResponse struct {
	CourseID int64 `json:"course_id"`
}

func (server *Server) createSubscribe(ctx *gin.Context) {
	var req createSubscribeRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		handleError(ctx, http.StatusBadRequest, err)
		return
	}

	arg := db.CreateSubscribeParams{
		UserID:   req.UserID,
		CourseID: req.CourseID,
	}

	subscribe, err := server.store.CreateSubscribe(ctx, arg)
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

	rsp := newSubscribeResponse(subscribe)
	ctx.JSON(http.StatusOK, rsp)
}

func (server *Server) subscribedUser(ctx *gin.Context) {
	var req subscribedUserRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		handleError(ctx, http.StatusBadRequest, err)
		return
	}

	subscribe, err := server.store.GetSubscribe(ctx, req.UserID)
	if err != nil {
		if err == sql.ErrNoRows {
			handleError(ctx, http.StatusNotFound, err)
			return
		}
		handleError(ctx, http.StatusInternalServerError, err)
		return
	}

	rsp := newSubscribeResponse(subscribe)
	ctx.JSON(http.StatusOK, rsp)
}




// package api

// import (
// 	"database/sql"
// 	"net/http"

// 	"github.com/gin-gonic/gin"
// 	db "github.com/it21152832/Learning-Backend/db/sqlc"
// 	"github.com/lib/pq"
// )

// type createSubscribeRequest struct {
// 	UserID         int64 `json:"user_id" binding:"required"`
// 	CourseID int64 `json:"course_id" binding:"required"`
// }

// type subscribeResponse struct{
// 	CourseID int64 `json:"course_id"`
// }

// func newSubscribeResponse(subscribe db.Subscribe) subscribeResponse{
// 	return subscribeResponse{
// 		CourseID : subscribe.CourseID,
// 	}
// }


// type subscribedUserRequest struct{
// 	UserID int64 `json:"user_id" binding:"required"`
// }

// type subscribedUserResponse struct{
// 	CourseID int64 `json:"course_id"`
// }

// func (server *Server) createSubscribe(ctx *gin.Context) {
// 	var req createSubscribeRequest
// 	if err := ctx.ShouldBindJSON(&req); err != nil {
// 		handleError(ctx, http.StatusBadRequest, err)
// 		return
// 	}

// 	arg := db.CreateSubscribeParams{
// 		UserID:       req.UserID,
// 		CourseID:        req.CourseID,
// 	}

// 	subscribe, err := server.store.CreateSubscribe(ctx, arg)
// 	if err != nil {
// 		if pqErr, ok := err.(*pq.Error); ok {
// 			switch pqErr.Code.Name() {
// 			case "foreign_key_violation", "unique_violation":
// 				handleError(ctx, http.StatusForbidden, err)
// 				return
// 			}
// 		}
// 		handleError(ctx, http.StatusInternalServerError, err)
// 		return
// 	}

// 	rsp := subscribeResponse{
// 		Subscribe: newSubscribeResponse(subscribe),
// 	}
// 	ctx.JSON(http.StatusOK, rsp)
// }


// func (server *Server) subscribedUser(ctx *gin.Context) {
// 	var req subscribedUserRequest

// 	if err := ctx.ShouldBindJSON(&req); err != nil {
// 		handleError(ctx, http.StatusBadRequest, err)
// 		return
// 	}

// 	subscribe, err := server.store.GetSubscribe(ctx, req.UserID)
// 	if err != nil {
// 		if err == sql.ErrNoRows {
// 			handleError(ctx, http.StatusNotFound, err)
// 			return
// 		}
// 		handleError(ctx, http.StatusInternalServerError, err)
// 		return
// 	}

// 	rsp := subscribedUserResponse {
// 		Subscribe: newSubscribeResponse(subscribe),
// 	}
// 	ctx.JSON(http.StatusOK, rsp)
// }

// func handleError(ctx *gin.Context, code int, err error) {
// 	ctx.JSON(code, errorResponse(err))
// }

// type getUserRequest struct {
// 	UserID int64 `uri:"user_id" binding:"required,min=1"`
// }

// func (server *Server) getUser(ctx *gin.Context) {
// 	var req getUserRequest
// 	if err := ctx.ShouldBindUri(&req); err != nil {
// 		ctx.JSON(http.StatusBadRequest, errorResponse(err))
// 		return
// 	}

// 	user, err := server.store.GetUser(ctx, req.UserID)
// 	if err != nil {
// 		if err == sql.ErrNoRows {
// 			ctx.JSON(http.StatusNotFound, errorResponse(err))
// 			return
// 		}
// 		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
// 		return
// 	}

// 	ctx.JSON(http.StatusOK, user)
// }
