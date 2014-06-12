package model

import (
	"time"

	"github.com/msecret/invcmp-b/util/clock"
)

// Base is a base model for all other models to inherit from. It has fields that
// are common to all modules.
type Base struct {
	CreatedAt time.Time `json:"created_at" bson:"create_at"`
	UpdatedAt time.Time `json:"updated_at" bson:"updated_at"`
}

// Update will update the current resource by setting its UpdatedAt value to
// the current time.
func (r *Base) Update() {
	r.UpdatedAt = clock.Now()
}

// Create will update the current resource by setting its UpdatedAt and CreatedAt
// value to the current time.
func (r *Base) Create() {
	r.CreatedAt = clock.Now()
	r.UpdatedAt = clock.Now()
}
