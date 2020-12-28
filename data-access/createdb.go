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
  "os"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

 func ConnectDB() *mongo.Collection{
  //  clientOptions := options.Client().ApplyURI("mongodb://127.0.0.1:27017")
  clientOptions := options.Client().ApplyURI(os.Getenv("MONGODB_URL"))
   client, err := mongo.Connect(context.TODO(), clientOptions)

   if err != nil {
	   log.Fatal(err)
   }

 
   user_collection := client.Database("task_tracker").Collection("users")
    
   return  user_collection
 }