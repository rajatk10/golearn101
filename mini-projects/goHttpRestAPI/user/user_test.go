package user

import (
	"os"
	"reflect"
	"strconv"
	"testing"

	"github.com/asdine/storm"
	"gopkg.in/mgo.v2/bson"
)

//Unit tests for CRUD Operations

func cleanUpDB() {
	os.RemoveAll("./data")
	os.Mkdir("./data", 0755)
	os.Create("./data/users_test.db")
	SetDBPath("./data/users_test.db")
}
func TestCRUD(t *testing.T) {
	t.Logf("TestCRUD Operations - Create")
	//Create a test database
	os.RemoveAll("./data")
	//Cleanup in each run
	os.Mkdir("./data", 0755)
	os.Create("./data/users_test.db")
	SetDBPath("./data/users_test.db")
	u := &User{
		ID:   bson.NewObjectId(),
		Name: "Maximus",
		Role: "Head of Army",
		Exec: true,
	}
	u11 := &User{
		ID:   bson.NewObjectId(),
		Name: "Maximus Son",
		Role: "Citizen",
		Exec: false,
	}
	u12 := &User{
		ID:   bson.NewObjectId(),
		Name: "Maximus Wife",
		Role: "Citizen",
		EM:   true,
		Exec: false,
	}
	err := u.Save()
	if err != nil {
		t.Fatalf("Unexpected error while saving : %s", err)
	}
	t.Log("TestCRUD Operations - Read")
	u2, err := One(u.ID)
	if err != nil {
		t.Fatalf("Unexpected error while reading : %s", err)
	}
	//if u.ID != u2.ID {
	//	t.Fatalf("Unexpected ID while reading : %s", err)
	//}
	if !reflect.DeepEqual(u, u2) {
		t.Error("Read Record do not match")
		//t.Fatalf("Unexpected data while reading : %s", err)
	}
	t.Log("TestCRUD Operations - Update")
	u.Role = "Head of Company"
	err = u.Save()
	if err != nil {
		t.Fatalf("Unexpected error while updating : %s", err)
	}
	u3, err := One(u.ID)
	if err != nil {
		t.Fatalf("Unexpected error while reading : %s", err)
	}
	if !reflect.DeepEqual(u, u3) {
		t.Error("Update Failed, Records do not match")
	}
	t.Log("TestCRUD - Delete Operations")
	err = Delete(u.ID)
	if err != nil {
		t.Errorf("Unexpected error while deleting : %s", err)
	}
	_, err = One(u.ID)
	if err == nil {
		t.Error("Delete Failed, Record still exists")
	}
	if err != storm.ErrNotFound {
		t.Errorf("Unexpected error while reading non existent record : %s", err)
	}

	t.Log("TestCRUD - ReadAll")
	err = u12.Save()
	if err != nil {
		t.Error("Unexpected error while saving : ", err)
	}
	err = u11.Save()
	if err != nil {
		t.Error("Unexpected error while saving : ", err)
	}
	usersAll, err := All()
	if len(usersAll) != 2 {
		t.Error("Unexpected number of records")
	}
	if !reflect.DeepEqual(usersAll, []User{*u11, *u12}) {
		t.Error("ReadAll Record do not match")
	}
}

func BenchmarkCRUD(b *testing.B) {
	b.ResetTimer()
	//t.Logf("TestCRUD Operations - Create")
	//Create a test database
	os.RemoveAll("./data")
	//Cleanup in each run
	os.Mkdir("./data", 0755)
	os.Create("./data/users_test.db")
	SetDBPath("./data/users_test.db")
	for i := 0; i < b.N; i++ {
		u := &User{
			ID:   bson.NewObjectId(),
			Name: "Maximus_" + strconv.Itoa(i),
			Role: "Head of Army",
			Exec: true,
		}
		//u11 := &User{
		//	ID:   bson.NewObjectId(),
		//	Name: "Maximus Son",
		//	Role: "Citizen",
		//	Exec: false,
		//}
		//u12 := &User{
		//	ID:   bson.NewObjectId(),
		//	Name: "Maximus Wife",
		//	Role: "Citizen",
		//	EM:   true,
		//	Exec: false,
		//}
		err := u.Save()
		if err != nil {
			b.Fatalf("Unexpected error while saving : %s", err)
		}
		//t.Log("TestCRUD Operations - Read")
		_, err = One(u.ID)
		if err != nil {
			b.Fatalf("Unexpected error while reading : %s", err)
		}
		//if u.ID != u2.ID {
		//	t.Fatalf("Unexpected ID while reading : %s", err)
		//}
		//if !reflect.DeepEqual(u, u2) {
		//	//t.Error("Read Record do not match")
		//	b.Fatalf("Unexpected data while reading : %s", err)
		//}
		//t.Log("TestCRUD Operations - Update")
		u.Role = "Head of Company"
		err = u.Save()
		if err != nil {
			b.Fatalf("Unexpected error while updating : %s", err)
		}
		_, err = One(u.ID)
		if err != nil {
			b.Fatalf("Unexpected error while reading : %s", err)
		}
		//t.Log("TestCRUD - Delete Operations")
		err = Delete(u.ID)
		if err != nil {
			b.Fatalf("Unexpected error while deleting : %s", err)
		}
		//
		////t.Log("TestCRUD - ReadAll")
		//err = u12.Save()
		//if err != nil {
		//	t.Error("Unexpected error while saving : ", err)
		//}
		//err = u11.Save()
		//if err != nil {
		//	t.Error("Unexpected error while saving : ", err)
		//}
		//usersAll, err := All()
		//if len(usersAll) != 2 {
		//	t.Error("Unexpected number of records")
		//}
		//if !reflect.DeepEqual(usersAll, []User{*u11, *u12}) {
		//	t.Error("ReadAll Record do not match")
		//}
	}
}

func BenchmarkCreate(b *testing.B) {
	b.ResetTimer()
	//t.Logf("TestCRUD Operations - Create")
	//Create a test database
	cleanUpDB()
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		u := &User{
			ID:   bson.NewObjectId(),
			Name: "Maximus_" + strconv.Itoa(i),
			Role: "Head of Army",
			Exec: true,
		}
		b.StartTimer()
		err := u.Save()
		if err != nil {
			b.Fatalf("Unexpected error while saving : %s", err)
		}
		b.StopTimer()
		Delete(u.ID)
	}
}

func BenchmarkRead(b *testing.B) {
	b.ResetTimer()
	SetDBPath("./data/users_test.db")
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		u := &User{
			ID:   bson.NewObjectId(),
			Name: "Maximus_" + strconv.Itoa(i),
			Role: "Head of Army",
			Exec: true,
		}

		err := u.Save()
		if err != nil {
			b.Fatalf("Unexpected error while saving : %s", err)
		}
		b.StartTimer()
		_, err = One(u.ID)
		if err != nil {
			b.Fatalf("Unexpected error while reading : %s", err)
		}
	}
}

func BenchmarkDelete(b *testing.B) {
	b.ResetTimer()
	cleanUpDB()
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		u := &User{
			ID:   bson.NewObjectId(),
			Name: "Maximus_" + strconv.Itoa(i),
			Role: "Head of Army",
			Exec: true,
		}
		err := u.Save()
		if err != nil {
			b.Fatalf("Unexpected error while saving : %s", err)
		}
		b.StartTimer()
		err = Delete(u.ID)
		if err != nil {
			b.Fatalf("Unexpected error while deleting : %s", err)
		}

	}
}

func BenchmarkUpdate(b *testing.B) {
	b.ResetTimer()
	cleanUpDB()
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		u := &User{
			ID:   bson.NewObjectId(),
			Name: "Maximus_" + strconv.Itoa(i),
			Role: "Head of Army",
			Exec: true,
		}
		err := u.Save()
		if err != nil {
			b.Fatalf("Unexpected error while saving : %s", err)
		}
		u.Role = "Head of Company_" + strconv.Itoa(i)
		b.StartTimer()
		err = u.Save()
		if err != nil {
			b.Fatalf("Unexpected error while updating : %s", err)
		}
	}
}
