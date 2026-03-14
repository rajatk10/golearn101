package handlers

import (
	"log"
	"net/http"
	"strings"

	"gopkg.in/mgo.v2/bson"
)

func UsersRouter(w http.ResponseWriter, r *http.Request) {
	//fmt.Printf("UsersRouter : %s\n", r.URL.Path)
	//fmt.Println("UsersRouterMethod : ", r.Method)
	path := strings.TrimSuffix(r.URL.Path, "/")
	if path == "/users" {
		switch r.Method {
		case "GET":
			log.Println("UsersRouter method: GET , GetAll Users")
			usersGetAll(w, r)
			return
		case "POST":
			log.Println("UsersRouter method: POST , Create User")
			userPostOne(w, r)
			return
		case "HEAD":
			log.Println("UsersRouter method: HEAD , Return empty if resource exists")
			usersGetAll(w, r)
			return
		case "OPTIONS":
			log.Println("UsersRouter method: OPTIONS , Return allowed methods")
			postOptionsResponse(w, []string{"GET", "POST", "HEAD", "OPTIONS"}, jsonResponse{})
			return
		default:
			postError(w, http.StatusMethodNotAllowed)
			return
		}
	}
	path = strings.TrimPrefix(path, "/users/")
	if !bson.IsObjectIdHex(path) {
		postError(w, http.StatusNotFound)
		return
	}
	id := bson.ObjectIdHex(path)
	switch r.Method {
	case http.MethodGet:
		log.Println("UsersRouter method: GET , Get One User by ID")
		usersGetOne(w, r, id)
		return
	case http.MethodHead:
		log.Println("UsersRouter method: HEAD , Return empty if resource exists")
		usersGetOne(w, r, id)
		return
	case http.MethodDelete:
		log.Println("UsersRouter method: DELETE , Delete One User by ID")
		usersDeleteOne(w, r, id)
		return
	case http.MethodPut:
		log.Println("UsersRouter method: PUT , Update One User by ID")
		userPutOne(w, r, id)
		return
	case http.MethodPatch:
		log.Println("UsersRouter method: PATCH , Update Partial User by ID")
		userPatchOne(w, r, id)
		return
	case http.MethodOptions:
		log.Println("UsersRouter method: OPTIONS , Return allowed methods for User by ID")
		postOptionsResponse(w, []string{"GET", "PUT", "PATCH", "DELETE", "HEAD", "OPTIONS"}, jsonResponse{})
		return
	default:
		postError(w, http.StatusMethodNotAllowed)
		return
	}
}
