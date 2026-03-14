package handlers

import (
	"bytes"
	"encoding/json"
	"goHttpRestAPI/user"
	"io/ioutil"
	"net/http"
	"reflect"
	"testing"

	"gopkg.in/mgo.v2/bson"
)

func TestBodyToUser(t *testing.T) {
	valid := &user.User{
		ID:   bson.NewObjectId(),
		Name: "Maximum",
		Role: "Head of Army",
		Exec: true,
	}
	js, err := json.Marshal(valid)
	if err != nil {
		t.Errorf("Unexpected error while marshalling : %s", err)
		t.FailNow()
	}
	ts := []struct {
		text string
		r    *http.Request
		u    *user.User
		err  bool
		exp  *user.User
	}{
		{
			text: "nil request",
			err:  true,
			exp:  nil,
		},
		{
			text: "nil request body",
			r: &http.Request{
				Body: nil,
			},
			err: true,
		},
		{
			text: "empty user",
			r: &http.Request{
				//Body: ioutil.NopCloser(bytes.NewReader([]byte("{}"))),
				Body: ioutil.NopCloser(bytes.NewBufferString("{}")),
			},
			//u:   &user.User{},
			err: true,
		},
		{
			text: "malformed Data",
			r: &http.Request{
				Body: ioutil.NopCloser(bytes.NewBufferString(`{"id": 12}`)),
			},
			u:   &user.User{},
			err: true,
		},
		{
			text: "Valid Request",
			r: &http.Request{
				Body: ioutil.NopCloser(bytes.NewBuffer(js)),
			},
			u:   &user.User{},
			exp: valid,
			err: false,
		},
	}
	for _, tc := range ts {
		t.Logf("Testing %s\n", tc.text)
		err := bodyToUser(tc.r, tc.u)
		if tc.err {
			if err == nil {
				t.Errorf("Expected error but got nil")
			}
			continue
		}
		if err != nil {
			t.Errorf("UnExpected error :  %s", err)
			continue
		}
		if !reflect.DeepEqual(tc.exp, tc.u) {
			t.Errorf("Expected %v but got %v", tc.exp, tc.u)
		}
	}

}
