package main

import (
	"encoding/json"
	"fmt"
)

// User structure
type User struct {
	Name  string `json:"name"`
	Phone string `json:"phone"`
	Age   int    `json:"age"`
}

func main() {
	user1 := &User{Name: "John", Phone: "512-000-000", Age: 35}
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
}
