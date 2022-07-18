package go_mw_service

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-mw/util"
	"net/http"
)

type GetRandomRequest struct {
	Length int `json:"length" binding:"gt=0"`
}

type GetRandomResponse struct {
	String string `json:"string"`
}

func (server *Server) getRandom(ctx *gin.Context) {
	var err error
	var request GetRandomRequest
	var response GetRandomResponse

	if err = ctx.ShouldBindJSON(&request); err == nil {
		randomString := util.RandomStringAlphabet(request.Length)

		response = GetRandomResponse{randomString}

		fmt.Printf("%v", response)
		ctx.JSON(http.StatusOK, response)
		return
	}
}
