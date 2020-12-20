package listservices

import (
	"big-todo-app/data-access/tododb"
	"big-todo-app/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetListsService(user_id primitive.ObjectID) (models.User, error) {
	result := tododb.GetAllList(user_id)

	return result, nil
}
