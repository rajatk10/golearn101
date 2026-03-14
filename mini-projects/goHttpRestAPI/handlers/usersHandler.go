package handlers

import (
	"encoding/json"
	"errors"
	"github.com/asdine/storm"
	"goHttpRestAPI/user"
	"gopkg.in/mgo.v2/bson"
	"io/ioutil"
	"log"
	"net/http"
)

func bodyToUser(r *http.Request, u *user.User) error {
	if r == nil {
		log.Println("bodyToUser: nil http.Request")
		return errors.New("request is nil")
	}
	if r.Body == nil {
		return errors.New("body is nil")
	}
	if u == nil {
		return errors.New("a user is required")
	}
	bd, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return err
	}
	return json.Unmarshal(bd, u)

}
func usersGetAll(w http.ResponseWriter, r *http.Request) {
	users, err := user.All()
	if err != nil {
		postError(w, http.StatusInternalServerError)
		return
	}
	if r.Method == "HEAD" {
		//w.WriteHeader(http.StatusOK)
		postBodyResponse(w, http.StatusOK, jsonResponse{})
		return
	}
	postBodyResponse(w, http.StatusOK, jsonResponse{"users": users})
}

func usersGetOne(w http.ResponseWriter, r *http.Request, id bson.ObjectId) {
	//queryId := r.URL.Query().Get("id")
	log.Printf("Here is the id being queries %v\n", id)
	u, err := user.One(id)
	if err != nil {
		log.Printf("Error while fetching the user with id: %v", id)
		if err == storm.ErrNotFound {
			postError(w, http.StatusNotFound)
		} else {
			postError(w, http.StatusInternalServerError)
		}
		return
	}
	if r.Method == "HEAD" {
		//w.WriteHeader(http.StatusOK)
		postBodyResponse(w, http.StatusOK, jsonResponse{})
		return
	}
	postBodyResponse(w, http.StatusOK, jsonResponse{"user": u})
}

func userPostOne(w http.ResponseWriter, r *http.Request) {
	u := new(user.User)
	err := bodyToUser(r, u)
	if err != nil {
		postError(w, http.StatusBadRequest)
		return
	}
	u.ID = bson.NewObjectId()
	err = u.Save()
	if err != nil {
		if err == user.ErrRecordInvalid {
			postError(w, http.StatusBadRequest)
		} else {
			postError(w, http.StatusInternalServerError)
		}
		return
	}
	w.Header().Set("Location", "/users/"+u.ID.Hex())
	w.WriteHeader(http.StatusCreated)
	//postBodyResponse(w, http.StatusCreated, jsonResponse{"user": u})
}

func userPutOne(w http.ResponseWriter, r *http.Request, id bson.ObjectId) {
	log.Printf("Here is the id being updated via HTTP PUT %v\n", id)
	_, err := user.One(id)
	if err != nil {
		log.Printf("Error while fetching the user with id: %v", id)
		if err == storm.ErrNotFound {
			postError(w, http.StatusNotFound)
		} else {
			postError(w, http.StatusInternalServerError)
		}
		return
	}
	u := new(user.User)
	if err := bodyToUser(r, u); err != nil {
		postError(w, http.StatusBadRequest)
		return
	}
	u.ID = id
	if err := u.Save(); err != nil {
		postError(w, http.StatusInternalServerError)
		return
	}
	//update header
	w.Header().Set("Location", "/users/"+u.ID.Hex())
	//w.WriteHeader(http.StatusOK)
	//update responde body
	postBodyResponse(w, http.StatusOK, jsonResponse{"user": u})
}

func userPatchOne(w http.ResponseWriter, r *http.Request, id bson.ObjectId) {
	//Get existing user
	log.Printf("Here is the id being updated via HTTP PATCH %v\n", id)
	existing_user, err := user.One(id)
	if err != nil {
		log.Printf("Error while fetching the user with id: %v", id)
		if err == storm.ErrNotFound {
			postError(w, http.StatusNotFound)
		} else {
			postError(w, http.StatusInternalServerError)
		}
		return
	}
	//Parse partial data in map
	var updates map[string]interface{}
	if err := json.NewDecoder(r.Body).Decode(&updates); err != nil {
		postError(w, http.StatusBadRequest)
		return
	}
	if len(updates) == 0 {
		postError(w, http.StatusBadRequest)
		return
	}

	//Define allowed fields
	allowedFields := map[string]bool{
		"yearsOfExperience": true,
		"role":              true,
		"location":          true,
	}
	for k := range updates {
		if !allowedFields[k] {
			log.Printf("Invalid field: %s", k)
			postError(w, http.StatusBadRequest)
			return
		}
	}
	//update only provided field
	if name, ok := updates["yearsOfExperience"]; ok {
		// In json float64 is used by default so need to convert it to int while marshalling/unmarshalling
		existing_user.YearsOfExperience = int(name.(float64))
	}
	if name, ok := updates["role"]; ok {
		existing_user.Role = name.(string)
	}
	if name, ok := updates["location"]; ok {
		existing_user.Location = name.(string)
	}

	if err := existing_user.Save(); err != nil {
		postError(w, http.StatusInternalServerError)
		return
	}
	w.Header().Set("Location", "/users/"+existing_user.ID.Hex())
	w.WriteHeader(http.StatusOK)
	//update responde body
	postBodyResponse(w, http.StatusOK, jsonResponse{"user": existing_user})
}

func usersDeleteOne(w http.ResponseWriter, r *http.Request, id bson.ObjectId) {
	//Check if ID exists
	log.Printf("UsersHandler method: DELETE , Delete User %v\n", id)
	_, err := user.One(id)
	if err != nil {
		log.Printf("Error while fetching the user with id: %v", id)
		if err == storm.ErrNotFound {
			postError(w, http.StatusNotFound)
		} else {
			postError(w, http.StatusInternalServerError)
		}
		return
	}
	//Delete now
	err = user.Delete(id)
	if err != nil {
		log.Printf("Error while deleting the user with id: %v", id)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("User deleted successfully\n"))
}
