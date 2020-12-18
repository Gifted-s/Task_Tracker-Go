package tododb
import (
	"big-todo-app/data-access"
	"big-todo-app/models"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// var year, month, day = time.Now().Date()

// var todaysDate = month.String() + " " +  strconv.Itoa(day) + " " + strconv.Itoa(year)

func CreateList(listDetails models.List) (*mongo.InsertOneResult, error) {

	collections := createdb.ConnectDB()
	todo_collection := collections[0]
	result, err := todo_collection.InsertOne(context.TODO(), listDetails)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func AddTask(list_id primitive.ObjectID, taskDetails models.Task) (models.List) {

	var listDetails models.List
	filter := bson.M{"_id": list_id}
	update := bson.M{"$push": bson.M{"tasks": taskDetails}}
	collections := createdb.ConnectDB()
	todo_collection := collections[0]
	todo_collection.FindOneAndUpdate(context.TODO(), filter, update).Decode(&listDetails)
	return listDetails
}

func EditTask(list_id primitive.ObjectID, taskDetails models.Task) (models.List) {

	var listDetails models.List
	filter := bson.M{"_id": list_id, "tasks._id": taskDetails.ID}
	update := bson.M{"$set": bson.M{"tasks.$.name": taskDetails.Name}}
	collections := createdb.ConnectDB()
	todo_collection := collections[0]
	todo_collection.FindOneAndUpdate(context.TODO(), filter, update).Decode(&listDetails)
	return listDetails
}

func CompleteTask(list_id primitive.ObjectID, task_id primitive.ObjectID)  models.List {

	var listDetails models.List
	filter := bson.M{"_id": list_id, "tasks._id": task_id}
	update := bson.M{"$set": bson.M{"tasks.$.completed": true}}
	collections := createdb.ConnectDB()
	todo_collection := collections[0]
	todo_collection.FindOneAndUpdate(context.TODO(), filter, update).Decode(&listDetails)
	return listDetails
}

func UndoTask(list_id primitive.ObjectID, task_id primitive.ObjectID) models.List {

	var listDetails models.List
	filter := bson.M{"_id": list_id, "tasks._id": task_id}
	update := bson.M{"$set": bson.M{"tasks.$.completed": false}}
	collections := createdb.ConnectDB()
	todo_collection := collections[0]
	todo_collection.FindOneAndUpdate(context.TODO(), filter, update).Decode(&listDetails)
	return listDetails
}
func DeleteTask(list_id primitive.ObjectID, task_id primitive.ObjectID)  models.List {

	var listDetails models.List
	filter := bson.M{"_id":list_id}
	update := bson.M{"$pull": bson.M{"tasks": bson.M{"_id":task_id}}}
	collections := createdb.ConnectDB()
	todo_collection := collections[0]
	todo_collection.FindOneAndUpdate(context.TODO(), filter, update).Decode(&listDetails)
	return listDetails
}

func EditList(_id primitive.ObjectID, name string) models.List {
	var listDetails models.List
	filter := bson.M{"_id": _id}
	update := bson.M{"$set": bson.M{"name": name}}
	collections := createdb.ConnectDB()
	todo_collection := collections[0]
	todo_collection.FindOneAndUpdate(context.TODO(), filter, update).Decode(&listDetails)
	return listDetails
}

func GetList(_id primitive.ObjectID) models.List {
	var listDetails models.List
	filter := bson.M{"_id": _id}
	collections := createdb.ConnectDB()
	todo_collection := collections[0]
	todo_collection.FindOne(context.TODO(), filter).Decode(&listDetails)
	return listDetails
}

func GetAllList() ([]models.List, error) {
	var lists []models.List

	collections := createdb.ConnectDB()
	todo_collection := collections[0]
	cursor, err := todo_collection.Find(context.TODO(), bson.M{})
	defer cursor.Close(context.TODO())
	if err != nil {
		return nil, err
	}
	for cursor.Next(context.TODO()) {
		var list models.List
		err := cursor.Decode(&list)
		if err != nil {
			return nil, err
		}

		lists = append(lists, list)
	}
	return lists, nil
}

func DeleteList(_id primitive.ObjectID) (*mongo.DeleteResult, error) {
	filter := bson.M{"_id": _id}
	collections := createdb.ConnectDB()
	todo_collection := collections[0]
	result, err := todo_collection.DeleteOne(context.TODO(), filter)
	if err != nil {
		return nil, err
	}
	return result, nil
}
