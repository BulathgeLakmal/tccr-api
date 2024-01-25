package api

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
	db "github.com/it21152832/Learning-Backend/db/sqlc"
	"github.com/lib/pq"
)

type createCourseRequest struct {
	CourseName string `json:"course_name" binding:"required"`
	CourseDesc string `json:"course_desc" binding:"required"`
	Category   string `json:"category" binding:"required"`
}

func (server *Server) createCourse(ctx *gin.Context) {
	var req createCourseRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.CreateCourseParams{
		CourseName:    req.CourseName,
		CourseDesc:    req.CourseDesc,
		Category:      req.Category,
	}

	course, err := server.store.CreateCourse(ctx, arg)
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

	ctx.JSON(http.StatusOK, course)
}


type getCourseRequest struct {
	CourseID int64 `uri:"course_id" binding:"required,min=1"`
}

func (server *Server) getCourse(ctx *gin.Context) {
	var req getCourseRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	course, err := server.store.GetCourse(ctx, req.CourseID)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, course)
}


type getAllCourseRequest struct {
	
}

func (server *Server) getAllCourse(ctx *gin.Context) {

	courses, err := server.store.GetAllCourse(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, courses)
}
