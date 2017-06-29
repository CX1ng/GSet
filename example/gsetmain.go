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

	set.Insert("peach")
	err = set.Insert(1)

	fmt.Println(err) //1因为不是string，所以插入失败
	fmt.Println(set.Size())

	set.Del("peach")
	fmt.Println(set.Size())

	fmt.Println(set.Type())
}
