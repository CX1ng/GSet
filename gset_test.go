package gset

import (
	"testing"
	"fmt"
)

func TestNewGSet(t *testing.T) {
	testExample := []struct{
		argtype interface{}
		ok error
	}{
		{"test", nil}, //字符串
		{123,nil}, //int
		{123.456, nil}, //float
		{true,nil},//bool
		{make(chan int), nil}, //chan
 		{[]int{1,2,3},ErrInitTypeError}, //slice
		{make(map[int] int),ErrInitTypeError}, //map
		{func(){
			fmt.Println("Error")
		},ErrInitTypeError},
		{[3]int{1,2,3},ErrInitTypeError},

	}

	for _,test := range testExample{
		if _,ok := NewGSet(test.argtype); ok != test.ok{
			t.Errorf("%s argument init failed!", test.argtype)
		}
	}
}
