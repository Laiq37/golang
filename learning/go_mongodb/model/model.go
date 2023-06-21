package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Netflix struct {
	Id        primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	MovieName string             `json:"moviename,omitempty"`
	IsWatched bool               `json:"iswatched,omitempty"`
}
