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
		Symbol string `json:"symbol" bson:"symbol" binding:"required"`
		Group  *Group `json:"group,omitempty" bson:",omitempty"`
		Fields bson.M `json:"fields"`
	}
	// InvestmentRepo is responsible for all actions on the database related to the
	// Investment model
	InvestmentRepo struct {
		Collection *mgo.Collection
	}

	// InvestmentRepository is an interface for something that can run basic crud
	// operations on an Investment model.
	InvestmentRepositor interface {
		CreateOne(toCreate Investment) (Investment, error)
		DeleteOne(id bson.ObjectId) error
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
	if toCreate.Group != nil {
		toCreate.Group.Create()
	}

	if toCreate.Fields == nil {
		toCreate.Fields = make(map[string]interface{})
	}

	err := r.Collection.Insert(toCreate)
	if err != nil {
		return Investment{}, err
	}

	return toCreate, nil
}

// DeleteOne will delete an investment from the repo by ID. It will return an
// error if the document wasn't found or there was some other error condition.
func (r *InvestmentRepo) DeleteOne(id bson.ObjectId) error {
	err := r.Collection.RemoveId(id)
	if err != nil {
		return err
	}
	return nil
}
