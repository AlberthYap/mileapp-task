package models

import (
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
)


type User struct {
    ID        bson.ObjectID `bson:"_id,omitempty" json:"id"`
    Email     string             `bson:"email" json:"email" binding:"required,email"`
    Password  string             `bson:"password" json:"-"`
    Name      string             `bson:"name" json:"name" binding:"required,min=2,max=100"`
    CreatedAt time.Time          `bson:"created_at" json:"created_at"`
    UpdatedAt time.Time          `bson:"updated_at" json:"updated_at"`
}
