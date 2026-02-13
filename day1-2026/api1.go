package main

import (
	"fmt"
	"io"
	"net/http"
)

func dummyRequest(endpoint string) {
	url := fmt.Sprintf("https://fake-json-api.mock.beeceptor.com/%s", endpoint)
	//use to format any string
	res1, err := http.Get(url)
	if err != nil {
		fmt.Printf("Error %w \n", err)
		return
	}
	defer res1.Body.Close() //Close res body
	body, _ := io.ReadAll(res1.Body)
	fmt.Println(string(body))
}

func main() {
	fmt.Println("Querying the user endpoint")
	dummyRequest("users")
	dummyRequest("companies")
}
