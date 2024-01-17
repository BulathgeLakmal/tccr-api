package api

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
	db "github.com/it21152832/Learning-Backend/db/sqlc"
	"github.com/lib/pq"
)

type createLectureRequest struct {
	CourseModuleID int64  `json:"course_module_id"`
	LectureDesc    string `json:"lecture_desc"  binding:"required"`
	LectureNumber  int32  `json:"lecture_number"  binding:"required"`
	VideoURL       string `json:"video_URL" binding:"required"`
	Status         string `json:"status"  binding:"required"`
}

func (server *Server) createLecture(ctx *gin.Context) {
	var req createLectureRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.CreateLectureParams{
		CourseModuleID: req.CourseModuleID,
		LectureDesc:    req.LectureDesc,
		LectureNumber:  req.LectureNumber,
		VideoURL:       req.VideoURL,
		Status:         req.Status,
	}

	lectures, err := server.store.CreateLecture(ctx, arg)
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

	ctx.JSON(http.StatusOK, lectures)
}

type getLectureRequest struct {
	LectureID int64 `uri:"lecture_id" binding:"required,min=1"`
}

func (server *Server) getLecture(ctx *gin.Context) {
	var req getLectureRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	lectures, err := server.store.GetLecture(ctx, req.LectureID)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, lectures)
}
