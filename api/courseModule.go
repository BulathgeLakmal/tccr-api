package api

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
	db "github.com/it21152832/Learning-Backend/db/sqlc"
	"github.com/lib/pq"
)

type createCourseModuleRequest struct {
	CourseID   int64 `json:"course_id" binding:"required"`
	ModuleName string        `json:"module_name" binding:"required"`
}

func (server *Server) createCourseModule(ctx *gin.Context) {
	var req createCourseModuleRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.CreateCourseModuleParams{
		CourseID:   req.CourseID,
		ModuleName: req.ModuleName,
	}

	courseModule, err := server.store.CreateCourseModule(ctx, arg)
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

	ctx.JSON(http.StatusOK, courseModule)
}

type getCourseModuleRequest struct {
	ModuleID int64 `uri:"module_id" binding:"required,min=1"`
}

func (server *Server) getCourseModule(ctx *gin.Context) {
	var req getCourseModuleRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	courseModule, err := server.store.GetCourseModule(ctx, req.ModuleID)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, courseModule)
}

// type deleteCourseModuleRequest struct {
// 	ModuleID int64 `uri:"module_id" binding:"required,min=1"`
// }

// func (server *Server) deleteCourseModule(ctx *gin.Context) {
// 	var req CourseModuleRequest
// 	if err := ctx.ShouldBindUri(&req); err != nil {
// 		ctx.JSON(http.StatusBadRequest, errorResponse(err))
// 		return
// 	}

// 	courseModule, err := server.store.GetCourseModule(ctx, req.ModuleID)
// 	if err != nil {
// 		if err == sql.ErrNoRows {
// 			ctx.JSON(http.StatusNotFound, errorResponse(err))
// 			return
// 		}
// 		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
// 		return
// 	}

// 	ctx.JSON(http.StatusOK, courseModule)
// }
