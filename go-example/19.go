package main

import (
	"fmt"
	"reflect"
)

func main() {

	user := User{
		name: "like",
		age:  22,
	}

	refTag(user, 0)
	refTag(user, 1)
}

// User is custom struct
type User struct {
	name string "this is user name"
	age  int    "this is user age"
}

// refTag can get your struct tag
func refTag(tt User, ix int) {
	ttType := reflect.TypeOf(tt)
	ixField := ttType.Field(ix)
	fmt.Printf("%v\n", ixField.Tag)
}
