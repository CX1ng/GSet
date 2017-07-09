package gset

import (
	"reflect"
	"sync"
)

type GSet struct {
	setType reflect.Type
	maxNum  int64
	gSet    map[interface{}]bool
	lock    *sync.RWMutex
}

var typeTest = []interface{}{
	reflect.Func,
	reflect.Array,
	reflect.Slice,
	reflect.Map,
	reflect.Interface,
}

//初始化GSet
func NewGSet(gsetType interface{}) (*GSet, error) {
	var gt GSet

	gt.setType = reflect.TypeOf(gsetType)
	for _, tp := range typeTest {
		if gt.setType.Kind() == tp {
			return nil, ErrInitTypeError
		}
	}
	gt.gSet = make(map[interface{}]bool)
	gt.gSet[gsetType] = true
	gt.lock = new(sync.RWMutex)

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
func (st GSet) Add(data interface{}) (int, error) {
	if reflect.TypeOf(data) != st.setType {
		return len(st.gSet), ErrTypeError
	}

	st.lock.Lock()
	defer st.lock.Unlock()
	st.gSet[data] = true

	return len(st.gSet), nil
}

//删除
func (st GSet) Remove(data interface{}) (int, error) {
	if reflect.TypeOf(data) != st.setType {
		return len(st.gSet), ErrTypeError
	}

	st.lock.Lock()
	defer st.lock.Unlock()
	delete(st.gSet, data)

	return len(st.gSet), nil
}

//插入多个
//请确保参数类型与集合相同，类型不同的项不会执行插入操作
//return: 执行插入的元素个数
func (st GSet) MultiAdd(data ...interface{}) int {
	var cnt int
	for _, item := range data {
		if reflect.TypeOf(item) == st.setType {
			st.lock.Lock()
			st.gSet[item] = true
			st.lock.Unlock()
			cnt++
		}
	}

	return cnt
}

//删除多个
//请确保参数类型与集合相同，类型不同的项不会执行删除操作
//return: 执行删除的元素个数
func (st GSet) MultiRemove(data ...interface{}) int {
	var cnt int
	for _, item := range data {
		if reflect.TypeOf(item) == st.setType {
			st.lock.Lock()
			delete(st.gSet, item)
			st.lock.Unlock()
			cnt++
		}
	}
	return cnt
}

//清空集合
func (st GSet) Clear() {
	//st.gSet = make(map[interface{}] bool) 这种操作不行??
	st.lock.Lock()
	defer st.lock.Unlock()

	for item := range st.gSet {
		delete(st.gSet, item)
		//fmt.Println("delete", item)
	}
}

//集合之间的并
func (st GSet) Union(other GSet) (*GSet, error) {
	if other.setType != st.setType {
		return nil, ErrSetTypeError
	}

	result := newGSetParamRType(st.setType)

	other.lock.RLock()
	defer other.lock.RUnlock()
	for key := range other.gSet {
		result.gSet[key] = true
	}

	st.lock.RLock()
	defer st.lock.RUnlock()
	for key := range st.gSet {
		result.gSet[key] = true
	}

	return &result, nil
}

//集合之间的交
func (st GSet) Intersect(other GSet) (*GSet, error) {
	if other.setType != st.setType {
		return nil, ErrSetTypeError
	}

	result := newGSetParamRType(st.setType)

	st.lock.RLock()
	defer st.lock.RUnlock()
	other.lock.RLock()
	defer other.lock.RUnlock()

	for key := range other.gSet {
		if _, ok := st.gSet[key]; ok {
			result.gSet[key] = true
		}
	}

	return &result, nil
}

//集合之间的差
func (st GSet) Except(other GSet) (*GSet, error) {
	if other.setType != st.setType {
		return nil, ErrSetTypeError
	}

	result := newGSetParamRType(st.setType)

	st.lock.RLock()
	defer st.lock.RUnlock()
	other.lock.RLock()
	defer other.lock.RUnlock()
	for key := range st.gSet {
		if _, ok := other.gSet[key]; !ok {
			result.gSet[key] = true
		}
	}

	return &result, nil
}

//新建GSet,接受reflect.type参数
func newGSetParamRType(setType reflect.Type) GSet {
	gset := GSet{
		setType: setType,
		gSet:    make(map[interface{}]bool),
		lock:    new(sync.RWMutex),
	}
	return gset
}

//查询集合中是否有指定值
func (st GSet) Exists(data interface{}) (bool, error) {
	if reflect.TypeOf(data) != st.setType {
		return false, ErrTypeError
	}

	st.lock.RLock()
	defer st.lock.RUnlock()
	_, ok := st.gSet[data]

	return ok, nil
}

//合并切片、数组
func (st GSet) BulkAdd(data interface{}) (int, error) {
	if reflect.TypeOf(data).Kind() != reflect.Slice && reflect.TypeOf(data).Kind() != reflect.Array {
		return 0, ErrArgTypeError
	}

	if reflect.TypeOf(data).Elem() != st.setType {
		return 0, ErrTypeError
	}

	s := reflect.ValueOf(data)

	st.lock.Lock()
	defer st.lock.Unlock()
	for i := 0; i < s.Len(); i++ {
		st.gSet[s.Index(i)] = true
	}
	return s.Len(), nil
}

//切片、数组删除
func (st GSet) BulkRemove(data interface{}) (int, error) {
	if reflect.TypeOf(data).Kind() != reflect.Slice && reflect.TypeOf(data).Kind() != reflect.Array {
		return 0, ErrArgTypeError
	}

	if reflect.TypeOf(data).Elem() != st.setType {
		return 0, ErrTypeError
	}

	s := reflect.ValueOf(data)

	st.lock.Lock()
	defer st.lock.Unlock()
	for i := 0; i < s.Len(); i++ {
		delete(st.gSet, s.Index(i))
	}
	return s.Len(), nil
}
