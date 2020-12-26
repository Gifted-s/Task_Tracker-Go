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

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson/primitive"
)


func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, PATCH")
}



func SignOut(w http.ResponseWriter, r *http.Request){
	enableCors(&w)
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
	return
}

func PostList(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	// access_details, err := auth.ExtractTokenMeta(r)
	// if err != nil {
	// 	responses.BadRequest(w, "Unathorized")
	// 	return
	// }

	// _, errToFetch := auth.FetchAuth(access_details)
	// if errToFetch != nil {
	// 	responses.BadRequest(w, "Unathorized")
	// 	return
	// }
	var list models.List
	params := mux.Vars(r)
	_id, _ := primitive.ObjectIDFromHex(params["id"])
   
	json.NewDecoder(r.Body).Decode(&list)
	response := listservices.CreateListService(_id, list)

	responses.SuccessResponse(w, 200, response)
}

func RemoveList(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	params := mux.Vars(r)
	_id, _ := primitive.ObjectIDFromHex(params["id"])
	user_id, _ :=primitive.ObjectIDFromHex(params["user_id"])
	response:= listservices.DeleteListService(_id, user_id)

	responses.SuccessResponse(w, 200, response)
}

func PatchList(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	params := mux.Vars(r)
	listMap := map[string]string{}
	json.NewDecoder(r.Body).Decode(&listMap)
	_id, _ := primitive.ObjectIDFromHex(params["id"])
	user_id, _ :=primitive.ObjectIDFromHex(params["user_id"])
	response := listservices.EditListService(_id,user_id, listMap["Name"])

	responses.SuccessResponse(w, 200, response)
}
func PostTask(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	params := mux.Vars(r)
	var task models.Task
	json.NewDecoder(r.Body).Decode(&task)
	_id, _ := primitive.ObjectIDFromHex(params["id"])
	user_id, _ :=primitive.ObjectIDFromHex(params["user_id"])
	task.DateCreated = modules.GetTodayDate()
	task.ID = primitive.NewObjectID()

	response := taskssevices.AddTaskService(_id,user_id, task)

	responses.SuccessResponse(w, 200, response)
}
func EditTask(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	params := mux.Vars(r)
	var task models.Task
	json.NewDecoder(r.Body).Decode(&task)
	_id, _ := primitive.ObjectIDFromHex(params["id"])
	task_id, _ := primitive.ObjectIDFromHex(params["task_id"])
	user_id, _ :=primitive.ObjectIDFromHex(params["user_id"])
	task.ID = task_id
	response:= taskssevices.EditTaskService(_id,user_id, task)

	responses.SuccessResponse(w, 200, response)
}
func CompleteTaskController(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	params := mux.Vars(r)

	_id, _ := primitive.ObjectIDFromHex(params["id"])
	task_id, _ := primitive.ObjectIDFromHex(params["task_id"])
	user_id, _ :=primitive.ObjectIDFromHex(params["user_id"])
	response := taskssevices.CompleteTaskService(_id,user_id, task_id)
	responses.SuccessResponse(w, 200, response)
}

func UndoTaskController(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	params := mux.Vars(r)

	_id, _ := primitive.ObjectIDFromHex(params["id"])
	task_id, _ := primitive.ObjectIDFromHex(params["task_id"])
	user_id, _ :=primitive.ObjectIDFromHex(params["user_id"])
	response := taskssevices.UndoTaskService(_id,user_id, task_id)
	responses.SuccessResponse(w, 200, response)
}

func DeleteTaskController(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	params := mux.Vars(r)
	_id, _ := primitive.ObjectIDFromHex(params["id"])
	task_id, _ := primitive.ObjectIDFromHex(params["task_id"])
	user_id, _ :=primitive.ObjectIDFromHex(params["user_id"])
	response := taskssevices.DeleteTaskService(_id,user_id, task_id)
	responses.SuccessResponse(w, 200, response)
}

func FetchLists(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	params := mux.Vars(r)
	user_id, _ :=primitive.ObjectIDFromHex(params["user_id"])
	response, err := listservices.GetListsService(user_id)
	if err != nil {
		responses.ServerError(w)
	}
	responses.SuccessResponse(w, 200, response)
}

func FetchList(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	params := mux.Vars(r)
	_id, _ := primitive.ObjectIDFromHex(params["id"])
	user_id, _ :=primitive.ObjectIDFromHex(params["user_id"])
	response := listservices.GetListService(_id,user_id)


	responses.SuccessResponse(w, 200, response)
}



func Refresh(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	mapToken := map[string]string{}
	decoder := json.NewDecoder(r.Body)
	decoder.Decode(&mapToken)
	refreshToken := mapToken["refresh_token"]
	tokens, err := auth.HandleRefreshToken(refreshToken)
	if err != "" {
		responses.BadRequest(w,err)
	}
	responses.SuccessResponse(w, 200, tokens)	
	return
}


func Signup(w  http.ResponseWriter, r *http.Request){
	var newUser models.User
	json.NewDecoder(r.Body).Decode(&newUser)
	inserted, err := auth.HandleSignup(newUser)
	if err != "" {
		responses.BadRequest(w,err)
		return
	}
	responses.SuccessResponse(w, 200, inserted)
	return
}

func Login(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	var newUser models.User
	json.NewDecoder(r.Body).Decode(&newUser)
	
	userExist, loginError := auth.HandleLogin(newUser)
	if loginError != "" {
		responses.BadRequest(w, loginError)
		return
	}
	// td, err := auth.CreateToken(userExist.ID)
	// if err != nil {
	// 	responses.BadRequest(w, "Token was not saved")
	// }
	// saveError := auth.SaveToRedis(userExist.ID, td)
	// if saveError != nil {
	// 	responses.BadRequest(w, "Token Detials failed to save to redis")
	// }
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(userExist)

}

