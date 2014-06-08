package model

import (
	"github.com/go-martini/martini"
	"labix.org/v2/mgo"
)

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
