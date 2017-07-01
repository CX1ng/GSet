package gset

import (
	"errors"
)

var (
	ErrTypeError    = errors.New("类型错误")
	ErrSetTypeError = errors.New("集合类型不一致")
)
