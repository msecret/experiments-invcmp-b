package main

import (
	"fmt"
	"os"

	"github.com/go-martini/martini"
	"labix.org/v2/mgo"
)

type StatusMessage struct {
	DbConf  string
	AppConf string
}

type User struct {
	Name  string
	Email string
}

func main() {
	DB_PORT := os.Getenv("DB_PORT_27017_TCP_PORT")
	DB_HOST := os.Getenv("DB_PORT_27017_TCP_ADDR")
	DB_NAME := "main"

	db_conn_params := fmt.Sprintf(
		"dbname=%s "+
			"host=%s "+
			"port=%s ", DB_NAME, DB_HOST, DB_PORT)

	session, err := mgo.Dial(fmt.Sprintf("%s:%s", DB_HOST, DB_PORT))
	if err != nil {
		panic(err)
	}
	defer session.Close()

	c := session.DB(DB_NAME).C("user")
	err = c.Insert(&User{"marco", "marco@minted.com"})
	if err != nil {
		panic(err)
	}
	fmt.Println("User created")

	m := martini.Classic()
	m.Get("/st", func() string {
		return "st: " + db_conn_params
	})

	m.Run()
}
