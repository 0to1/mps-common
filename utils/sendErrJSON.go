package utils

import (
	"net/http"

	"github.com/kataras/iris/v12"
	"sync"
)

var mu sync.Mutex

// SendErrJSON 有错误发生时，发送错误JSON
func SendErrJSON(msg string, args ...interface{}) {
	if len(args) == 0 {
		panic("missing iris.Context")
	}

	var ctx iris.Context
	var errNo = ErrorCode.ERROR

	if len(args) == 1 {
		theCtx, ok := args[0].(iris.Context)
		if !ok {
			panic("missing iris.Context")
		}
		ctx = theCtx
	} else if len(args) == 2 {
		theErrNo, ok := args[0].(int)
		if !ok {
			panic("errNo不正确")
		}
		errNo = theErrNo
		theCtx, ok := args[1].(iris.Context)
		if !ok {
			panic("missing iris.Context")
		}
		ctx = theCtx
	}

	ctx.StatusCode(http.StatusBadRequest)

	mu.Lock()
	ctx.JSON(iris.Map{
		"errNo": errNo,
		"msg":   msg,
		"data":  struct{}{},
	})
	mu.Unlock()

	// 终止请求链
	// ctx.EndRequest()
}
