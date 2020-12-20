package listservices

import (
	"big-todo-app/data-access/tododb"
	"big-todo-app/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)
func EditListService(_id primitive.ObjectID,user_id primitive.ObjectID, name string) (models.User) {
	result := tododb.EditList(_id,user_id, name )
	return result
}
