package main

import (
	"fmt"
)

type notifier interface {
	notify()
	changeEmail(string)
}

type user struct {
	name  string
	email string
}

func (u user) notify() {
	fmt.Println(u.name, u.email)
}

func (u *user) changeEmail(email string) {
	u.email = email
}

// // 继承struct
// type admin struct {
// 	user
// 	level string
// }

// 继承struct指针
type admin struct {
	*user
	level string
}

// // 继承接口
// type admin struct {
// 	notifier
// 	level string
// }

func main() {
	tom := admin{
		user: &user{
			name:  "tom",
			email: "tome@email.com",
		},
		level: "super",
	}
	tom.notify()

	tom.changeEmail("tom@newdomain.com")
	tom.notify()

	fmt.Println(tom.name, tom.email)
}
