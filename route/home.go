package route

import (
	"github.com/go-martini/martini"
	"labix.org/v2/mgo"
)

func InitHomeRoutes(api martini.Router, config map[string]string) (
	martini.Router, error) {
	api.Get("/st", func(db *mgo.Database) string {
		return "st: " + config["DbName"]
	})

	return api, nil
}
