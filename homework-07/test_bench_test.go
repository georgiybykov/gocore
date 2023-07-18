package testbench

import (
	"fmt"
	"math/rand"
	"reflect"
	"sort"
	"testing"
)

func TestSort_Ints(t *testing.T) {
	got := []int{4134134, 4, 1, -54323, 2, -4, 75, 101, 0}
	want := []int{-54323, -4, 0, 1, 2, 4, 75, 101, 4134134}

	if sort.Ints(got); !reflect.DeepEqual(got, want) {
		t.Errorf("Ints() = %v, want %v", got, want)
	}
}

func TestSort_Strings(t *testing.T) {
	type args struct {
		x []string
	}

	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "Test case #1",
			args: args{
				x: []string{"b", "c", "a"},
			},
			want: []string{"a", "b", "c"},
		},
		{
			name: "Test case #2",
			args: args{
				x: []string{"242523", "%#%^#^#", "", "abcde"},
			},
			want: []string{"", "%#%^#^#", "242523", "abcde"},
		},
		{
			name: "Test case #3",
			args: args{
				x: []string{""},
			},
			want: []string{""},
		},
		{
			name: "Test case #4",
			args: args{
				x: []string{},
			},
			want: []string{},
		},
		{
			name: "Test case #5",
			args: args{
				x: nil,
			},
			want: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.args.x
			if sort.Strings(got); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Strings() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkSort_Ints(b *testing.B) {
	tables := []struct {
		volume int
	}{
		{volume: 10_000},
		{volume: 100_000},
		{volume: 1_000_000},
	}

	for _, tt := range tables {
		b.Run(fmt.Sprintf("%d", tt.volume), func(b *testing.B) {
			s := genIntSlice(tt.volume)
			for i := 0; i < b.N; i++ {
				sort.Ints(s)
			}
		})
	}
}

func BenchmarkSort_Float64s(b *testing.B) {
	tables := []struct {
		volume int
	}{
		{volume: 10_000},
		{volume: 100_000},
		{volume: 1_000_000},
	}

	for _, tt := range tables {
		b.Run(fmt.Sprintf("%d", tt.volume), func(b *testing.B) {
			s := genFloatSlice(tt.volume)
			for i := 0; i < b.N; i++ {
				sort.Float64s(s)
			}
		})
	}
}

func genIntSlice(n int) []int {
	s := make([]int, 0, n)

	for i := 0; i < n; i++ {
		s = append(s, rand.Int())
	}
	return s
}

func genFloatSlice(n int) []float64 {
	s := make([]float64, 0, n)

	for i := 0; i < n; i++ {
		s = append(s, rand.Float64())
	}
	return s
}

/*
goos: linux
goarch: amd64
pkg: gocore/homework-07
cpu: Intel(R) Core(TM) i7-8550U CPU @ 1.80GHz
BenchmarkSort_Ints/10000-8 	           42200	     28089 ns/op	      25 B/op	       1 allocs/op
BenchmarkSort_Ints/100000-8         	    4335	    282011 ns/op	     209 B/op	       1 allocs/op
BenchmarkSort_Ints/1000000-8        	     302	   3372987 ns/op	   26525 B/op	       1 allocs/op
BenchmarkSort_Float64s/10000-8      	   37861	     30712 ns/op	      26 B/op	       1 allocs/op
BenchmarkSort_Float64s/100000-8     	    3817	    302611 ns/op	     234 B/op	       1 allocs/op
BenchmarkSort_Float64s/1000000-8    	     266	   3790915 ns/op	   30112 B/op	       1 allocs/op
PASS
ok  	gocore/homework-07	11.482s
*/
