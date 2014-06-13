package model

import (
	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
)

type (
	// Investments is a list of Investment(s).
	Investments []Investment

	// Investment models a single entity on the stock market searchable by
	// its symbol. It may have a group associated with it. It's fields is a
	// investment schema:
	//   /public/schema/investment.json
	Investment struct {
		Base
		Id     bson.ObjectId `json:"id" bson:"_id"`
		Symbol string        `json:"symbol" bson:"symbol"`
		Group  Group         `json:"group" bson:"group"`
		Fields bson.M        `json:"fields" bson:"fields"`
	}
	// InvestmentRepo is responsible for all actions on the database related to the
	// Investment model
	InvestmentRepo struct {
		Collection *mgo.Collection
	}
)

// NewInvestmentRepo acreates a new investment repository, sets the correct
// database collection and ensures any indexes on the collection. Returns
// the new repository.
func NewInvestmentRepo(sesh *mgo.Database) InvestmentRepo {
	idx := mgo.Index{
		Key:        []string{"symbol"},
		Unique:     true,
		DropDups:   true,
		Background: true,
		Sparse:     true,
	}
	r := InvestmentRepo{
		Collection: sesh.C("investments"),
	}
	r.Collection.EnsureIndex(idx)

	return r
}

// CreateOne will attempt to update the base time fields and then insert one
// Investment into the database by taking a Investment to be created. Returns
// the created investment.
func (r *InvestmentRepo) CreateOne(toCreate Investment) (Investment, error) {
	toCreate.Create()
	err := r.Collection.Insert(toCreate)
	if err != nil {
		return Investment{}, err
	}

	return toCreate, nil
}
