package main

import (
	"github.com/annuums/ploio/common/config"
	"github.com/annuums/ploio/users"
	"github.com/annuums/solanum"
	"github.com/gin-gonic/gin"
)

func main() {
	server := *solanum.NewSolanum(5050)
	server.GetGinEngine().Use(gin.Logger(), gin.Recovery())

	var userModule solanum.Module
	userUri := "/users"
	userModule, _ = users.NewUserModule(
		server.GetGinEngine().Group(userUri),
		userUri,
	)

	cfgReader := config.NewConfigReader()
	cfgReader.Initialize()

	server.AddModule(&userModule)

	server.Run()
}
