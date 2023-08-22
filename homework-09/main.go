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

// 	result := MaxAge(employee1, employee2, employee3, customer1, customer2, customer3)

// 	fmt.Printf("max age: %v", result)
// }

func MaxAge(users ...User) (age int) {
	for _, user := range users {
		if age < user.Age() {
			age = user.Age()
		}
	}
	return age
}
