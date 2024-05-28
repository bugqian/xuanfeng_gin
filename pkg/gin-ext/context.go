package ginext

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
	"sync"
	"xuanfeng_gin/pkg/ecode"
)

var contextPool = sync.Pool{New: func() interface{} { return &Context{} }}

func H(handler func(*Context)) gin.HandlerFunc {
	return func(original *gin.Context) {
		ctx := contextPool.Get().(*Context)
		ctx.Context = original

		handler(ctx)

		ctx.Context = nil
		contextPool.Put(ctx)
	}
}

// 自定义上下文
type Context struct {
	*gin.Context
}

type RES struct {
	Code    int         `json:"code"`
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
}

type BodyRES struct {
	RES
	Sign      string `json:"sign"`
	Timestamp int64  `json:"timestamp"`
}

// JSON .
func (ctx *Context) JSON(data interface{}, err error) {
	code := ecode.Cause(err)

	if !code.Equal(ecode.OK) {
		// 如果有错误，不返回数据
		data = nil
	}
	var message string
	message = code.Message()
	// 返回数据
	res := RES{
		Data:    data,
		Code:    code.Code(),
		Message: message,
		//Message: gin18n.MustGetMessage(code.Message()),
	}
	ctx.Context.JSON(http.StatusOK, res)

	// 记录返回的数据
	ctx.Context.Set("JSONRES", res)
}

func (ctx *Context) ClientIP() string {
	return clientIP(ctx.Context)
}

func clientIP(ctx *gin.Context) string {
	// 先获取slb头
	addr := ctx.GetHeader("X-Forwarded-For")
	if addr != "" {
		addr = strings.Split(addr, ",")[0]
		if addr != "" {
			return addr
		}
	}

	// 获取不到再获取默认的
	return ctx.ClientIP()
}

// TODO .
func TODO(ctx *Context) {
	ctx.JSON("TODO", nil)
}

// JSON .
func (ctx *Context) JSON3(data interface{}, err error) {
	code := ecode.Cause(err)

	if !code.Equal(ecode.OK) {
		// 如果有错误，不返回数据
		data = data
	}
	var message string
	message = code.Message()
	// 返回数据
	res := RES{
		Data:    data,
		Code:    code.Code(),
		Message: message,
		//Message: gin18n.MustGetMessage(code.Message()),
	}
	ctx.Context.JSON(http.StatusOK, res)

	// 记录返回的数据
	ctx.Context.Set("JSONRES", res)
}
