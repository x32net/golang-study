package main

import (
	"errors"
	"fmt"
)

type user struct {
	name string
	age  int64
}

func ValidateUserAge(age int64) (int64, error) {
	if age < 18 {
		return age, errors.New("User is not over 18 years old")
	}
	return age, nil
}

func main() {
	boss := user{name: "Zhang"}
	// syntax error: field appers on left side of :=
	//boss.age, err := ValidateUserAge(22)
	age, err := ValidateUserAge(22)
	if err != nil {
		fmt.Println(err)
	}

	boss.age = age
}
