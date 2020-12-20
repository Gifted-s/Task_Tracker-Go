package listservices

import (
	"big-todo-app/data-access/tododb"
	"big-todo-app/models"
	"go.mongodb.org/mongo-driver/bson/primitive"	
)

func DeleteListService(_id primitive.ObjectID, user_id primitive.ObjectID) (models.User) {

	result:= tododb.DeleteList(_id, user_id)
	
	return result
}
