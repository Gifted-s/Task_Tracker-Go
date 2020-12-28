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
	"github.com/rs/cors"
	"log"
	"net/http"
	"big-todo-app/routers"
	// "os"
	// "strconv"
	// "strings"
	// "time"
)





func main(){
	router := routers.Router()
	fmt.Println("Server is Listening")
    c := cors.New(cors.Options{
        AllowedOrigins: []string{"*"},
        AllowCredentials: true,
    })

    handler := c.Handler(router)
    log.Fatal(http.ListenAndServe(":4000", handler))
}