package main

import (
	"fmt"
	"sync"
)

func main() {
	PrintOnce()
	PrintOnce()
	PrintOnce()
}

var once sync.Once

// PrintOnce 这个方法，不管调用几次，只会输出一次
func PrintOnce() {
	// 你可能把这个东西声明在局部变量里，那就是没效果的
	//var once sync.Once
	once.Do(func() {
		fmt.Println("只输出一次")
	})
}
