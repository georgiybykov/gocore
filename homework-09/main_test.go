package main

import "testing"

func TestMaxAge(t *testing.T) {
	tests := []struct {
		name  string
		users []User
		want  int
	}{
		{
			name: "When the users are employees",
			users: []User{
				Employee{age: 21},
				Employee{age: 64},
				Employee{age: 51},
			},
			want: 64,
		},
		{
			name: "When the users are customers",
			users: []User{
				Customer{age: 31},
				Customer{age: 73},
				Customer{age: 47},
			},
			want: 73,
		},
		{
			name: "When the users are employees and customers",
			users: []User{
				Employee{age: 44},
				Employee{age: 37},
				Customer{age: 55},
				Customer{age: 23},
			},
			want: 55,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MaxAge(tt.users...); got != tt.want {
				t.Errorf("got MaxAge() = %v, want %v", got, tt.want)
			}
		})
	}
}
