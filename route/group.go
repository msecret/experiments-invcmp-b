package route

import (
	"github.com/codegangsta/martini-contrib/render"
	"github.com/go-martini/martini"
	"labix.org/v2/mgo"

	"github.com/msecret/invcmp-b/model"
)

var repo model.GroupRepo

// InitGroupRoutes Initializes all routes for the group schema.
// Takes a router to add routes to and config  for the api.
// Returns the same router with new routes added.
func InitGroupRoutes(api martini.Router, db *mgo.Database) (
	martini.Router, error) {

	repo = model.NewGroupRepo(db)

	api.Get("/group/:name", GetOneByName)

	return api, nil
}

// GetOneByName takes martini params and a database session and returns one
// Group model. Should search for the one by name which should be present in
// params.
func GetOneByName(params martini.Params,
	sesh *mgo.Database, r render.Render) {
	repo.Collection = sesh.C("groups")
	group, err := repo.GetOne(params["name"])
	if err != nil {
		if err.Error() == "not found" {
			r.JSON(404, map[string]interface{}{"status": "failure",
				"error_message": err.Error()})
			return
		}
		r.JSON(500, map[string]interface{}{"status": "failure",
			"error_message": err.Error()})
	}

	r.JSON(200, map[string]interface{}{"status": "success", "data": group})
}
