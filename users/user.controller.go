package users

import (
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
	handlers := NewUserHandlers()

	helloWorldController.AddHandler(handlers...)
}

//* Middleware
