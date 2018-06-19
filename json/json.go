package main

import (
	"encoding/json"
	"fmt"
	"time"
)

// User structure
type User struct {
	Name      string    `json:"name"`
	Phone     string    `json:"phone"`
	Age       int       `json:"age"`
	CreatedTs time.Time `json:"createdTs,string"`
}

func main() {
	now := time.Now()
	user1 := &User{Name: "John", Phone: "512-000-000", Age: 35, CreatedTs: now}
	b, err := json.Marshal(user1) // b is an array of bytes
	if err != nil {
		fmt.Println(err)
		return
	}

	// Create string from array of bytes...
	fmt.Println(string(b))

	var user2 User
	error := json.Unmarshal(b, &user2)
	if error != nil {
		fmt.Println(error)
		return
	}

	fmt.Println(user2)

	var jsonStr = `{
		"name": "John",
		"phone": "512-000-000",
		"age": 35,
		"createdTs": "2018-06-19T12:29:42.647435475-04:00"
	  }`

	fmt.Println(jsonStr)

	error = json.Unmarshal([]byte(jsonStr), &user2)
	if error != nil {
		fmt.Println(error)
		return
	}

	fmt.Printf("%+v\n", user2)
}
