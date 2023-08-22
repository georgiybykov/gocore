package aquarium

import (
	"bytes"
	"reflect"
	"testing"
)

func Test_MaxAge(t *testing.T) {
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

func Test_OldestUser(t *testing.T) {
	tests := []struct {
		name  string
		users []any
		want  any
	}{
		{
			name: "When the users are employees",
			users: []any{
				Employee{age: 21},
				Employee{age: 64},
				Employee{age: 51},
			},
			want: Employee{age: 64},
		},
		{
			name: "When the users are customers",
			users: []any{
				Customer{age: 31},
				Customer{age: 73},
				Customer{age: 47},
			},
			want: Customer{age: 73},
		},
		{
			name: "When the users are employees and customers",
			users: []any{
				Employee{age: 44},
				Employee{age: 37},
				Customer{age: 55},
				Customer{age: 23},
			},
			want: Customer{age: 55},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := OldestUser(tt.users...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("got OldestUser() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_Print(t *testing.T) {
	tests := []struct {
		name string
		args []any
		want string
	}{
		{
			name: "Test case #1",
			args: []any{1, true, "Print ", 3423, 4534.354, "this ", "\n ", nil, "line"},
			want: "Print this \n line",
		},
		{
			name: "Test case #2",
			args: []any{1, true, 3423, 4534.354, nil},
			want: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := &bytes.Buffer{}
			Print(w, tt.args...)

			if got := w.String(); got != tt.want {
				t.Errorf("got Print() = %v, want %v", got, tt.want)
			}
		})
	}
}
