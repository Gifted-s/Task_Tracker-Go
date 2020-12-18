package createdb

// "fmt"
// "time"
// "strconv"
// "net/http"
// "github.com/gorilla/mux"
// "log"
// "context"
// "github.com/Gifted-s/todo/models"
// "encoding/json"
// "go.mongodb.org/mongo-driver/bson"
// "go.mongodb.org/mongo-driver/bson/primitive"
// "go.mongodb.org/mongo-driver/mongo/options"
// "go.mongodb.org/mongo-driver/mongo"
import (
	"context"
	"log"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

 func ConnectDB() [2]*mongo.Collection{
   clientOptions := options.Client().ApplyURI("mongodb://127.0.0.1:27017")
   client, err := mongo.Connect(context.TODO(), clientOptions)

   if err != nil {
	   log.Fatal(err)
   }

   todo_collection  := client.Database("sunky_todo").Collection("todos")
   user_collection := client.Database("sunky_todo").Collection("users")
    
   return [2]*mongo.Collection{todo_collection, user_collection}
 }