package users

import (
	"context"
	"log"
	"net/http"
	"strconv"

	"github.com/annuums/ploio/common/storages"
	usermodel "github.com/annuums/ploio/users/dba"
	"github.com/annuums/solanum"
	"github.com/gin-gonic/gin"
)

// Listing User Without Conditions
func NewUserHandlers() []*solanum.SolaService {
	return []*solanum.SolaService{
		//* List Users
		{
			Uri:        "/",
			Method:     http.MethodGet,
			Handler:    listHandler,
			Middleware: nil,
		},
		//* List Single User By Id
		{
			Uri:     "/:id",
			Method:  http.MethodGet,
			Handler: getOneHandler,
			//* Need Authorization
			Middleware: nil,
		},
		//* Create User
		{
			Uri:        "/",
			Method:     http.MethodPost,
			Handler:    createUserHandler,
			Middleware: nil,
		},
		{
			Uri:     "/",
			Method:  http.MethodPatch,
			Handler: updateUserHandler,
			//* Need Authorization
			Middleware: nil,
		},
	}
}

func listHandler(ctx *gin.Context) {
	//*id := ctx.Query("lastname") // shortcut for ctx.Request.URL.Query().Get("lastname")
	//*id := ctx.Param("id")

	code := http.StatusOK

	/* Prepare Database Session */
	db := storages.NewDatabase()
	conn := db.Connect()
	tx, err := conn.Begin()

	if err != nil {
		log.Fatal("[listHandler]: Fail to start transaction")
	}

	defer tx.Rollback()

	query := usermodel.New(conn)
	qtx := query.WithTx(tx)

	/* End of Prepare Database Session. Need to be separated */

	result, err := qtx.ListUsers(context.Background())

	if err != nil {
		log.Fatal("[listHandler]: Fail to get users: ", err)
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code":  code,
		"users": result,
	})
}

func getOneHandler(ctx *gin.Context) {
	//*id := ctx.Query("lastname") // shortcut for ctx.Request.URL.Query().Get("lastname")
	//*id := ctx.Param("id")
	id := ctx.Param("id")
	_id, err := strconv.Atoi(id)
	code := http.StatusOK
	if err != nil {
		code = http.StatusBadRequest
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code":    code,
			"message": err.Error(),
		})
		return
	}

	/* Prepare Database Session */
	db := storages.NewDatabase()
	conn := db.Connect()
	tx, err := conn.Begin()

	if err != nil {
		log.Fatal("[listHandler]: Fail to start transaction")
	}

	defer tx.Rollback()

	query := usermodel.New(conn)
	qtx := query.WithTx(tx)

	/* End of Prepare Database Session. Need to be separated */

	result, err := qtx.GetUser(context.Background(), int64(_id))

	if err != nil {
		log.Fatal("[listHandler]: Fail to get users: ", err)
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code": code,
		"user": result,
	})
}

func createUserHandler(ctx *gin.Context) {
	//*id := ctx.Query("lastname") // shortcut for ctx.Request.URL.Query().Get("lastname")
	//*id := ctx.Param("id")
	code := http.StatusCreated

	/* Prepare Database Session */
	db := storages.NewDatabase()
	conn := db.Connect()
	tx, err := conn.Begin()

	if err != nil {
		log.Fatal("[createUserHandler]: Fail to start transaction")
	}

	defer tx.Rollback()

	query := usermodel.New(conn)
	qtx := query.WithTx(tx)

	/* End of Prepare Database Session. Need to be separated */

	var user usermodel.CreateUserParams
	err = ctx.BindJSON(&user)

	log.Println(user)

	if err != nil {
		log.Printf("[createUserHandler]: Fail to create user:: %v\n", err)
		code = http.StatusBadRequest

		ctx.JSON(code, gin.H{
			"code":    code,
			"message": err.Error(),
		})
		return
	}

	result, err := qtx.CreateUser(context.Background(), user)

	if err != nil {
		log.Printf("[createUserHandler]: Fail to create user:: %v\n", err)
		code = http.StatusInternalServerError

		ctx.JSON(code, gin.H{
			"code":    code,
			"message": err.Error(),
		})
		return
	}

	log.Println(result.LastInsertId())
	ctx.JSON(code, gin.H{
		"code": code,
		"user": result,
	})

	err = tx.Commit()
	if err != nil {
		log.Printf("[createUserHandler]: Fail to create user:: %v\n", err)
		code = http.StatusBadRequest

		ctx.JSON(code, gin.H{
			"code":    code,
			"message": err.Error(),
		})
		return
	}
}

func updateUserHandler(ctx *gin.Context) {
	//*id := ctx.Query("lastname") // shortcut for ctx.Request.URL.Query().Get("lastname")
	//*id := ctx.Param("id")
	code := http.StatusCreated

	/* Prepare Database Session */
	db := storages.NewDatabase()
	conn := db.Connect()
	tx, err := conn.Begin()

	if err != nil {
		log.Fatal("[updateUserHandler]: Fail to start transaction")
	}

	defer tx.Rollback()

	// query := usermodel.New(conn)
	// qtx := query.WithTx(tx)

	/* End of Prepare Database Session. Need to be separated */

	ctx.JSON(code, gin.H{
		"code":       code,
		"need_to_be": "implemented",
	})

	err = tx.Commit()
	if err != nil {
		log.Printf("[updateUserHandler]: Fail to create user:: %v\n", err)
		code = http.StatusBadRequest

		ctx.JSON(code, gin.H{
			"code":    code,
			"message": err.Error(),
		})
		return
	}
}
