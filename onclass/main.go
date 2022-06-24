package main

import (
	"fmt"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "这是主页")
}

func user(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "这是用户")
}

func createUser(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "这是创建用户")
}

func order(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "这是订单")
}

//func main() {
//	http.HandleFunc("/", home)
//	http.HandleFunc("/user", user)
//	http.HandleFunc("/user/create", createUser)
//	http.HandleFunc("/order", order)
//	http.ListenAndServe(":8080", nil)
//}

func main() {
	server := NewHttpServer("test-server")
	serverV2 := NewHttpServerV2("test-server-v2")
	server.Route("/", home)
	server.Route("/user", user)
	server.Route("/user/create", createUser)
	server.Route("/user/signup", SignUp)
	serverV2.RouteV2("/user/signupv2", SignUpV2)
	server.Route("/order", order)
	server.Start(":8080")
	serverV2.StartV2(":8081")
}
