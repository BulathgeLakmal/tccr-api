package api

import (
	"github.com/gin-gonic/gin"
	db "github.com/it21152832/Learning-Backend/db/sqlc"
	"github.com/it21152832/Learning-Backend/util"
)

// Server serves HTTP requests for our banking service
type Server struct {
	config     util.Config
	store      *db.Store
	router     *gin.Engine
}

// NewServer creates a new HTTP server and sets up routing.
func NewServer(config util.Config , store *db.Store) (*Server, error) {


	server := &Server{
		store:      store,
	}

	router := gin.Default()
	router.POST("/category", server.createCategory)

	server.router = router
	// server.setupRouter()
	return server, nil// Return server with nil error as we've handled error in tokenMaker creation
}

// Start runs the HTTP server on a specific address.
func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
