package main

import (
	"fmt"
	"net/http"
)

type ServerV3 interface {

	// RouteV3 设定一个路由，命中该路由的会执行 handlerFunc 的代码
	// method POST, GET, PUT
	RouteV3(method string, pattern string, handlerFunc func(ctx *Context))

	// StartV3 启动我们的服务器
	StartV3(address string) error
}

type ServerV2 interface {

	// RouteV2 设定一个路由，命中该路由的会执行 handlerFunc 的代码
	//Route(pattern string, handlerFunc http.HandlerFunc)
	RouteV2(pattern string, handlerFunc func(ctx *Context))

	// StartV2 启动我们的服务器
	StartV2(address string) error
}

// Server 是 http server 的顶级抽象
type Server interface {

	// Route 设定一个路由，命中该路由的会执行 handlerFunc 的代码
	Route(pattern string, handlerFunc http.HandlerFunc)

	// Start 启动我们的服务器
	Start(address string) error
}

// sdkHttpServerV3 这个是基于 net/http 这个包实现的 http server
type sdkHttpServerV3 struct {
	// Name server 的名字，给个标记，日志输出的时候用的上
	Name string
}

// sdkHttpServerV2 这个是基于 net/http 这个包实现的 http server
type sdkHttpServerV2 struct {
	// Name server 的名字，给个标记，日志输出的时候用的上
	Name string
}

// sdkHttpServer 这个是基于 net/http 这个包实现的 http server
type sdkHttpServer struct {
	// Name server 的名字，给个标记，日志输出的时候用的上
	Name string
}

// RouteV3 注册路由
func (s *sdkHttpServerV3) RouteV3(method string, pattern string, handlerFunc func(ctx *Context)) {
	http.HandleFunc(pattern, func(w http.ResponseWriter, r *http.Request) {
		//ctx := &Context{
		//	W: w,
		//	R: r,
		//}
		ctx := NewContext(w, r)
		handlerFunc(ctx) // 调用传进来的函数，函数的入参是在这个方法里面构建的
	})
}

func (s *sdkHttpServerV2) RouteV2(pattern string, handlerFunc func(ctx *Context)) {
	http.HandleFunc(pattern, func(w http.ResponseWriter, r *http.Request) {
		//ctx := &Context{
		//	W: w,
		//	R: r,
		//}
		ctx := NewContext(w, r)
		handlerFunc(ctx) // 调用传进来的函数，函数的入参是在这个方法里面构建的
	})
}

func (s *sdkHttpServer) Route(pattern string, handlerFunc http.HandlerFunc) {
	http.HandleFunc(pattern, handlerFunc)
}

func (s *sdkHttpServerV3) StartV3(address string) error {
	return http.ListenAndServe(address, nil)
}

func (s *sdkHttpServerV2) StartV2(address string) error {
	return http.ListenAndServe(address, nil)
}

func (s *sdkHttpServer) Start(address string) error {
	return http.ListenAndServe(address, nil)
}

func NewHttpServerV3(name string) ServerV3 {
	// 返回一个实际类型是我实现接口的时候，需要取址
	return &sdkHttpServerV3{
		Name: name,
	}
}

func NewHttpServerV2(name string) ServerV2 {
	// 返回一个实际类型是我实现接口的时候，需要取址
	return &sdkHttpServerV2{
		Name: name,
	}
}

func NewHttpServer(name string) Server {
	// 返回一个实际类型是我实现接口的时候，需要取址
	return &sdkHttpServer{
		Name: name,
	}
}

func SignUpV2(ctx *Context) {
	req := &signUpReq{}

	//ctx := &Context{
	//	W: w,
	//	R: r,
	//}
	err := ctx.ReadJson(req)
	if err != nil {
		//fmt.Fprintf(w, "err: %v", err)
		ctx.BadRequestJson(err)
		return
	}

	// 返回一个虚拟的 user id 表示注册成功了
	//fmt.Fprintf(w, "%d", 123)
	fmt.Fprintf(ctx.W, "%d", 123)

	// 返回 json 对象
	resp := &commonResponse{
		Data: 123,
	}

	err = ctx.WriteJson(http.StatusOK, resp)
	if err != nil {
		fmt.Printf("写入响应失败：%v", err)
	}
}

func SignUp(w http.ResponseWriter, r *http.Request) {
	req := &signUpReq{}
	//body, err := io.ReadAll(r.Body)
	//if err != nil {
	//	fmt.Fprintf(w, "read body failed: %v", err)
	//	// 要返回掉，不然就会继续执行后面的代码
	//	return
	//}
	//err = json.Unmarshal(body, req)
	//if err != nil {
	//	fmt.Fprintf(w, "deserialized failed: %v", err)
	//	// 要返回掉，不然就会继续执行后面的代码
	//	return
	//}
	ctx := &Context{
		W: w,
		R: r,
	}
	err := ctx.ReadJson(req)
	if err != nil {
		fmt.Fprintf(w, "err: %v", err)
		return
	}

	// 返回一个虚拟的 user id 表示注册成功了
	fmt.Fprintf(w, "%d", 123)

	// 返回 json 对象
	resp := &commonResponse{
		Data: 123,
	}

	//w.WriteHeader(http.StatusOK)
	//respJson, err := json.Marshal(resp)
	//if err != nil {
	//
	//}
	//fmt.Fprintf(w, string(respJson)) // []byte 和 string 可以互转
	err = ctx.WriteJson(http.StatusOK, resp)
	if err != nil {
		fmt.Printf("写入响应失败：%v", err)
	}
}

type signUpReq struct {
	Email             string `json:"email"`
	Password          string `json:"password"`
	ConfirmedPassword string `json:"confirmed_password"`
}

type commonResponse struct {
	BizCode int         `json:"biz_code"`
	Msg     string      `json:"msg"`
	Data    interface{} `json:"data"`
}