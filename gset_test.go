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
