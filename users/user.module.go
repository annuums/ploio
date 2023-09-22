package users

import (
	"log"

	"github.com/annuums/solanum"
	"github.com/gin-gonic/gin"
)

var userModule *solanum.SolaModule

func NewUserModule(router *gin.RouterGroup, uri string) (*solanum.SolaModule, error) {
	if userModule == nil {
		userModule, _ = solanum.NewModule(router, uri)
		attachControllers()
	}

	return userModule, nil
}

func attachControllers() {
	//* Attatching Controller Directly
	var (
		ctr solanum.Controller
		err error
	)

	ctr, err = NewUserController()

	if err != nil {
		log.Fatal(err)
	}

	userModule.SetControllers(&ctr)
}
