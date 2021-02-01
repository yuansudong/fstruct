package fstruct

import "errors"

var (
	// ErrIOIsNil 对象是一个空值
	ErrIOIsNil error = errors.New("object is nil")
	// ErrNoPtr 对象不是一个指针
	ErrNoPtr error = errors.New("object is not ptr")
	// ErrNoStruct 对象不是一个指针
	ErrNoStruct error = errors.New("object is not struct")
)
