package listservices

import (
	"big-todo-app/data-access/tododb"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func DeleteListService(_id primitive.ObjectID) (*mongo.DeleteResult, error) {

	result, err := tododb.DeleteList(_id)
	if err != nil {
		return nil, err
	}
	return result, nil
}
