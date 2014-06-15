package model

import (
	"errors"
	"fmt"
	log "github.com/cihub/seelog"
	"reflect"

	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
)

const ERR_NOT_FOUND = "not_found"

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

// GetOne will search for an investment and return it if found, return a not
// found error if not found, or return other error if one occurred.
func (r *InvestmentRepo) FindOne(id string) (Investment, error) {
	toReturnInvestment := Investment{}
	bsonId := bson.ObjectIdHex(id)
	err := r.Collection.FindId(bsonId).One(&toReturnInvestment)
	if err != nil {
		if err.Error() == mgo.ErrNotFound.Error() {
			return toReturnInvestment, errors.New(ERR_NOT_FOUND)
		} else {
			return toReturnInvestment, err
		}
	}
	return toReturnInvestment, nil
}

// GetOneBySymbol will search for an investment by its symbol and will return
// it if it exists. If the "not found" error occurs, will return a NOT_FOUND
// error, or just the error otherwise.
// Because symbol is unique, this is not part of a list of investments but just
// a single investment.
func (r *InvestmentRepo) FindOneBySymbol(symbol string) (Investment, error) {
	toReturnInvestment := Investment{}
	err := r.Collection.Find(bson.M{"symbol": symbol}).One(&toReturnInvestment)
	if err != nil {
		if err.Error() == mgo.ErrNotFound.Error() {
			return toReturnInvestment, errors.New(ERR_NOT_FOUND)
		} else {
			return toReturnInvestment, err
		}
	}
	return toReturnInvestment, nil
}

// Get multiple will take a map of params to search on and return a list of
// investments that fulfill the params.
func (r *InvestmentRepo) FindMultiple(params map[string]interface{}) (
	Investments, error) {
	toReturnInvestments := Investments{}
	bsonParams := mappingToFieldMatches(params)

	err := r.Collection.Find(bsonParams).All(&toReturnInvestments)
	if err != nil {
		return toReturnInvestments, err
	}

	return toReturnInvestments, nil
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
func (r *InvestmentRepo) DeleteOne(id string) error {
	bsonId := bson.ObjectIdHex(id)
	err := r.Collection.RemoveId(bsonId)
	if err != nil {
		if err.Error() == mgo.ErrNotFound.Error() {
			return errors.New(ERR_NOT_FOUND)
		} else {
			log.Error(err.Error())
			return err
		}
	}
	return nil
}

func mappingToFieldMatches(mapping map[string]interface{}) bson.M {
	/*
	   //start
	   {
	     fields: {
	       cap: poop
	     }
	     group: {
	       name: poop
	     }
	   }
	   //end
	   {
	     fields.cap: poop,
	     group.name: poop
	   }
	*/
	fields := bson.M{}
	for key, value := range mapping {
		fieldPrefix := key
		if reflect.TypeOf(value).Kind() == reflect.Map {
			for innerKey, innerValue := range value.(map[string]interface{}) {
				fieldKey := fmt.Sprintf("%s.%s", fieldPrefix, innerKey)
				fields[fieldKey] = innerValue
			}
		}
	}

	return fields
}
