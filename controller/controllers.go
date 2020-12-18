package controller

import (
	"big-todo-app/models"
	"big-todo-app/modules"
	"big-todo-app/modules/responses"
	"big-todo-app/services/auth"
	"big-todo-app/services/listservices"
	"big-todo-app/services/taskssevices"
	"encoding/json"

	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID       uint64 `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

var user = User{
	ID:       1,
	Username: "username",
	Password: "password",
}

func Login(w http.ResponseWriter, r *http.Request) {
	var newUser User
	json.NewDecoder(r.Body).Decode(&newUser)
	if newUser.Password != user.Password || user.Username != newUser.Username {
		responses.BadRequest(w, "Invalid Password Or user name")
		return
	}
	td, err := auth.CreateToken(newUser.ID)
	if err != nil {
		responses.BadRequest(w, "Token was not saved")
	}
	saveError := auth.SaveToRedis(newUser.ID, td)
	if saveError != nil {
		responses.BadRequest(w, "Token Detials failed to save to redis")
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"AccessToken":  td.AccessToken,
		"RefreshToken": td.RefreshToken,
		"Username":     newUser.Username,
		"ID":           strconv.Itoa(int(user.ID)),
	})

}


func SignOut(w http.ResponseWriter, r *http.Request){
	access_details, err := auth.ExtractTokenMeta(r)
	if err != nil {
		responses.BadRequest(w, "Unathorized")
		return
	}

	deleteResponse, saveErr := auth.DeleteAt(access_details.AccessUuid)
	if saveErr != nil {
		responses.BadRequest(w, "Unathorized")
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(map[string]int64{"deleted_count": deleteResponse})
}

func PostList(w http.ResponseWriter, r *http.Request) {
	access_details, err := auth.ExtractTokenMeta(r)
	if err != nil {
		responses.BadRequest(w, "Unathorized")
		return
	}

	_, errToFetch := auth.FetchAuth(access_details)
	if errToFetch != nil {
		responses.BadRequest(w, "Unathorized")
		return
	}
	var list models.List
	json.NewDecoder(r.Body).Decode(&list)
	response, err := listservices.CreateListService(list)
	if err != nil {
		responses.ServerError(w)
	}
	responses.SuccessResponse(w, 200, response)
}

func RemoveList(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	_id, _ := primitive.ObjectIDFromHex(params["id"])
	response, err := listservices.DeleteListService(_id)
	if err != nil {
		responses.ServerError(w)
	}
	responses.SuccessResponse(w, 200, response)
}

func PatchList(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	listMap := map[string]string{}
	json.NewDecoder(r.Body).Decode(&listMap)
	_id, _ := primitive.ObjectIDFromHex(params["id"])

	response := listservices.EditListService(_id, listMap["Name"])

	responses.SuccessResponse(w, 200, response)
}
func PostTask(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var task models.Task
	json.NewDecoder(r.Body).Decode(&task)
	_id, _ := primitive.ObjectIDFromHex(params["id"])
	task.DateCreated = modules.GetTodayDate()
	task.ID = primitive.NewObjectID()

	response := taskssevices.AddTaskService(_id, task)

	responses.SuccessResponse(w, 200, response)
}
func EditTask(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var task models.Task
	json.NewDecoder(r.Body).Decode(&task)
	_id, _ := primitive.ObjectIDFromHex(params["id"])
	task_id, _ := primitive.ObjectIDFromHex(params["task_id"])
	task.ID = task_id
	response := taskssevices.EditTaskService(_id, task)
	responses.SuccessResponse(w, 200, response)
}
func CompleteTaskController(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	_id, _ := primitive.ObjectIDFromHex(params["id"])
	task_id, _ := primitive.ObjectIDFromHex(params["task_id"])

	response := taskssevices.CompleteTaskService(_id, task_id)
	responses.SuccessResponse(w, 200, response)
}

func UndoTaskController(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	_id, _ := primitive.ObjectIDFromHex(params["id"])
	task_id, _ := primitive.ObjectIDFromHex(params["task_id"])

	response := taskssevices.UndoTaskService(_id, task_id)
	responses.SuccessResponse(w, 200, response)
}

func DeleteTaskController(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	_id, _ := primitive.ObjectIDFromHex(params["id"])
	task_id, _ := primitive.ObjectIDFromHex(params["task_id"])
	response := taskssevices.DeleteTaskService(_id, task_id)
	responses.SuccessResponse(w, 200, response)
}

func FetchLists(w http.ResponseWriter, r *http.Request) {
	response, err := listservices.GetListsService()
	if err != nil {
		responses.ServerError(w)
	}
	responses.SuccessResponse(w, 200, response)
}

func FetchList(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	_id, _ := primitive.ObjectIDFromHex(params["id"])

	response, err := listservices.GetListService(_id)
	if err != nil {
		responses.ServerError(w)
	}

	responses.SuccessResponse(w, 200, response)
}



func Refresh(w http.ResponseWriter, r *http.Request) {
	mapToken := map[string]string{}
	decoder := json.NewDecoder(r.Body)
	decoder.Decode(&mapToken)
	refreshToken := mapToken["refresh_token"]
	tokens, err := auth.HandleRefreshToken(refreshToken)
	if err != "" {
		responses.BadRequest(w,err)
	}
	responses.SuccessResponse(w, 200, tokens)	
}

