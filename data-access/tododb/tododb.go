package tododb

import (
	"big-todo-app/data-access"
	"big-todo-app/models"
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// var year, month, day = time.Now().Date()

// var todaysDate = month.String() + " " +  strconv.Itoa(day) + " " + strconv.Itoa(year)

func CreateList(user_id primitive.ObjectID, listDetails models.List) models.User {
	var user models.User
	user_collection := createdb.ConnectDB()
	filter := bson.M{"_id": user_id}
	update := bson.M{"$push": bson.M{"lists": listDetails}}
	user_collection.FindOneAndUpdate(context.TODO(), filter, update).Decode(&user)
	return user
}

func AddTask(list_id primitive.ObjectID, user_id primitive.ObjectID, taskDetails models.Task) models.User {
	var user models.User
	filter := bson.M{"_id": user_id, "lists._id": list_id}
	update := bson.M{"$push": bson.M{"lists.$.tasks": taskDetails}}
	user_collection := createdb.ConnectDB()
	user_collection.FindOneAndUpdate(context.TODO(), filter, update).Decode(&user)
	return user
}

func EditTask(list_id primitive.ObjectID, user_id primitive.ObjectID, taskDetails models.Task) models.User {
	// var user models.User
	user_collection := createdb.ConnectDB()
	filter := bson.M{
		"_id":       user_id,
		"lists._id": list_id,
	}
	updateQuery := bson.M{"$set": bson.M{"lists.$.tasks.$[elem].name": taskDetails.Name}}
	options := options.Update().SetArrayFilters(options.ArrayFilters{
		Filters: []interface{}{bson.M{
			"elem._id": taskDetails.ID,
		}},
	})
	user_collection.UpdateOne(
		context.TODO(),
		filter,
		updateQuery,
		options,
	)
	user := GetUserById(user_id)
	return user

}
func CompleteTask(list_id primitive.ObjectID, user_id primitive.ObjectID, task_id primitive.ObjectID) models.User {
	user_collection := createdb.ConnectDB()
	filter := bson.M{
		"_id":       user_id,
		"lists._id": list_id,
	}
	updateQuery := bson.M{"$set": bson.M{"lists.$.tasks.$[elem].completed": true}}
	options := options.Update().SetArrayFilters(options.ArrayFilters{
		Filters: []interface{}{bson.M{
			"elem._id": task_id,
		}},
	})
	user_collection.UpdateOne(
		context.TODO(),
		filter,
		updateQuery,
		options,
	)
	user := GetUserById(user_id)
	return user
}

func UndoTask(list_id primitive.ObjectID, user_id primitive.ObjectID, task_id primitive.ObjectID) models.User {
	user_collection := createdb.ConnectDB()
	filter := bson.M{
		"_id":       user_id,
		"lists._id": list_id,
	}
	updateQuery := bson.M{"$set": bson.M{"lists.$.tasks.$[elem].completed": false}}
	options := options.Update().SetArrayFilters(options.ArrayFilters{
		Filters: []interface{}{bson.M{
			"elem._id": task_id,
		}},
	})
	user_collection.UpdateOne(
		context.TODO(),
		filter,
		updateQuery,
		options,
	)
	user := GetUserById(user_id)
	return user
}
func DeleteTask(list_id primitive.ObjectID, user_id primitive.ObjectID, task_id primitive.ObjectID) models.User {
    fmt.Println(list_id, user_id, task_id)
	var user models.User
	filter := bson.M{"_id": user_id, "lists._id": list_id}
	update := bson.M{"$pull": bson.M{"lists.$.tasks": bson.M{"_id": task_id}}}
	user_collection := createdb.ConnectDB()
	user_collection.FindOneAndUpdate(context.TODO(), filter, update).Decode(&user)
	return user
}


func EditList(_id primitive.ObjectID, user_id primitive.ObjectID, name string) models.User {
	var user models.User
	filter := bson.M{"_id": user_id, "lists._id": _id}
	update := bson.M{"$set": bson.M{"lists.$.name": name}}
	collections := createdb.ConnectDB()
	user_collection := collections
	user_collection.FindOneAndUpdate(context.TODO(), filter, update).Decode(&user)
	return user
}

func GetList(_id primitive.ObjectID, user_id primitive.ObjectID) models.User {
	var user models.User
	filter := bson.M{"_id": user_id, "lists._id": _id}
	collections := createdb.ConnectDB()
	user_collection := collections
	user_collection.FindOne(context.TODO(), filter).Decode(&user)
	return user
}

func GetAllList(user_id primitive.ObjectID) models.User {
	var user models.User
	collections := createdb.ConnectDB()
	user_collection := collections
	filter := bson.M{"_id": user_id}
	user_collection.FindOne(context.TODO(), filter).Decode(&user)
	return user
}

func DeleteList(_id primitive.ObjectID, user_id primitive.ObjectID) models.User {
	var user models.User
	user_collection := createdb.ConnectDB()
	filter := bson.M{"_id": user_id}
	update := bson.M{"$pull": bson.M{"lists": bson.M{"_id": _id}}}
	user_collection.FindOneAndUpdate(context.TODO(), filter, update).Decode(&user)
	return user
}

func InsertUser(user models.User) (*mongo.InsertOneResult, error) {
	user_collection := createdb.ConnectDB()
	result, err := user_collection.InsertOne(context.TODO(), user)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func GetUser(email string) models.User {
	var user models.User
	filter := bson.M{"email": email}
	collections := createdb.ConnectDB()
	user_collection := collections
	user_collection.FindOne(context.TODO(), filter).Decode(&user)
	return user
}

func GetUserById(id primitive.ObjectID) models.User {
	var user models.User
	filter := bson.M{"_id": id}
	collections := createdb.ConnectDB()
	user_collection := collections
	user_collection.FindOne(context.TODO(), filter).Decode(&user)
	return user
}
