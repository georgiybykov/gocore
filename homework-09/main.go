package main

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

// func main() {
// 	fmt.Println("a")

// 	employee1 := Employee{age: 30}
// 	employee2 := Employee{age: 13}
// 	employee3 := Employee{age: 23}

// 	customer1 := Customer{age: 15}
// 	customer2 := Customer{age: 54}
// 	customer3 := Customer{age: 34}

// 	result1 := MaxAge(employee1, employee2, employee3, customer1, customer2, customer3)

// 	fmt.Printf("max age: %v \n", result1)

// 	result2 := OldestUser(employee1, employee2, employee3, customer1, customer2, customer3)

// 	fmt.Printf("max age: %v \n type: %T", result2, result2)
// }

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
