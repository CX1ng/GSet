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
func (st GSet) Insert(data interface{}) error {
	if reflect.TypeOf(data) != st.setType {
		return ErrTypeError
	}

	st.gSet[data] = true
	return nil
}

//删除
func (st GSet) Del(data interface{}) error {
	if reflect.TypeOf(data) != st.setType {
		return ErrTypeError
	}

	delete(st.gSet, data)
	return nil
}

//random key

//bulk insert

//remove

//
