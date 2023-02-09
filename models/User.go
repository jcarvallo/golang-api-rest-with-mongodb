package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID         primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Name       string             `bson:"name" json:"name"`
	Email      string             `bson:"email" json:"email"`
	Phone      string             `bson:"phone" json:"phone"`
	Website    string             `bson:"website" json:"website"`
	Username   string             `bson:"username" json:"username"`
	Created_At time.Time          `json:"created_at,omitempty"`
	Updated_At time.Time          `json:"updated_at,omitempty"`
}

type Users []*User
