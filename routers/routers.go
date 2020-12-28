package routers

import (
	"big-todo-app/controller"
	"github.com/gorilla/mux"
	// "github.com/twinj/uuid"

	// "github.com/twinj/uuid"
	// "log"
	"net/http"
)

func HandleRoot(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	w.Write([]byte("Hello world"))
}
func Router() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/", HandleRoot).Methods("GET")
	router.HandleFunc("/lists/{id}", controller.PostList).Methods("POST", "OPTIONS")
	router.HandleFunc("/lists/{user_id}",controller.FetchLists).Methods("GET", "OPTIONS")
	router.HandleFunc("/edit-list/{id}/{user_id}", controller.PatchList).Methods("POST", "OPTIONS")
	router.HandleFunc("/delete-list/{id}/{user_id}", controller.RemoveList).Methods("GET", "OPTIONS")
	router.HandleFunc("/lists/{id}/{user_id}", controller.FetchList).Methods("GET", "OPTIONS")
	router.HandleFunc("/lists/add-task/{id}/{user_id}",controller.PostTask).Methods("POST", "OPTIONS")
	router.HandleFunc("/lists/edit-task/{id}/{task_id}/{user_id}", controller.EditTask).Methods("POST", "OPTIONS")
	router.HandleFunc("/lists/delete-task/{id}/{task_id}/{user_id}", controller.DeleteTaskController).Methods("GET", "OPTIONS")
	router.HandleFunc("/lists/complete-task/{id}/{task_id}/{user_id}", controller.CompleteTaskController).Methods("GET", "OPTIONS")
	router.HandleFunc("/lists/undo-task/{id}/{task_id}/{user_id}", controller.UndoTaskController).Methods("GET", "OPTIONS")
	router.HandleFunc("/login", controller.Login).Methods("POST", "OPTIONS")
	router.HandleFunc("/signout", controller.SignOut).Methods("POST", "OPTIONS")
	router.HandleFunc("/refresh", controller.Refresh).Methods("POST", "OPTIONS")
	router.HandleFunc("/signup", controller.Signup).Methods("POST", "OPTIONS")
	return router

}
