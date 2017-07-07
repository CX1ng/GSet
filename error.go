package gset

import (
	"errors"
)

var (
	ErrTypeError    = errors.New("类型错误")
	ErrInitTypeError = errors.New("初始化类型不支持")
	ErrSetTypeError = errors.New("集合类型错误")
)


