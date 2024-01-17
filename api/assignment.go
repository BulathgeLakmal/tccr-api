package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	db "github.com/it21152832/Learning-Backend/db/sqlc"
	"github.com/lib/pq"
)

type createAssignmentRequest struct {
	UserID         int64 `json:"user_id" binding:"required"`
	CourseModuleID int64 `json:"course_module_id" binding:"required"`
	LectureID      int64 `json:"lecture_id" binding:"required"`
}

func (server *Server) createAssignment(ctx *gin.Context) {
	var req createAssignmentRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.CreateAssignmentParams{
		UserID:         req.UserID,
		CourseModuleID: req.CourseModuleID,
		LectureID:      req.LectureID,
	}

	assignment, err := server.store.CreateAssignment(ctx, arg)
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

	ctx.JSON(http.StatusOK, assignment)
}

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
