package main

import (
	"fmt"
	. "gset"
	"os"
)

func main() {
	str := "apple"

	set, err := NewGSet(str) //创建集合，且接受类型为string的数据
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	set.Add("peach")
	_, err = set.Add(1)

	fmt.Println(err) //1因为不是string，所以插入失败
	fmt.Println("Size:", set.Size())

	set.Remove("peach")
	fmt.Println("Size:", set.Size())

	fmt.Println("Type:", set.Type())

	set.MultiAdd("orange", "banana", 1, "lemon")
	fmt.Println("Size:", set.Size())

	set.MultiRemove("watermelon", 2, "banana")
	fmt.Println("Size:", set.Size())

	//set.Clear()
	//fmt.Println("Size:", set.Size())

	set2, err := NewGSet("Golang")

	set2.MultiAdd("C", "Java", "C++")
	fmt.Println("Set2 Size:", set2.Size())

	setUnion, _ := set.Union(*set2)
	fmt.Println("Set Union Size:", setUnion.Size())

	setIntersect, _ := set.Intersect(*set2)
	fmt.Println("Set Intersect Size:", setIntersect.Size())

	setExcept, _ := set.Except(*set2)
	fmt.Println("Set Except Size:", setExcept.Size())

	fmt.Println("Before bulk add slice:", set.Size())
	testSlice := []string{"Python", "R", "Ruby"}
	_, err = set.BulkAdd(testSlice)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("After bulk add slice:", set.Size())

	testArray := [3]string{"Scala", "Lua", "Rust"}
	_, err = set.BulkAdd(testArray)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("After bulk add array:", set.Size())

	_, err = set.BulkRemove(testArray)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("After bulk remove array:", set.Size())
	_, err = set.BulkRemove(testSlice)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("After bulk remove slice:", set.Size())

}
