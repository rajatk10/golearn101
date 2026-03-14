package user

import (
	"errors"
	"log"

	"github.com/asdine/storm"
	"gopkg.in/mgo.v2/bson"
)

// Handling user data
type User struct {
	ID                bson.ObjectId `json:"id" storm:"id"`
	Name              string        `json:"name"`
	Role              string        `json:"role"`
	Location          string        `json:"location"`
	YearsOfExperience int           `json:"yearsOfExperience"`
	IC                bool          `json:"individualContributor"`
	EM                bool          `json:"managerial"`
	Exec              bool          `json:"executive"`
}

var (
	ErrRecordInvalid = errors.New("record is invalid")
	dbPath           = "data/users.db"
)

func SetDBPath(path string) {
	dbPath = path
}

// Read all user
func All() ([]User, error) {
	db, err := storm.Open(dbPath)
	if err != nil {
		log.Println("Unable to open the database")
		return nil, err
	}
	defer db.Close()
	users := []User{}
	if err := db.All(&users); err != nil {
		log.Println("Unable to read from the database")
		return nil, err
	}
	return users, nil
}

func One(id bson.ObjectId) (*User, error) {
	db, err := storm.Open(dbPath)
	if err != nil {
		log.Printf("Unable to open the database %s", dbPath)
		return nil, err
	}
	defer db.Close()
	u := new(User)
	if err := db.One("ID", id, u); err != nil {
		log.Printf("Unable to read from the database %s", dbPath)
		return nil, err
	}
	return u, nil
}

// Delete the record
func Delete(id bson.ObjectId) error {
	db, err := storm.Open(dbPath)
	if err != nil {
		log.Printf("Unable to open the database %s", dbPath)
		return err
	}
	defer db.Close()
	u := new(User)
	if err := db.One("ID", id, u); err != nil {
		log.Printf("Unable to read from the database %s", dbPath)
		return err
	}
	return db.DeleteStruct(u)
}

// Save the record
func (u *User) Save() error {
	if err := u.Validate(); err != nil {
		return err
	}
	db, err := storm.Open(dbPath)
	if err != nil {
		log.Printf("Unable to open the database %s", dbPath)
		return err
	}
	defer db.Close()
	return db.Save(u)
}

// Validate the entry
func (u *User) Validate() error {
	if u.Name == "" || u.Role == "" {
		return ErrRecordInvalid
	}
	return nil
}
