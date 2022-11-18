package api

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	db "github.com/lucianocorreia/simplebank/db/sqlc"
	"github.com/lucianocorreia/simplebank/token"
	"github.com/lucianocorreia/simplebank/util"
)

// Server serves Http requests for our banking service
type Server struct {
	config     util.Config
	store      db.Store
	tokenMaker token.Maker
	router     *gin.Engine
}

// NewServer creates a new Http Server
func NewServer(config util.Config, store db.Store) (*Server, error) {
	tokenMaker, err := token.NewPasetoMaker(config.TokenSymmetricKey)
	if err != nil {
		return nil, fmt.Errorf("cannot create token maker %s", err)
	}
	server := &Server{
		config:     config,
		store:      store,
		tokenMaker: tokenMaker,
	}

	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("currency", validCurrency)
	}

	server.setupRoutes()

	return server, nil
}

func (s *Server) setupRoutes() {
	router := gin.Default()

	router.POST("/users", s.createUser)
	router.POST("/login", s.loginUser)

	router.POST("/accounts", s.createAccount)
	router.GET("/accounts/:id", s.getAccount)
	router.GET("/accounts", s.listAccount)
	router.PUT("/accounts/:id", s.updateAccount)
	router.DELETE("/accounts/:id", s.deleteAccount)

	router.POST("/transfers", s.createTransfer)

	s.router = router
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
