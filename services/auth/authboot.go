package auth

import (
	// "time"
	"strconv"
	// "encoding/json"
	// "encoding/json"
	"fmt"
	"net/http"

	// "strconv"
	"big-todo-app/data-access/tododb"
	"big-todo-app/models"
	"big-todo-app/modules"
	"strings"

	"time"
	// "go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"

	jwt "github.com/dgrijalva/jwt-go"
    
	// "github.com/dgrijalva/jwt-go"
	"github.com/go-redis/redis/v7"
	// "big-todo-app/services/listservices"
	"github.com/twinj/uuid"
	// "github.com/twinj/uuid"
	// "log"
	"os"
	// "strconv"
	// "strings"
	// "time"
)

var client *redis.Client

//Initializing redis

func ConnectRedis() {
	dsn := os.Getenv("REDIS_DSN")
	if len(dsn) == 0 {
		dsn = "localhost:6379"
	}
	client = redis.NewClient(&redis.Options{
		Addr: dsn, //redis port
	})
	_, err := client.Ping().Result()
	if err != nil {
		panic(err)
	}
}

func CreateToken(userId uint64) (*models.TokenDetailsStruct, error) {

	ConnectRedis()

	td := &models.TokenDetailsStruct{}
	td.AccessUuid = uuid.NewV4().String()
	td.RefreshUuid = uuid.NewV4().String()
	td.AceessExpiryDate = time.Now().Add(time.Minute * 15).Unix()
	td.RefreshExpiryDate = time.Now().Add(time.Hour * 24 * 7).Unix()
	os.Setenv("ACCESS_KEY", "1234")
	atClaims := jwt.MapClaims{}
	atClaims["authorized"] = true
	atClaims["user_id"] = userId
	atClaims["exp"] = td.AceessExpiryDate
	atClaims["access_uuid"] = td.AccessUuid
	access_token := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	at, err := access_token.SignedString([]byte(os.Getenv("ACCESS_KEY")))
	if err != nil {
		return nil, err
	}
	td.AccessToken = at

	os.Setenv("REFRESH_KEY", "1234")
	rtClaims := jwt.MapClaims{}

	rtClaims["user_id"] = userId
	rtClaims["exp"] = td.RefreshExpiryDate
	rtClaims["refresh_uuid"] = td.RefreshUuid
	refresh_token := jwt.NewWithClaims(jwt.SigningMethodHS256, rtClaims)
	rt, err := refresh_token.SignedString([]byte(os.Getenv("REFRESH_KEY")))
	if err != nil {
		return nil, err
	}
	td.RefreshToken = rt

	return td, nil
}

func SaveToRedis(userID uint64, td *models.TokenDetailsStruct) error {
	access_expiry := time.Unix(td.AceessExpiryDate, 0)
	refresh_expiry := time.Unix(td.RefreshExpiryDate, 0)
	now := time.Now()

	err1 := client.Set(td.AccessUuid, strconv.Itoa(int(userID)), access_expiry.Sub(now)).Err()
	if err1 != nil {
		return err1
	}
	err2 := client.Set(td.RefreshUuid, strconv.Itoa(int(userID)), refresh_expiry.Sub(now)).Err()
	if err2 != nil {
		return err2
	}

	return nil
}

func ExtractToken(r *http.Request) string {
	bearerToken := r.Header.Get("Authorization")
	splitedToken := strings.Split(bearerToken, " ")
	if len(splitedToken) == 2 {
		return splitedToken[1]
	}

	return ""
}

func CheckSigninMethod(r *http.Request) (*jwt.Token, error) {
	extractedToken := ExtractToken(r)
	token, err := jwt.Parse(extractedToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("String cannot be used for this signin method")
		}
		return []byte(os.Getenv("ACCESS_KEY")), nil
	})
	if err != nil {
		return nil, err
	}
	return token, nil
}

func CheckIfTokenStillValid(r *http.Request) error {
	token, err := CheckSigninMethod(r)
	if err != nil {
		return err
	}
	_, ok := token.Claims.(jwt.Claims)
	if !ok && !token.Valid {
		return err
	}
	return nil
}

func ExtractTokenMeta(r *http.Request) (*models.AccessDetail, error) {
	token, err := CheckSigninMethod(r)
	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		uuid, ok := claims["access_uuid"].(string)
		if !ok {
			return nil, err
		}
		userId, err := strconv.ParseUint(fmt.Sprintf("%.f", claims["user_id"]), 10, 64)
		if err != nil {
			return nil, err
		}

		return &models.AccessDetail{
			AccessUuid: uuid,
			UserId:     userId,
		}, nil
	}
	return nil, err

}

func FetchAuth(access_detail *models.AccessDetail) (uint64, error) {

	dsn := os.Getenv("REDIS_DSN")
	if len(dsn) == 0 {
		dsn = "localhost:6379"
	}
	client = redis.NewClient(&redis.Options{
		Addr: dsn, //redis port
	})
	_, err := client.Ping().Result()
	if err != nil {
		panic(err)
	}

	user_id, err := client.Get(access_detail.AccessUuid).Result()
	if err != nil {
		return 0, err
	}
	result, err := strconv.ParseUint(user_id, 10, 64)
	if err != nil {
		return 0, err
	}
	return result, nil
}



func DeleteAt(access_uuid string)(int64, error){
	dsn := os.Getenv("REDIS_DSN")
	if len(dsn) == 0 {
		dsn = "localhost:6379"
	}
	client = redis.NewClient(&redis.Options{
		Addr: dsn, //redis port
	})
	_, err := client.Ping().Result()
	if err != nil {
		panic(err)
	}
 result,err := client.Del(access_uuid).Result()
 if err !=nil {
	 return 0, err
 }
 return result,nil
}


func HandleRefreshToken(refreshToken  string)(map[string]string, string){
	//verify the token
	os.Setenv("REFRESH_KEY", "1234") //this should be in an env file
	token, err := jwt.Parse(refreshToken, func(token *jwt.Token) (interface{}, error) {
		//Make sure that the token method conform to "SigningMethodHMAC"
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		
		return []byte(os.Getenv("REFRESH_KEY")), nil
	})
	//if there is an error, the token must have expired
	if err != nil {
		
		return nil, "Token cannot be used for this signing method"
	}
	//is token valid?
	if _, ok := token.Claims.(jwt.Claims); !ok && !token.Valid {
		return nil, "String is no longer valid"
	
	}
	//Since token is valid, get the uuid:
	claims, ok := token.Claims.(jwt.MapClaims) //the token claims should conform to MapClaims
	if ok && token.Valid {
		refreshUuid, ok := claims["refresh_uuid"].(string) //convert the interface to string
		if !ok {
			
			return nil, "Failed to get refresh uuid"
		
		}
		userId, err := strconv.ParseUint(fmt.Sprintf("%.f", claims["user_id"]), 10, 64)
		if err != nil {
			return nil, "Failed to get refresh user id"
		}
		//Delete the previous Refresh Token
		deleted, delErr := DeleteAt(refreshUuid)
		if delErr != nil || deleted == 0 { //if any goes wrong
			return nil, "Error deleting prevoius refresh token"
		}
		//Create new pairs of refresh and access tokens
		ts, createErr := CreateToken(userId)
		if createErr != nil {
			return nil, "Error creating new tokens"
		}
		//save the tokens metadata to redis
		saveErr := SaveToRedis(userId, ts)
		if saveErr != nil {
			return nil, "Error creating new tokens"
		}
		tokens := map[string]string{
			"Access_Token":  ts.AccessToken,
			"Refresh_Token": ts.RefreshToken,
		}
		return tokens , ""
	} else {
		return nil, "Token is invalid"
	}
}

func HandleSignup(user models.User)(*mongo.InsertOneResult, string){
userExist := tododb.GetUser(user.Email)
if userExist.Email != "" {
	return nil, "User already exist"
}
hashedPassword,errToHash := modules.HashPassword(user.Password)
if errToHash!= nil {
	return nil, "Error hashing password"
}
user.Password = hashedPassword
user.DateCreated = modules.GetTodayDate()
insertResponse, err :=tododb.InsertUser(user)
if err!= nil {
	return nil, "Error Inserting User"
}

return insertResponse, ""
}


func HandleLogin(user models.User)(models.User, string){
	userExist := tododb.GetUser(user.Email)
	if userExist.Email == "" {
		return user, "User does e notxist"
	}
	
	passwordValid := modules.CheckPasswordHash(user.Password, userExist.Password )
	if !passwordValid {
		return user, "Error hashing password"
	}

	
	return userExist, ""
	}
	