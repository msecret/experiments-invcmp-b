package model

import (
	"time"

	"labix.org/v2/mgo/bson"

	"github.com/msecret/invcmp-b/util/clock"
)

type Repositor interface {
	CreateOne(m Base) (Base, error)
}

// Base is a base model for all other models to inherit from. It has fields that
// are common to all modules.
type Base struct {
	Id        bson.ObjectId `json:"id" bson:"_id"`
	CreatedAt time.Time     `json:"created_at" bson:"create_at"`
	UpdatedAt time.Time     `json:"updated_at" bson:"updated_at"`
}

// Update will update the current resource by setting its UpdatedAt value to
// the current time.
func (r *Base) Update() {
	r.UpdatedAt = clock.Now()
}

// Create will update the current resource by setting its UpdatedAt and CreatedAt
// value to the current time.
func (r *Base) Create() {
	id := bson.NewObjectId()
	r.Id = id
	r.CreatedAt = clock.Now()
	r.UpdatedAt = clock.Now()
}
