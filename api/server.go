package api

import (
	"fmt"

	"github.com/gin-gonic/gin"
	db "github.com/it21152832/Learning-Backend/db/sqlc"
	"github.com/it21152832/Learning-Backend/token"
	"github.com/it21152832/Learning-Backend/util"
)

// Server serves HTTP requests for our banking service
type Server struct {
	config     util.Config
	store      *db.Store
	tokenMaker token.Maker
	router     *gin.Engine
}

// NewServer creates a new HTTP server and sets up routing.
func NewServer(config util.Config, store *db.Store) (*Server, error) {
	tokenMaker, err := token.NewPasetoMaker(config.TokenSymmetricKey)
	if err != nil {
		return nil, fmt.Errorf("cannot create token maker: %v", err)
	}
	server := &Server{
		config:     config,
		store:      store,
		tokenMaker: tokenMaker,
	}

	router := gin.Default()

	//API calls for category
	router.POST("/category", server.createCategory)
	router.GET("/category/:category_id", server.getCategory)

	//API calls for category
	router.POST("/course", server.createCourse)
	router.GET("/course/:course_id", server.getCourse)

	//API calls for course module
	router.POST("/courseModule", server.createCourseModule)
	router.GET("/courseModule/:module_id", server.getCourseModule)

	//API calls for lectures
	router.POST("/lectures", server.createLecture)
	router.GET("/lectures/:lecture_id", server.getLecture)

	//API calls for user role
	router.POST("/userRole", server.createUserRole)

	//API calls for user
	router.POST("/user", server.createUser)
	router.POST("/user/login", server.loginUser)
	// router.GET("/user/:user_id", server.getUser)

	//API calls for user details
	// router.POST("/userDetails", server.createUserDetails)
	// router.GET("/userDetails/:user_details_id", server.getUserDetails)

	//API calls for assignment
	router.POST("/assignment", server.createAssignment)
	// router.GET("/assignment/:assignment_id", server.getUserDetails)

	//API calls for assignment file
	router.POST("/assignment_file", server.createAssignmentFile)

	//API calls for subscribe
	router.POST("/subscribe", server.createSubscribe)

	server.router = router
	// server.setupRouter()
	return server, nil // Return server with nil error as we've handled error in tokenMaker creation
}

// Start runs the HTTP server on a specific address.
func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
