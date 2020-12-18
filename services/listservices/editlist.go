package listservices

import (
	"big-todo-app/data-access/tododb"
	"big-todo-app/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)
func EditListService(_id primitive.ObjectID, name string) (models.List) {
	result := tododb.EditList(_id, name )
	return result
}
