package gset

import (
	"reflect"
)

type GSet struct {
	setType reflect.Type
	maxNum  int64
	gSet    map[interface{}]bool
}

//初始化GSet
func NewGSet(gsetType interface{}) (*GSet, error) {
	var gt GSet
	gt.setType = reflect.TypeOf(gsetType)
	gt.gSet = make(map[interface{}]bool)
	gt.gSet[gsetType] = true

	return &gt, nil
}

//获取数量
func (st GSet) Size() int {
	return len(st.gSet)
}

//获取当前类型
func (st GSet) Type() string {
	return st.setType.String()
}

//插入
func (st GSet) Insert(data interface{})(int, error){
	if reflect.TypeOf(data) != st.setType {
		return len(st.gSet), ErrTypeError
	}

	st.gSet[data] = true
	return len(st.gSet),nil
}

//删除
func (st GSet) Del(data interface{})(int,error){
	if reflect.TypeOf(data) != st.setType {
		return len(st.gSet),ErrTypeError
	}

	delete(st.gSet, data)
	return len(st.gSet), nil
}

//插入多个
//请确保参数类型与集合相同，类型不同的项不会执行插入操作
//return: 执行插入的元素个数
func (st GSet) MultiInsert(data ...interface{})(int){
	var cnt int
	for _, item := range data{
		if reflect.TypeOf(item) == st.setType {
			st.gSet[item] = true
			cnt++
		}
	}

	return cnt
}

//删除多个
//请确保参数类型与集合相同，类型不同的项不会执行删除操作
//return: 执行删除的元素个数
func (st GSet) MultiDel(data ...interface{})(int){
	var cnt int
	for _, item := range data{
		if reflect.TypeOf(item) == st.setType {
			delete(st.gSet,item)
			cnt++
		}
	}
	return cnt
}

//
