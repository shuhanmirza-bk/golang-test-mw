package go_mw_service

import (
	"github.com/gin-gonic/gin"
)

const PREFIX_PATH = ""

type Server struct {
	router *gin.Engine
}

func NewServer() *Server {
	server := &Server{}
	router := gin.Default()

	router.POST(PREFIX_PATH+"/random/", server.getRandom)

	server.router = router

	return server
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{
		"error": err.Error(),
	}
}
