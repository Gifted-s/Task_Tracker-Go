package taskssevices


import (
	"big-todo-app/data-access/tododb"
	"big-todo-app/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CompleteTaskService(_id primitive.ObjectID,user_id primitive.ObjectID, task_id primitive.ObjectID) (models.User) {
	result := tododb.CompleteTask(_id,user_id, task_id )
	return result
}
