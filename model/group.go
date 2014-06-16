package model

import (
	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
)

type (
	Groups []Group
	// Group is a struct representing a group schema. The json schema is defined
	// as follows:
	// {
	//   id: int,
	//   name: string,
	// }
	Group struct {
		Base `bson:",inline"`
		Name string `json:"name" bson"name"`
	}
	// GroupRepo is responsible for all actions on the database related to the
	// Group model.
	GroupRepo struct {
		Collection *mgo.Collection
	}
)

// NewGroupRepo creates a new group repository, sets the correct
// database collection and ensures any indexes on the collection. Returns
// the new repository.
func NewGroupRepo(sesh *mgo.Database) GroupRepo {
	idx := mgo.Index{
		Key:        []string{"name"},
		Unique:     true,
		DropDups:   true,
		Background: true,
		Sparse:     true,
	}
	r := GroupRepo{
		Collection: sesh.C("groups"),
	}
	r.Collection.EnsureIndex(idx)

	return r
}

// GetOne returns one group from the database that matches based on the name
// string passed in.
// Returns the group or an error if one occurred.
func (r *GroupRepo) GetOne(name string) (Group, error) {
	result := Group{}
	err := r.Collection.Find(bson.M{"name": name}).One(&result)
	if err != nil {
		return result, err
	}
	return result, nil
}

// Create will insert a new document of schema group to the database.
// Returns the group struct that was inserted or an error if one was
// encountered.
func (r *GroupRepo) Create(toCreate Group) (Group, error) {
	toCreate.Create()
	err := r.Collection.Insert(toCreate)
	if err != nil {
		panic(err) // TODO handle_error
	}

	return toCreate, nil
}
