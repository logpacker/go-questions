package main

import (
	"fmt"
	"reflect"
)

func main() {
	type User struct {
		Name string `key:"name" maxlength:"128"`
	}

	u := User{}
	st := reflect.TypeOf(u)
	field := st.Field(0)
	fmt.Println(field.Tag.Get("key"), field.Tag.Get("maxlength"))
}
