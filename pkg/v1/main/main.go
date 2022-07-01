package main

import (
	"fmt"
	webv1 "geektime/toy-web/pkg/v1"
	"net/http"
)

func main() {
	server := webv1.NewSdkHttpServer("test-server")
	//server.Route("/", home)
	//server.Route("/user", user)
	//server.Route("/user/create", createUser)
	server.Route(http.MethodGet, "/user/signup", SignUp)
	//server.Route("/order", order)
	err := server.Start(":8080")
	if err != nil {
		panic(err)
	}
}

func SignUp(ctx *webv1.Context) {
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
