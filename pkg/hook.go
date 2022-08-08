package web

import (
	"context"
	"fmt"
	"sync"
	"time"
)

// Hook 是一个钩子函数。注意，
// ctx 是一个有超时机制的 context.Context
// 所以你必须处理超时的问题
type Hook func(ctx context.Context) error

// BuildCloseServerHook 这里其实可以考虑使用 errgroup，
// 但是我们这里不用是希望每个 server 单独关闭
// 互相之间不影响
func BuildCloseServerHook(servers ...Server) Hook {
	return func(ctx context.Context) error {
		wg := sync.WaitGroup{}
		doneCh := make(chan struct{})
		wg.Add(len(servers)) // 添加 server 的数量，因为我想一个个关闭完

		for _, s := range servers { // 遍历 server
			go func(svr Server) { // 开 goroutine 每一个 server 单独关闭，为了不影响其他 server
				err := svr.Shutdown(ctx)
				if err != nil {
					fmt.Printf("server shutdown error: %v \n", err)
				}
				time.Sleep(time.Second)
				wg.Done() // 这里就是减一
			}(s)
		}
		go func() {
			wg.Wait() // server 没有关完就会卡在这里，关完了就会继续往下执行
			doneCh <- struct{}{}
		}()
		select {
		case <-ctx.Done(): // 执行的时候没关完
			fmt.Printf("closing servers timeout \n")
			return ErrorHookTimeout
		case <-doneCh: // 执行的时候关完了
			fmt.Printf("close all servers \n")
			return nil
		}
	}
}
