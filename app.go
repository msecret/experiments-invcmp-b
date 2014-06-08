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

	"github.com/msecret/invcmp-b/model"
	"github.com/msecret/invcmp-b/route"
)

func main() {
	DB_PORT := os.Getenv("DB_PORT_27017_TCP_PORT")
	DB_HOST := os.Getenv("DB_PORT_27017_TCP_ADDR")
	DB_NAME := "main"
	VERSION := os.Getenv("API_VERSION")
	VERSION = "v0"

	config := map[string]string{
		"DbName":  DB_NAME,
		"Version": VERSION,
	}
	db_conn_params := fmt.Sprintf(
		"%s:%s", DB_HOST, DB_PORT)

	m := martini.Classic()
	m.Use(model.DB(db_conn_params, DB_NAME))

	m, err := route.InitHomeRoutes(m, config)
	if err != nil {
		// TODO handle error
		panic(err)
	}

	m.Run()
}
