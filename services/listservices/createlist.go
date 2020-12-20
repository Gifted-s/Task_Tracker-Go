package listservices

import (
	"big-todo-app/models"
	"big-todo-app/modules"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"big-todo-app/data-access/tododb"
)

func CreateListService(user_id primitive.ObjectID, listDetails models.List) (models.User) {
	listDetails.DateCreated = modules.GetTodayDate()
	listDetails.Tasks[0].DateCreated = modules.GetTodayDate()
	listDetails.Tasks[0].ID = primitive.NewObjectID()
	listDetails.ID = primitive.NewObjectID()
	result := tododb.CreateList(user_id, listDetails)

	return result
}
