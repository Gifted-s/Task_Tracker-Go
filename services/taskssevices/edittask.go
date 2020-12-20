package taskssevices


import (
	"big-todo-app/data-access/tododb"
	"big-todo-app/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func EditTaskService(_id primitive.ObjectID,user_id primitive.ObjectID, taskDetails models.Task) (models.User) {
	result := tododb.EditTask(_id,user_id, taskDetails )
	return result
}
