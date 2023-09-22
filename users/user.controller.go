package users

import (
	userhandlers "github.com/annuums/ploio/users/handlers"
	"github.com/annuums/solanum"
)

var helloWorldController *solanum.SolaController

func NewUserController() (*solanum.SolaController, error) {
	if helloWorldController == nil {
		helloWorldController, _ = solanum.NewController()
		addHandlers()
	}

	return helloWorldController, nil
}

func addHandlers() {
	getOneHandler := userhandlers.NewUserGetOneHandler()
	getListHandler := userhandlers.NewUserListHandler()

	helloWorldController.AddHandler(getOneHandler, getListHandler)
}

//* Middleware
