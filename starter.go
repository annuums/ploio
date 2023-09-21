package main

import (
	"github.com/annuums/ploio/users"
	"github.com/annuums/solanum"
)

func main() {
	server := *solanum.NewSolanum(5050)

	userUri := "/users"
	userModule, _ := users.NewUserModule(
		server.GetGinEngine().Group(userUri),
		userUri,
	)

	server.AddModule(&userModule)

	// server.GET("/posts", func(ctx *gin.Context) {
	// 	ctx.JSON(
	// 		http.StatusOK,
	// 		videoController.FindAll(),
	// 	)
	// })

	// server.POST("/posts", func(ctx *gin.Context) {
	// 	ctx.JSON(
	// 		http.StatusCreated,
	// 		videoController.Save(ctx),
	// 	)
	// })

	server.Run()
}
