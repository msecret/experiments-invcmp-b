package model

import (
	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
)

type (
	Groups []Group
	Group  struct {
		Base
		Id   bson.ObjectId `json:"id" bson:"_id"`
		Name string        `json:"name" bson"name"`
	}
	GroupRepo struct {
		Collection *mgo.Collection
	}
)

func (r *GroupRepo) GetOne(name string) (Group, error) {
	result := Group{}
	err := r.Collection.Find(bson.M{"name": name}).One(&result)
	if err != nil {
		return result, err
	}
	return result, nil
}

func (r *GroupRepo) Create(toCreate Group) (Group, error) {
	toCreate.Create()
	err := r.Collection.Insert(toCreate)
	if err != nil {
		panic(err) // TODO handle_error
	}

	return toCreate, nil
}
