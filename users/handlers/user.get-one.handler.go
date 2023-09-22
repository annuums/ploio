package userhandlers

import (
	"log"
	"net/http"
	"strconv"

	"github.com/annuums/solanum"
	"github.com/gin-gonic/gin"
)

func NewUserGetOneHandler() *solanum.SolaService {
	return &solanum.SolaService{
		Uri:        "/:id",
		Method:     http.MethodGet,
		Handler:    getOneHandler,
		Middleware: nil,
	}
}

func getOneHandler(ctx *gin.Context) {
	//*id := ctx.Query("lastname") // shortcut for ctx.Request.URL.Query().Get("lastname")
	//*id := ctx.Param("id")
	id := ctx.Param("id")

	log.Println("id: ", id)
	code := http.StatusOK
	if _, err := strconv.Atoi(id); err != nil {
		code = http.StatusBadRequest
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code":    code,
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code": code,
		"id":   id,
	})
}
