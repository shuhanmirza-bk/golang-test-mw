package go_mw_service

import (
	"github.com/gin-gonic/gin"
	"go-mw/util"
)

const PREFIX_PATH = ""

var GLOBAL_CONFIG util.GlobalConfig

type Server struct {
	router *gin.Engine
}

func NewServer(config util.GlobalConfig) *Server {
	GLOBAL_CONFIG = config
	server := &Server{}
	router := gin.Default()

	router.POST(PREFIX_PATH+"/random/", server.getRandom)
	router.GET(PREFIX_PATH+"/hello-world/", server.getHelloWorld)
	router.GET(PREFIX_PATH+"/hello-world/delayed", server.getHelloWorldDelayed)
	router.GET(PREFIX_PATH+"/hello-world/very-delayed", server.getHelloWorldVeryDelayed)

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
