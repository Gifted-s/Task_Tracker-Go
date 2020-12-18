package taskssevices


import (
	"big-todo-app/data-access/tododb"
	"big-todo-app/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func AddTaskService(_id primitive.ObjectID, taskDetails models.Task) (models.List) {
	result := tododb.AddTask(_id, taskDetails )
	return result
}
