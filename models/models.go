
package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type List struct {
	ID  primitive.ObjectID  `json:"_id,omitempty" bson:"_id,omitempty"`
	Name string  `json:"name,omitempty" bson:"name,omitempty"`
	DateCreated string  `json:"date_created,omitempty" bson:"date_created,omitempty"`
	Tasks []Task `json:"tasks,omitempty" bson:"tasks,omitempty"`
}

type Task struct {
	ID  primitive.ObjectID  `json:"_id,omitempty" bson:"_id,omitempty"`
	Name string  `json:"name,omitempty" bson:"name,omitempty"`
	DateCreated string  `json:"date_created,omitempty" bson:"date_created,omitempty"`
	Completed bool  `json:"completed" bson:"completed"`
}


type User struct {
	ID  primitive.ObjectID  `json:"_id,omitempty" bson:"_id,omitempty"`
	Email string  `json:"email,omitempty" bson:"email,omitempty"`
	Password string `json:"password,omitempty" bson:"password,omitempty"`
	DateCreated string  `json:"date_created,omitempty" bson:"date_created,omitempty"`
	Lists []List `json:"lists,omitempty" bson:"lists,omitempty"`
}

type TokenDetailsStruct struct {
	RefreshUuid  string
	AccessUuid   string
	RefreshToken string
	AccessToken  string
	RefreshExpiryDate int64
	AceessExpiryDate int64
}

type AccessDetail struct {
	AccessUuid string
	UserId     uint64
}




