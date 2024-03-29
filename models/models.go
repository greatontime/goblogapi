package models

import "go.mongodb.org/mongo-driver/bson/primitive"

//Blog struct
type Blog struct {
	ID     primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Name string `json:"name,omitempty"`
	Status bool `json:"status,omitempty"`
}