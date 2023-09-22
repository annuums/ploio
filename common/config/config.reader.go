package config

import (
	"log"
	"os"
	"path/filepath"
	"sync"

	"github.com/joho/godotenv"
)

var once sync.Once

type kv struct {
	key   string
	value string
}

type ConfigReader struct {
	ch   chan kv
	data map[string]string
}

var r *ConfigReader

func NewConfigReader() *ConfigReader {
	if r == nil {
		once.Do(func() {
			r = &ConfigReader{
				ch:   make(chan kv),
				data: make(map[string]string),
			}
		})
	}

	return r
}

func (r *ConfigReader) Initialize() {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}
	environmentPath := filepath.Join(dir, ".env")
	godotenv.Load(environmentPath)

	env := os.Getenv("GO_ENV")

	if env == "" {
		environmentPath = "/Users/eugene/Playground/gospace/annuums/ploio/.env"
		godotenv.Load(environmentPath)
	}

	log.Println("Initializing Config Reader...")

	r.Put("GO_ENV", os.Getenv("GO_ENV"))
	//* Dev Env
	{
		r.Put("DB_HOSTNAME_DEV", os.Getenv("DB_HOSTNAME_DEV"))
		r.Put("DB_PORT_DEV", os.Getenv("DB_PORT_DEV"))
		r.Put("DB_USER_DEV", os.Getenv("DB_USER_DEV"))
		r.Put("DB_PASSWORD_DEV", os.Getenv("DB_PASSWORD_DEV"))
		r.Put("DB_SCHEME_DEV", os.Getenv("DB_SCHEME_DEV"))
	}
	//* Prod Env
	{
		r.Put("DB_HOSTNAME", os.Getenv("DB_HOSTNAME"))
		r.Put("DB_PORT", os.Getenv("DB_PORT"))
		r.Put("DB_USER", os.Getenv("DB_USER"))
		r.Put("DB_PASSWORD", os.Getenv("DB_PASSWORD"))
		r.Put("DB_SCHEME", os.Getenv("DB_SCHEME"))
	}

	close(r.ch)
}

func (r *ConfigReader) Get(key string) string {
	return r.data[key]
}

func (r *ConfigReader) Put(key string, value string) {
	go func() {
		_kv, ok := <-r.ch

		if !ok {
			log.Println("Channel already closed")
			return
		}

		r.data[_kv.key] = _kv.value
	}()

	r.ch <- kv{
		key:   key,
		value: value,
	}
}
