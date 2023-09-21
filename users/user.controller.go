package users

import "github.com/annuums/solanum"

var helloWorldController solanum.Controller

func NewUserController() (solanum.Controller, error) {
	if helloWorldController == nil {
		helloWorldController, _ = solanum.NewController()
		addHandlers()
	}

	return helloWorldController, nil
}

func addHandlers() {
	indexHandler := NewUserIndexHandler()
	helloHandler := NewUserHelloHandler()

	helloWorldController.AddHandler(indexHandler, helloHandler)
}
