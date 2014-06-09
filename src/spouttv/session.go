package spouttv

import (
  // "labix.org/v2/mgo"
  "labix.org/v2/mgo/bson"
  "time"
)

type Session struct {
  Id          bson.ObjectId `json:"id"         bson:"_id,omitempty"`
  Token       string        `json:"token"      bson:"token"`
  CreatedAt   time.Time     `json:"created_at" bson:"created_at"`
}

