package route

import (
	"github.com/go-martini/martini"
	"labix.org/v2/mgo"
)

func InitHomeRoutes(m *martini.ClassicMartini, config map[string]string) (
	*martini.ClassicMartini, error) {
	m.Get("/"+config["Version"]+"/st", func(db *mgo.Database) string {
		return "st: " + config["DbName"]
	})

	return m, nil
}
