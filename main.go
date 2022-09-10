package main

import (
	"fmt"
	"main/app/Business/TodoList"
	"main/app/ErrorHandle"
	"main/app/Router"
	"os"
	"os/signal"
	"runtime"
	"sync"
	"syscall"
	"time"
)

func init() {
	// 初始化log
	ErrorHandle.Init(os.Stdout, os.Stdout, os.Stderr, os.Stderr)

	TodoList.InitTodoListFile()
}
func main() {
	// 監聽外部輸入已關閉
	runtime.GOMAXPROCS(runtime.NumCPU())
	wg := &sync.WaitGroup{}
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM, syscall.SIGQUIT)
	go func() {
		sig := <-c
		_ = sig
		wg.Done()
	}()
	wg.Add(1)

	// -- 主要功能
	fmt.Println(time.Now().Unix())
	go Router.RouterStart()
	// --

	wg.Wait()
	// 服務結束時保險再存檔一次
	TodoList.WriteTodoListFile()
	ErrorHandle.Info.Printf("Server Shutdown\n")
}
