package main
import (
	// "time"
	// "strconv"
	// "encoding/json"
	// "encoding/json"
	"fmt"
	// "strconv"
	// "strings"
	// "time"

	// jwt "github.com/dgrijalva/jwt-go"
	// "github.com/dgrijalva/jwt-go"
	// "github.com/go-redis/redis/v7"
	// "big-todo-app/services/listservices"
	// "github.com/twinj/uuid"

	// "github.com/twinj/uuid"
	// "log"
	"net/http"
	"big-todo-app/routers"
	"github.com/gorilla/handlers"
	// "os"
	// "strconv"
	// "strings"
	// "time"
)





func main(){
	router := routers.Router()
	fmt.Println("Server is Listening")
	http.ListenAndServe(":4000",handlers.CORS(handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"}), handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "OPTIONS"}), handlers.AllowedOrigins([]string{"*"})(router))
}