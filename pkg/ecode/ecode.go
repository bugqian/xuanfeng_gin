package ecode

import (
	"fmt"

	"github.com/pkg/errors"
)

var _codes = make(map[int]struct{}) // register codes.

func add(c int, m string) Code {
	if _, ok := _codes[c]; ok {
		panic(fmt.Sprintf("ecode: %d 已经存在", c))
	}
	_codes[c] = struct{}{}
	return Code{
		code:    c,
		message: m,
	}
}

// New .
func New(c int, m string) Code {
	if c < 1000 {
		panic("非公共的状态码不能小于1000")
	}
	return add(c, m)
}

func NewCustom(c int, m string) Code {
	if c < 1000 {
		panic("非公共的状态码不能小于1000")
	}

	return Code{
		code:       c,
		message:    m,
		noNeedI18n: true,
	}
}

// Code .
type Code struct {
	code       int
	message    string
	noNeedI18n bool
}

// Error 实现 error interface
func (e Code) Error() string { return e.message }

// Code 返回错误代码
func (e Code) Code() int { return e.code }

// Code 返回错误消息
func (e Code) Message() string { return e.message }

// Equal 是否等于某个code
func (e Code) Equal(c Code) bool { return e == c }

// Cause .
func Cause(err error) Code {
	if err == nil {
		return OK
	}

	if ec, ok := errors.Cause(err).(Code); ok {
		return ec
	}

	return ServerErr
}

func (e Code) NoNeedI18n() bool {
	return e.noNeedI18n
}
