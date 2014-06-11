package model

import (
	"time"

	"github.com/msecret/invcmp-b/util/clock"
)

type Base struct {
	CreatedAt time.Time `json:"created_at" bson:"create_at"`
	UpdatedAt time.Time `json:"updated_at" bson:"updated_at"`
}

func (r *Base) Update() {
	r.UpdatedAt = clock.Now()
}

func (r *Base) Create() {
	r.CreatedAt = clock.Now()
	r.UpdatedAt = clock.Now()
}
