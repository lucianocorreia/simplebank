package gapi

import (
	"fmt"

	db "github.com/lucianocorreia/simplebank/db/sqlc"
	"github.com/lucianocorreia/simplebank/pb"
	"github.com/lucianocorreia/simplebank/token"
	"github.com/lucianocorreia/simplebank/util"
)

// Server serves GRP requersts for our bank service
type Server struct {
	pb.UnimplementedSimpleBankServer
	config     util.Config
	store      db.Store
	tokenMaker token.Maker
}

// NewServer creates a new GRP server
func NewServer(config util.Config, store db.Store) (*Server, error) {
	tokenMaker, err := token.NewPasetoMaker(config.TokenSymmetricKey)
	if err != nil {
		return nil, fmt.Errorf("cannot create token maker: %s", err)
	}

	server := &Server{
		config:     config,
		store:      store,
		tokenMaker: tokenMaker,
	}

	return server, nil
}
