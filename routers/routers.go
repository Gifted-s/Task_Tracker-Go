package routers

import (
	"big-todo-app/controller"
	"big-todo-app/middlewares"
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
	router.HandleFunc("/lists/{id}/{user_id}", controller.PatchList).Methods("PATCH", "OPTIONS")
	router.HandleFunc("/lists/{id}/{user_id}", controller.RemoveList).Methods("DELETE", "OPTIONS")
	router.HandleFunc("/lists/{id}/{user_id}", controller.FetchList).Methods("GET", "OPTIONS")
	router.HandleFunc("/lists/add-task/{id}/{user_id}",controller.PostTask).Methods("POST", "OPTIONS")
	router.HandleFunc("/lists/edit-task/{id}/{task_id}/{user_id}", controller.EditTask).Methods("PATCH", "OPTIONS")
	router.HandleFunc("/lists/delete-task/{id}/{task_id}/{user_id}", controller.DeleteTaskController).Methods("DELETE", "OPTIONS")
	router.HandleFunc("/lists/complete-task/{id}/{task_id}/{user_id}", controller.CompleteTaskController).Methods("PATCH", "OPTIONS")
	router.HandleFunc("/lists/undo-task/{id}/{task_id}/{user_id}", controller.UndoTaskController).Methods("PATCH", "OPTIONS")
	router.HandleFunc("/login", controller.Login).Methods("POST", "OPTIONS")
	router.HandleFunc("/signout", middlewares.TokenAuthMiddleware(controller.SignOut)).Methods("POST", "OPTIONS")
	router.HandleFunc("/refresh", controller.Refresh).Methods("POST", "OPTIONS")
	router.HandleFunc("/signup", controller.Signup).Methods("POST", "OPTIONS")
	return router

}
