package listservices

import (
	"big-todo-app/data-access/tododb"
	"big-todo-app/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetListService(_id primitive.ObjectID) (models.List, error) {
	result := tododb.GetList(_id)
	return result, nil
}
