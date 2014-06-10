package model

import (
	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
	"time"
)

type (
	Groups []Group
	Group  struct {
		Id   bson.ObjectId `json:"id" bson:"_id"`
		Name string        `json:"name" bson"name"`
		// Make this defined on every model somehow
		CreatedAt time.Time `json:"created_at" bson:"create_at"`
		UpdatedAt time.Time `json:"updated_at" bson:"updated_at"`
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
	toCreate.CreatedAt = time.Now()
	toCreate.UpdatedAt = time.Now()
	r.Collection.UpsertId(toCreate.Id, toCreate)

	return toCreate, nil
}
