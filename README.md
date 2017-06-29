GSet
---
一个简单的Golang集合包

目前支持的操作：
+ 获取集合元素数
+ 增加元素
+ 删除元素
+ 增加多个元素
+ 删除多个元素

初始化
-----
    str := "apple"
    set, err := gset.NewGSet(str)
NewGSet()传入的第一个参数，其类型指定了集合存储的元素类型
