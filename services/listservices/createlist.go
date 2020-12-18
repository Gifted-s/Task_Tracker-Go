package listservices

import (
	"big-todo-app/models"
	"big-todo-app/modules"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"big-todo-app/data-access/tododb"
)

func CreateListService(listDetails models.List) (*mongo.InsertOneResult, error) {
	listDetails.DateCreated = modules.GetTodayDate()
	listDetails.Tasks[0].DateCreated = modules.GetTodayDate()
	listDetails.Tasks[0].ID = primitive.NewObjectID()
	result, err := tododb.CreateList(listDetails)
	if err != nil {
		return nil, err
	}
	return result, nil
}
