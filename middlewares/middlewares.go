package middlewares

import (
	"big-todo-app/services/auth"
	"net/http"
	"big-todo-app/modules/responses"
)
func TokenAuthMiddleware(handler func(w http.ResponseWriter, r *http.Request)) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		err := auth.CheckIfTokenStillValid(r)
		if err != nil {
		    responses.BadRequest(w,"Unauthorized")
			return
		}
		handler(w, r)
	}
}