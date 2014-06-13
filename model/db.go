package model

import (
	"github.com/go-martini/martini"
	"labix.org/v2/mgo"
)

// CreateSesh will create a new database session with the database params and
// database name passed in as strings. Returns a poiner to an mgo.Session or an
// error if one occurred.
func CreateSesh(dbParams string, dbName string) (*mgo.Session, error) {
	session, err := mgo.Dial(dbParams)
	if err != nil {
		return &mgo.Session{}, err
	}

	return session, nil
}

// DB will return a function to create a new db session by taking in a string
// of database params and a string of the db names and returning a handler for
// setting up the db session context.
func DB(dbParams string, dbName string) martini.Handler {
	session, err := CreateSesh(dbParams, dbName)
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
