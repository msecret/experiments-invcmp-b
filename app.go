/*
Copyright (c) 2013
All Rights Reserved
Licensed MIT

https://github.com/msecret/experiments-invcmp-b
stockarator v0.0.1
*/

package main

import (
	"fmt"
	"os"

	"github.com/go-martini/martini"
	"labix.org/v2/mgo"

	"github.com/msecret/invcmp-b/model"
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
		"%s:%s", DB_HOST, DB_PORT)

	m := martini.Classic()
	m.Use(model.DB(db_conn_params, DB_NAME))

	m.Get("/st", func(db *mgo.Database) string {
		return "st: " + db_conn_params
	})

	m.Run()
}
