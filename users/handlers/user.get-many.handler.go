package userhandlers

import (
	"net/http"

	usermodel "github.com/annuums/ploio/users/dba"
	"github.com/annuums/solanum"
	"github.com/gin-gonic/gin"
)

// Listing User Without Conditions
func NewUserListHandler() *solanum.SolaService {
	return &solanum.SolaService{
		Uri:        "/",
		Method:     http.MethodGet,
		Handler:    listHandler,
		Middleware: nil,
	}
}

func listHandler(ctx *gin.Context) {
	//*id := ctx.Query("lastname") // shortcut for ctx.Request.URL.Query().Get("lastname")
	//*id := ctx.Param("id")

	code := http.StatusOK

	ctx.JSON(http.StatusOK, gin.H{
		"code":  code,
		"users": make([]usermodel.User, 0),
	})
}
