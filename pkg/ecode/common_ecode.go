package ecode

var OK = add(1, "Ok")
var ServerErr = add(500, "请求异常")           // 服务器错误
var RequestErr = add(400, "错误的请求")         // 请求错误
var ExceptionRequestErr = add(401, "异常请求") // 异常请求
var RepeatRequestErr = add(402, "重复的请求")   // 重复的请求
var RequestFastErr = add(403, "操作频繁")      // 操作频繁
