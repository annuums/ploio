package users

import (
	"log"
	"net/http"

	"github.com/annuums/solanum"
	"github.com/gin-gonic/gin"
)

func NewUserIndexHandler() *solanum.Service {
	return &solanum.Service{
		Uri:        "/",
		Method:     http.MethodGet,
		Handler:    indexHandler,
		Middleware: indexMiddleware,
	}
}

func indexHandler(ctx *gin.Context) {
	ctx.JSON(200, "Hello, World! From User Index Handler.")
}

func indexMiddleware(ctx *gin.Context) {
	log.Println("Hello Index Middleware")
	ctx.Next()
}

func NewUserHelloHandler() *solanum.Service {
	return &solanum.Service{
		Uri:        "/hello",
		Method:     http.MethodGet,
		Handler:    indexHandler2,
		Middleware: indexMiddleware2,
	}
}

func indexHandler2(ctx *gin.Context) {
	ctx.JSON(200, "Hello, World! From User Hello Handler.")
}

func indexMiddleware2(ctx *gin.Context) {
	log.Println("Hello Index Middleware")
	ctx.Next()
}
