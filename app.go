package main

import (
	"fmt"
	"os"

	"github.com/go-martini/martini"
)

type StatusMessage struct {
	DbConf  string
	AppConf string
}

func main() {
	DB_PORT := os.Getenv("DB_PORT_27017_TCP_PORT")
	DB_HOST := os.Getenv("DB_PORT_27017_TCP_ADDR")
	DB_NAME := "main"

	db_conn_params := fmt.Sprintf(
		"dbname=%s "+
			"host=%s "+
			"port=%s ", DB_NAME, DB_HOST, DB_PORT)

	m := martini.Classic()
	m.Get("/st", func() string {
		return "st: " + db_conn_params
	})

	m.Run()
}
