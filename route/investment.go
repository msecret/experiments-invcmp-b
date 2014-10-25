package route

import (
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"strings"

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

	api.Get("/investment/:id", GetOne)
	api.Get("/investment", GetOneBySymbol)
	api.Get("/investments", GetMultiple)
	api.Post("/investments", binding.Bind(model.InvestmentRequest{}), CreateOne)
	api.Delete("/investment/:id", DeleteOne)

	return api, nil
}

// GetOne will attempt to find a resource by id in the database and return
// it if exists, 404 error if it doesn't.
//
// curl -i -H "Accept: application/json" http://localhost:49182/api/v0/investment/{id}
func GetOne(params martini.Params, sesh *mgo.Database, r render.Render) {
	investmentRepo.Collection = sesh.C("investments")
	investment, err := investmentRepo.FindOne(params["id"])
	if err != nil {
		if err.Error() == model.ERR_NOT_FOUND {
			r.JSON(ResponseNotFound())
		} else {
			log.Error(err.Error())
			r.JSON(ResponseInternalServerError(err))
		}
		return
	}

	r.JSON(ResponseSuccess(investment, "investment"))
}

// GetOneBySymbol will attempt to find a resource by symbol name in the database
// and return it if exists, 404 error if it doesn't.
//
// curl -i -H "Accept: application/json"
//   http://localhost:49182/api/v0/investment?symbol={symbol}
func GetOneBySymbol(req *http.Request, sesh *mgo.Database, r render.Render) {
	query := req.URL.Query()
	symbolParam := query["symbol"]
	if len(symbolParam) < 1 {
		r.JSON(ResponseBadRequest(errors.New("Request missing required name field")))
		return
	}

	symbol := symbolParam[0]
	investmentRepo.Collection = sesh.C("investments")
	investment, err := investmentRepo.FindOneBySymbol(symbol)
	if err != nil {
		if err.Error() == model.ERR_NOT_FOUND {
			r.JSON(ResponseNotFound())
		} else {
			log.Error(err.Error())
			r.JSON(ResponseInternalServerError(err))
		}
		return
	}

	r.JSON(ResponseSuccess(investment, "investment"))
}

// GetMultiple will get multiple resources. It will look for query params and
// get by query params if valid ones exist. Otherwise it will get all
// investments.
// curl -i -H "Accept: application/json"
//   http://localhost:49182/api/v0/investments
//   http://localhost:49182/api/v0/investments?cap=20&price=15
//   http://localhost:49182/api/v0/investments?group-name=test
func GetMultiple(req *http.Request, sesh *mgo.Database, r render.Render) {
	var (
		err    error
		params map[string]interface{}
	)
	query := req.URL.Query()
	if len(query) > 0 {
		params, err = handleParams(query)
	}

	investmentRepo.Collection = sesh.C("investments")
	investment, err := investmentRepo.FindMultiple(params)
	if err != nil {
		r.JSON(ResponseInternalServerError(err))
	}

	r.JSON(ResponseSuccess(investment, "investments"))
}

// CreateOne is the handler for when a new resource is being created with
// a POST request. It will take an Investment model and return the investment
// model after it was inserted as JSON.
func CreateOne(investment model.InvestmentRequest, params martini.Params,
	sesh *mgo.Database, r render.Render) {
	investmentRepo.Collection = sesh.C("investments")
	createdInvestment, err := investmentRepo.CreateOne(investment.Investment)
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
//
// curl -i -H "Accept: application/json" -X DELETE http://localhost:49182/api/v0/investment/{id}
func DeleteOne(params martini.Params, sesh *mgo.Database, r render.Render) {
	investmentRepo.Collection = sesh.C("investments")
	err := investmentRepo.DeleteOne(params["id"])
	if err != nil {
		// If type of error is mg.ErrNotFound, return a 404.
		if err.Error() == model.ERR_NOT_FOUND {
			r.JSON(ResponseNotFound())
		} else {
			log.Error(err.Error())
			r.JSON(ResponseInternalServerError(err))
		}
		return
	}

	r.JSON(ResponseSuccessNoData())
	return
}

func handleParams(query url.Values) (map[string]interface{}, error) {
	params, errs := TransformQueryToMapping(query)
	toReturnErrs := []string{}
	if len(errs) > 0 {
		for _, err := range errs {
			toReturnErrs = append(toReturnErrs, err.Error())
		}
		errorString := strings.Join(toReturnErrs, ", ")
		err := errors.New(
			fmt.Sprintf("Request has invalid fields: %s", errorString))
		return params, err
	}
	return params, nil
}
