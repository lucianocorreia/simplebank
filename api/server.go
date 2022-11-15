package api

import (
	"github.com/gin-gonic/gin"
	db "github.com/lucianocorreia/simplebank/db/sqlc"
)

// Server serves Http requests for our banking service
type Server struct {
	store  db.Store
	router *gin.Engine
}

// NewServer creates a new Http Server
func NewServer(store db.Store) *Server {
	server := &Server{
		store: store,
	}
	router := gin.Default()

	router.POST("/accounts", server.createAccount)
	router.GET("/accounts/:id", server.getAccount)
	router.GET("/accounts", server.listAccount)
	router.PUT("/accounts/:id", server.updateAccount)
	router.DELETE("/accounts/:id", server.deleteAccount)

	server.router = router
	return server
}

// Start runs the http server on the given address
func (s *Server) Start(address string) error {
	return s.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{
		"error": err.Error(),
	}
}
