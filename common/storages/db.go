package storages

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/annuums/ploio/common/config"
	_ "github.com/go-sql-driver/mysql"
)

type Storage struct {
}

type Database struct {
}

func NewDatabase() *Database {
	return &Database{}
}

func (db *Database) Connect() *sql.DB {
	cfgReader := config.NewConfigReader()

	env := cfgReader.Get("GO_ENV")
	url := ""

	switch env {
	case "dev":
		url = fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?parseTime=true",
			cfgReader.Get("DB_USER_DEV"),
			cfgReader.Get("DB_PASSWORD_DEV"),
			cfgReader.Get("DB_HOSTNAME_DEV"),
			cfgReader.Get("DB_PORT_DEV"),
			cfgReader.Get("DB_SCHEME_DEV"),
		)
	case "production":
		url = fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?parseTime=true",
			cfgReader.Get("DB_USER"),
			cfgReader.Get("DB_PASSWORD"),
			cfgReader.Get("DB_HOSTNAME"),
			cfgReader.Get("DB_PORT"),
			cfgReader.Get("DB_SCHEME"),
		)
	}

	fmt.Printf("Env: %v, Connection Created: %v\n", env, url)

	connection, err := sql.Open("mysql", url)

	if err != nil {
		log.Fatal(err)
	}

	return connection
}

func (db *Database) ConnectTx() (*sql.Tx, error) {
	_db := db.Connect()

	return _db.Begin()
}
