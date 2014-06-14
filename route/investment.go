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

	return api, nil
}

func CreateOne(investment model.Investment, params martini.Params,
	sesh *mgo.Database, r render.Render) {
	log.Info("in CreateOne")
	investmentRepo.Collection = sesh.C("investments")
	createdInvestment, err := investmentRepo.CreateOne(investment)
	if err != nil {
		r.JSON(500, map[string]interface{}{"status": "failure",
			"error_message": err.Error()})
		return
	}

	r.JSON(200, map[string]interface{}{"status": "success",
		"data": map[string]interface{}{
			"investment": createdInvestment},
	})

	return
}
