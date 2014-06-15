package route

import (
	log "github.com/cihub/seelog"
	"github.com/codegangsta/martini-contrib/render"
	"github.com/go-martini/martini"
	"github.com/martini-contrib/binding"
	"labix.org/v2/mgo"

	"github.com/msecret/invcmp-b/model"
)

var investmentRepo model.InvestmentRepo

// InitInvestmentRoutes Initializes all routes for the investment schema.
// Takes a router to add routes to and config  for the api.
// Returns the same router with new routes added.
func InitInvestmentRoutes(api martini.Router, db *mgo.Database) (
	martini.Router, error) {

	investmentRepo = model.NewInvestmentRepo(db)

	api.Post("/investments", binding.Bind(model.Investment{}), CreateOne)
	api.Delete("/investments/:id", DeleteOne)

	return api, nil
}

// CreateOne is the handler for when a new resource is being created with
// a POST request. It will take an Investment model and return the investment
// model after it was inserted as JSON.
func CreateOne(investment model.Investment, params martini.Params,
	sesh *mgo.Database, r render.Render) {
	investmentRepo.Collection = sesh.C("investments")
	createdInvestment, err := investmentRepo.CreateOne(investment)
	if err != nil {
		r.JSON(ResponseInternalServerError(err))
		return
	}

	r.JSON(ResponseSuccess(createdInvestment, "investment"))

	return
}

// DeleteOne will take an id from the url params and request that the resource
// be deleted from the database. Returns a 404 response if the resource was not
// found, a 500 for other errors and a success response with no data on success.
func DeleteOne(params martini.Params, sesh *mgo.Database, r render.Render) {
	investmentRepo.Collection = sesh.C("investments")
	err := investmentRepo.DeleteOne(params["id"])
	if err != nil {
		// If type of error is mg.ErrNotFound, return a 404.
		if err.Error() == model.ERR_NOT_FOUND {
			r.JSON(ResponseNotFound())
			return
		} else {
			log.Error(err.Error())
			r.JSON(ResponseInternalServerError(err))
			return
		}
	}

	r.JSON(ResponseSuccessNoData())
	return
}
