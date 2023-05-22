package introduction

import (
	"errors"
	"fmt"
)

func Errors() {
	users := []user{{"张三", "123"}, {"李四", "123"}}

	u, err := findUser(users, "张三")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(u.name)
	if u, err = findUser(users, "王五"); err != nil {
		if err != nil {
			fmt.Println(err)
			return
		} else {
			fmt.Println(u.name)
		}
	}
}

// 错误处理
type user struct {
	name     string
	password string
}

func findUser(users []user, name string) (v *user, err error) {
	for _, user := range users {
		if user.name == name {
			return &user, nil
		}
	}
	return nil, errors.New("not found user")
}
