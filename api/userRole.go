package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	db "github.com/it21152832/Learning-Backend/db/sqlc"
	"github.com/lib/pq"
)

type createUserRoleRequest struct {
	Role string `json:"role" binding:"required"`
}

func (server *Server) createUserRole(ctx *gin.Context) {
	var req createUserRoleRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.CreateUserRoleParams{
		Role:    req.Role,
	}

	userRole, err := server.store.CreateUserRole(ctx, arg)
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

	ctx.JSON(http.StatusOK, userRole)
}


// type getCourseRequest struct {
// 	CourseID int64 `uri:"course_id" binding:"required,min=1"`
// }

// func (server *Server) getCourse(ctx *gin.Context) {
// 	var req getCourseRequest
// 	if err := ctx.ShouldBindUri(&req); err != nil {
// 		ctx.JSON(http.StatusBadRequest, errorResponse(err))
// 		return
// 	}

// 	course, err := server.store.GetCourse(ctx, req.CourseID)
// 	if err != nil {
// 		if err == sql.ErrNoRows {
// 			ctx.JSON(http.StatusNotFound, errorResponse(err))
// 			return
// 		}
// 		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
// 		return
// 	}

// 	ctx.JSON(http.StatusOK, course)
// }

