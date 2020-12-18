package taskssevices


import (
	"big-todo-app/data-access/tododb"
	"big-todo-app/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func UndoTaskService(_id primitive.ObjectID, task_id primitive.ObjectID) (models.List) {
	result := tododb.UndoTask(_id, task_id )
	return result
}
