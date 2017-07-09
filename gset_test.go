package gset

import (
	"fmt"
	"testing"
)

func TestNewGSet(t *testing.T) {
	testExample := []struct {
		argtype interface{}
		ok      error
	}{
		{"test", nil},                         //字符串
		{123, nil},                            //int
		{123.456, nil},                        //float
		{true, nil},                           //bool
		{make(chan int), nil},                 //chan
		{[]int{1, 2, 3}, ErrInitTypeError},    //slice
		{make(map[int]int), ErrInitTypeError}, //map
		{func() {
			fmt.Println("Error")
		}, ErrInitTypeError},
		{[3]int{1, 2, 3}, ErrInitTypeError},
	}

	for _, test := range testExample {
		if _, ok := NewGSet(test.argtype); ok != test.ok {
			t.Errorf("%s argument init failed!", test.argtype)
		}
	}
}

func TestGSet_Add(t *testing.T) {
	str := "test"
	set, err := NewGSet(str)
	if err != nil {
		fmt.Println(err)
	}

	testExample := []struct {
		data interface{}
		ok   error
	}{
		{"test", nil},                     //字符串
		{123, ErrTypeError},               //int
		{123.456, ErrTypeError},           //float
		{true, ErrTypeError},              //bool
		{make(chan int), ErrTypeError},    //chan
		{[]int{1, 2, 3}, ErrTypeError},    //slice
		{make(map[int]int), ErrTypeError}, //map
		{func() {
			fmt.Println("Error")
		}, ErrTypeError},
		{[3]int{1, 2, 3}, ErrTypeError},
		{"golang", nil},
		{"c++", nil},
	}
	for _, test := range testExample {
		_, ok := set.Add(test.data)
		if ok != test.ok {
			t.Errorf("GSet error, when add value %v\n", test.data)
		}
	}
}

func TestGSet_Remove(t *testing.T) {
	str := "test"
	set, err := NewGSet(str)
	if err != nil {
		fmt.Println(err)
	}

	testExample := []struct {
		data interface{}
		ok   error
	}{
		{"test", nil},                     //字符串
		{123, ErrTypeError},               //int
		{123.456, ErrTypeError},           //float
		{true, ErrTypeError},              //bool
		{make(chan int), ErrTypeError},    //chan
		{[]int{1, 2, 3}, ErrTypeError},    //slice
		{make(map[int]int), ErrTypeError}, //map
		{func() {
			fmt.Println("Error")
		}, ErrTypeError},
		{[3]int{1, 2, 3}, ErrTypeError},
		{"golang", nil},
		{"c++", nil},
	}
	for _, test := range testExample {
		_, ok := set.Remove(test.data)
		if ok != test.ok {
			t.Errorf("GSet error, when remove value %v\n", test.data)
		}
	}
}

func TestGSet_Clear(t *testing.T) {
	str := "test"
	set, err := NewGSet(str)
	if err != nil {
		fmt.Println(err)
	}
	set.Clear()
	if set.Size() != 0 {
		t.Errorf("GSet error, when clear gset\n")
	}
}

func TestGSet_Size(t *testing.T) {
	str := "test"
	set, err := NewGSet(str)
	if err != nil {
		fmt.Println(err)
	}
	testExample := []struct {
		value interface{}
		cnt   int
	}{
		{"test", 1},            //字符串
		{123, 1},               //int
		{123.456, 1},           //float
		{true, 1},              //bool
		{make(chan int), 1},    //chan
		{[]int{1, 2, 3}, 1},    //slice
		{make(map[int]int), 1}, //map
		{func() {
			fmt.Println("Error")
		}, 1},
		{[3]int{1, 2, 3}, 1},
		{"golang", 2},
		{"c++", 3},
	}
	for _, test := range testExample {
		set.Add(test.value)
		count := set.Size()
		if count != test.cnt {
			t.Errorf("GSet Error, When get size of gset")
		}
	}
}
