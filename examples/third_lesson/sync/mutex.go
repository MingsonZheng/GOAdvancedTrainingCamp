package main

import (
	"sync"
)

var mutex sync.Mutex
var rwMutex sync.RWMutex

func Mutex() {
	mutex.Lock()
	// 你的业务代码
	// 然后你这里 panic 掉了
	// 所以 mutex 解锁需要使用 defer
	defer mutex.Unlock()
	// 你的代码
}

// RwMutex 一般情况下都用读写锁
func RwMutex() {
	// 加读锁
	rwMutex.RLock()
	defer rwMutex.RUnlock()

	// 也可以加写锁
	rwMutex.Lock()
	defer rwMutex.Unlock()
}

// Failed1 不可重入例子
func Failed1() {
	mutex.Lock()
	defer mutex.Unlock()

	// 这一句会死锁
	// 但是如果你只有一个goroutine，那么这一个会导致程序崩溃
	mutex.Lock()
	defer mutex.Unlock()
}

// Failed2 不可升级
func Failed2() {
	rwMutex.RLock()
	defer rwMutex.RUnlock()

	// 这一句会死锁
	// 但是如果你只有一个goroutine，那么这一个会导致程序崩溃
	mutex.Lock()
	defer mutex.Unlock()
}
