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
	"github.com/martini-contrib/strip"

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
	apiPrefix := fmt.Sprintf("/api/%s", VERSION)

	m := martini.Classic()
	m.Use(model.DB(db_conn_params, DB_NAME))

	api := martini.NewRouter()
	api, err := route.InitHomeRoutes(api, config)
	if err != nil {
		// TODO handle error
		panic(err)
	}

	// Prefix all api requests with api/{version id}/
	m.Get(apiPrefix+"/**", strip.Prefix(apiPrefix), api.Handle)
	m.Run()
}
