package go_mw_service

import (
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
)

func (server *Server) httpGetRequest(url string) (string, error) {
	response, err := http.Get(url)
	if err != nil {
		return "", err
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return "", err
	}

	return string(responseData), nil
}

func (server *Server) getHelloWorld(ctx *gin.Context) {
	var err error

	url := GLOBAL_CONFIG.TargetAddress + "/hello-world"
	response, err := server.httpGetRequest(url)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, response)
}

func (server *Server) getHelloWorldDelayed(ctx *gin.Context) {
	var err error

	url := GLOBAL_CONFIG.TargetAddress + "/hello-world/delayed"
	response, err := server.httpGetRequest(url)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, response)
}

func (server *Server) getHelloWorldVeryDelayed(ctx *gin.Context) {
	var err error

	url := GLOBAL_CONFIG.TargetAddress + "/hello-world/very-delayed"
	response, err := server.httpGetRequest(url)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, response)
}
