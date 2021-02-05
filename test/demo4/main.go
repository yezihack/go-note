package main

import (
	"errors"
	"fmt"
)

type UserNotFoundErr struct {
}

func (e UserNotFoundErr) Error() string {
	return fmt.Sprint("can not found user")
}

func oops() error {
	return UserNotFoundErr{}
}

func main() {
	if err := oops(); err != nil {

		var another UserNotFoundErr

		fmt.Println(err)                     //can not found user
		fmt.Println(err == another)          //true
		fmt.Println(errors.Is(err, another)) //true

		uid := 51002
		err2 := fmt.Errorf("%w:%v", err, uid) //can not found user:51002
		fmt.Println(err2)
		fmt.Println(err2 == another)          //false <==
		fmt.Println(errors.Is(err2, another)) //true

	}
}
