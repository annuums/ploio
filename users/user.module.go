package users

import (
	"github.com/annuums/solanum"
	"github.com/gin-gonic/gin"
)

var userModule solanum.Module

func NewUserModule(router *gin.RouterGroup, uri string) (solanum.Module, error) {
	if userModule == nil {
		userModule, _ = solanum.NewModule(router, uri)
		attachControllers()
	}

	return userModule, nil
}

func attachControllers() {
	//* Attatching Controller Directly
	ctr, _ := NewUserController()
	// ctr2, _ := NewAnotherController()
	//	...

	userModule.SetControllers(ctr)
}
