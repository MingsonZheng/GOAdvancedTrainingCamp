package main

import (
	"net/http"
	"sync"
)

type handlerFunc func(ctx *Context)

type Routable interface {
	RouteV4(method string, pattern string, handlerFunc handlerFunc) // server 可以把 Route 委托给这边的 Handler
}

type Handler interface {
	//ServeHTTP http.Handler
	ServeHTTP(c *Context)
	//Routable Route(method string, pattern string, handlerFunc func(ctx *Context)) // server 可以把 Route 委托给这边的 Handler
	Routable
}

type HandlerBasedOnMap struct {
	// key 应该是 method + url
	//handlers map[string]func(ctx *Context)
	handlers sync.Map
}

// RouteV4 注册路由
func (h *HandlerBasedOnMap) RouteV4(method string, pattern string, handlerFunc handlerFunc) {
	key := h.key(method, pattern)
	//h.handlers[key] = handlerFunc
	h.handlers.Store(key, handlerFunc)
}

//// RouteV3 注册路由
//func (s *sdkHttpServerV3) RouteV3(method string, pattern string, handlerFunc func(ctx *Context)) {
//	key := s.handler.key(method, pattern)
//	s.handler.handlers[key] = handlerFunc
//}

func (h *HandlerBasedOnMap) ServeHTTP(c *Context) {
	//key := h.key(c.R.Method, c.R.URL.Path)
	//// 判定路由是否已经注册
	//if handler, ok := h.handlers[key]; ok {
	//	handler(c)
	//} else {
	//	c.W.WriteHeader(http.StatusNotFound)
	//	c.W.Write([]byte("Not Found"))
	//}
	request := c.R
	key := h.key(request.Method, request.URL.Path)
	handler, ok := h.handlers.Load(key)
	if !ok {
		c.W.WriteHeader(http.StatusNotFound)
		c.W.Write([]byte("Not Found"))
		return
	}
	handler.(func(c *Context))(c)
}

//func (h *HandlerBasedOnMap) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
//	key := h.key(request.Method, request.URL.Path)
//	// 判定路由是否已经注册
//	if handler, ok := h.handlers[key]; ok {
//		handler(NewContext(writer, request))
//	} else {
//		writer.WriteHeader(http.StatusNotFound)
//		writer.Write([]byte("Not Found"))
//	}
//}

func (h *HandlerBasedOnMap) key(method string, pattern string) string {
	return method + "#" + pattern
}

func NewHandlerBasedOnMap() Handler {
	return &HandlerBasedOnMap{
		//handlers: make(map[string]func(ctx *Context)),
		handlers: sync.Map{},
	}
}
