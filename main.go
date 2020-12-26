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
	header:=handlers.AllowedOrigins([]string{"X-Requested-With","Content-Type", "Authorization" })
	methods:= handlers.AllowedMethods([]string{"GET", "POST", "DELETE", "PATCH", "PUT"})
	origins:= handlers.AllowedOrigins([]string{"*"})
	http.ListenAndServe(":4000", handlers.CORS(header, methods, origins)(router))
}