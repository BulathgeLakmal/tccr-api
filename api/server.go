// package api

// import (
// 	"fmt"
// 	"time"

// 	"github.com/gin-contrib/cors"
// 	"github.com/gin-gonic/gin"
// 	db "github.com/it21152832/Learning-Backend/db/sqlc"
// 	"github.com/it21152832/Learning-Backend/token"
// 	"github.com/it21152832/Learning-Backend/util"
// )

// // Server serves HTTP requests for our banking service
// type Server struct {
// 	config     util.Config
// 	store      *db.Store
// 	tokenMaker token.Maker
// 	router     *gin.Engine
// }

// // NewServer creates a new HTTP server and sets up routing.
// func NewServer(config util.Config, store *db.Store) (*Server, error) {
// 	tokenMaker, err := token.NewPasetoMaker(config.TokenSymmetricKey)
// 	if err != nil {
// 		return nil, fmt.Errorf("cannot create token maker: %v", err)
// 	}
// 	server := &Server{
// 		config:     config,
// 		store:      store,
// 		tokenMaker: tokenMaker,
// 	}

// 	router := gin.Default()

// 	//API calls for category
// 	router.POST("/category", server.createCategory)
// 	router.GET("/category/:category_id", server.getCategory)

// 	//API calls for category
// 	router.POST("/course", server.createCourse)
// 	router.GET("/course/:course_id", server.getCourse)

// 	//API calls for course module
// 	router.POST("/courseModule", server.createCourseModule)
// 	router.GET("/courseModule/:module_id", server.getCourseModule)

// 	//API calls for lectures
// 	router.POST("/lectures", server.createLecture)
// 	router.GET("/lectures/:lecture_id", server.getLecture)

// 	//API calls for user role
// 	router.POST("/userRole", server.createUserRole)

// 	//API calls for user
// 	router.POST("/user", server.createUser)
// 	router.POST("/user/login", server.loginUser)
// 	// router.GET("/user/:user_id", server.getUser)

// 	//API calls for user details
// 	router.POST("/userDetails", server.createUserDetails)
// 	router.POST("/userDetails/details", server.detailsUser)
// 	// router.GET("/userDetails/:user_details_id", server.getUserDetails)

// 	//API calls for assignment
// 	router.POST("/assignment", server.createAssignment)
// 	router.POST("/assignment/assignmentDetails", server.assignmentUser)
// 	// router.GET("/assignment/:assignment_id", server.getUserDetails)

// 	//API calls for assignment file
// 	router.POST("/assignment_file", server.createAssignmentFile)

// 	//API calls for subscribe
// 	router.POST("/subscribe", server.createSubscribe)
// 	router.POST("/subscribe/subscribedUsers", server.subscribedUser)

// 	router.Use(cors.New(cors.Config{
// 		AllowOrigins:     []string{"http://", "https://", ""},
// 		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
// 		AllowHeaders:     []string{""},
// 		ExposeHeaders:    []string{"Content-Length"},
// 		AllowCredentials: true,
// 		MaxAge:           12 * time.Hour,
// 	}))

// 	server.router = router
// 	// server.setupRouter()
// 	return server, nil // Return server with nil error as we've handled error in tokenMaker creation
// }

// // Start runs the HTTP server on a specific address.
// func (server *Server) Start(address string) error {
// 	return server.router.Run(address)
// }

// func errorResponse(err error) gin.H {
// 	return gin.H{"error": err.Error()}
// }

package api

import (
	"fmt"
	"time"

	"github.com/gin-contrib/cors"
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

	// router.Use(cors.New(cors.Config{
	// 	AllowOrigins:     []string{"http://localhost:3000"},
	// 	AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
	// 	AllowHeaders:     []string{"Origin", "Content-Type"},
	// 	ExposeHeaders:    []string{"Content-Length"},
	// 	AllowCredentials: true,
	// 	MaxAge:           12 * time.Hour,
	// }))

	// Start the server
	// router.Run(":8080")

	// API calls for category
	router.POST("/category", server.createCategory)
	router.GET("/category/:category_id", server.getCategory)

	// API calls for course
	router.POST("/course", server.createCourse)
	router.GET("/course/:course_id", server.getCourse)
	router.GET("/course", server.getAllCourse)

	// API calls for course module
	router.POST("/courseModule", server.createCourseModule)
	router.GET("/courseModule/:module_id", server.getCourseModule)
	router.GET("/courseModule", server.GetAllCourseModule)

	// API calls for lectures
	router.POST("/lectures", server.createLecture)
	router.GET("/lectures/:lecture_id", server.getLecture)

	// API calls for user role
	router.POST("/userRole", server.createUserRole)

	// API calls for user
	router.POST("/user", server.createUser)
	router.POST("/user/login", server.loginUser)

	// API calls for user details
	router.POST("/userDetails", server.createUserDetails)
	router.POST("/userDetails/details", server.detailsUser)

	// API calls for assignment
	router.POST("/assignment", server.createAssignment)
	router.POST("/assignment/assignmentDetails", server.assignmentUser)

	// API calls for assignment file
	router.POST("/assignment_file", server.createAssignmentFile)

	// API calls for subscribe
	router.POST("/subscribe", server.createSubscribe)
	router.POST("/subscribe/subscribedUsers", server.subscribedUser)

	    // Set up CORS middleware

		router.Use(cors.New(cors.Config{
			AllowOrigins:     []string{"http://localhost:3000"},
			AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
			AllowHeaders:     []string{"Origin", "Content-Type"},
			ExposeHeaders:    []string{"Content-Length"},
			AllowCredentials: true,
			MaxAge:           12 * time.Hour,
		}))
	
		// Start the server
		router.Run(":8080")

	server.router = router
	return server, nil
}

// Start runs the HTTP server on a specific address.
func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
