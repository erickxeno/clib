// 对logs的再次包装
package xlog

import (
	"context"
	"net/http"
	"strconv"
	"sync"

	"github.com/erickxeno/clib/logid"
	"github.com/erickxeno/clib/logs"
)

const (
	defaultReqidKey = "X-Reqid"
	defaultLogKey   = "X-Log"
	maxErrorLen     = 32
)

var (
	reqidKey = defaultReqidKey
	logKey   = defaultLogKey
)

func SetReqidKey(key string) {
	reqidKey = key
}

func SetLogKey(key string) {
	logKey = key
}

type key int

const (
	xlogKey key = 0
)

func NewContext(ctx context.Context, xl *Logger) context.Context {
	if xl == nil {
		panic("nil logger")
	}
	var newCtx = ctx
	newCtx = context.WithValue(newCtx, xlogKey, xl)
	newCtx = context.WithValue(newCtx, reqidKey, xl.ReqId())
	return newCtx
}

type Logger struct {
	ctx   context.Context
	reqId string

	lock *sync.RWMutex // protect header
	h    http.Header   // header for response
}

func New(w http.ResponseWriter, req *http.Request) *Logger {
	reqId := req.Header.Get(reqidKey)
	if reqId == "" {
		//reqId = genReqId()
		reqId = strconv.FormatUint(logid.GetID(), 10)
		req.Header.Set(logs.LogIDCtxKey, reqId)
	}
	h := w.Header() // log X-Log in ResponseWriter header
	h.Set(reqidKey, reqId)

	l := &Logger{
		reqId: reqId,
		lock:  new(sync.RWMutex),
		h:     h,
	}
	l.WithContext(req.Context())
	return l
}

func (xlog *Logger) ReqId() string {
	return xlog.reqId
}

/*
WithContext binding a context with Logger
without new logger created
*/
func (xlog *Logger) WithContext(ctx context.Context) {
	if ctx == nil {
		panic("nil context")
	}
	xlog.ctx = NewContext(ctx, xlog)
}
