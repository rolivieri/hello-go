package main

import (
	"fmt"
	"reflect"
)

type User struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Email     string `json:"email"`
}

// The starting point for any Go program is the main function.
func main() {
	var f1 User
	users := myReflect(f1).([]User)
	fmt.Printf("It rocks: %v \n", users)
	// f1.Firstname = "John"
	// fType := reflect.TypeOf(f1)
	// fVal := reflect.New(fType)
	// f2 := fVal.Elem().Interface().(User)
	// f2.Firstname = "Peter"
	// fmt.Println("Hello World: " + f2.Firstname)
	// fmt.Println("Hello World: " + f1.Firstname)
	// inter := fVal.Elem().Interface()
	// myUser := inter.(User)
	// myUser.Firstname = "Phil"
	//fmt.Println("Hello World: " + f1.Firstname)

}

func myReflect(ptr interface{}) interface{} {
	fType := reflect.TypeOf(ptr)
	fVal := reflect.New(fType)
	f2 := fVal.Elem().Interface().(User)
	f2.Firstname = "Peter"
	fmt.Println("Hello World: " + f2.Firstname)
	//fmt.Println("Hello World: " + f1.Firstname)
	inter := fVal.Elem().Interface()
	myUser := inter.(User)
	myUser.Firstname = "Phil"
	fmt.Println("Hello World: " + myUser.Firstname)

	//var users []User

	//users := make([]User, 0)
	//sliceType := reflect.TypeOf(users)
	//var counties []commonLib.County
	//userSliceReflect := reflect.MakeSlice(sliceType, 0, 0)
	userSliceReflect := reflect.MakeSlice(reflect.SliceOf(fType), 0, 0)

	a := reflect.ValueOf(myUser)
	userSliceReflect = reflect.Append(userSliceReflect, a)
	//fmt.Println("Hello World: " + userSliceReflect)
	fmt.Printf("Hello World: %v \n", userSliceReflect)
	t := userSliceReflect.Interface().([]User)
	fmt.Printf("Hello World: %v \n", t)
	return userSliceReflect.Interface()

	//return userSliceReflect
}
