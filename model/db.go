package model

import (
	"github.com/go-martini/martini"
	"labix.org/v2/mgo"
)

// DB will return a function to create a new db session by taking in a string
// of database params and a string of the db names and returning a handler for
// setting up the db session context.
func DB(dbParams string, dbName string) martini.Handler {
	session, err := mgo.Dial(dbParams)
	if err != nil {
		panic(err)
	}

	return func(c martini.Context) {
		s := session.Clone()
		c.Map(s.DB(dbName))
		defer s.Close()
		c.Next()
	}
}
