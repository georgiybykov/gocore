package aquarium

import (
	"io"
)

type User interface {
	Age() int
}

type Employee struct {
	age int
}

func (e Employee) Age() int {
	return e.age
}

type Customer struct {
	age int
}

func (c Customer) Age() int {
	return c.age
}

func MaxAge(users ...User) (age int) {
	for _, user := range users {
		if user.Age() > age {
			age = user.Age()
		}
	}
	return age
}

func OldestUser(users ...any) (user any) {
	var maxAge int

	for _, u := range users {
		if e, ok := u.(Employee); ok {
			if e.age > maxAge {
				user, maxAge = e, e.age
			}
		}

		if c, ok := u.(Customer); ok {
			if c.age > maxAge {
				user, maxAge = c, c.age
			}
		}
	}
	return user
}

func Print(w io.Writer, args ...any) {
	for _, arg := range args {
		if s, ok := arg.(string); ok {
			w.Write([]byte(s))
		}
	}
}
