package route

import (
	"github.com/go-martini/martini"
	"labix.org/v2/mgo"

	"github.com/msecret/invcmp-b/model"
)

// InitGroupRoutes Initializes all routes for the group schema.
// Takes a router to add routes to and config  for the api.
// Returns the same router with new routes added.
func InitGroupRoutes(api martini.Router, config map[string]string) (
	martini.Router, error) {

	repo := model.GroupRepo{}

	api.Get("/group/:name",
		func(params martini.Params, sesh *mgo.Database) string {
			repo.Collection = sesh.C("group")
			group, err := repo.GetOne(params["name"])
			if err != nil {
				panic(err) // TODO handle err
			}

			return group.Name
		})

	return api, nil
}
