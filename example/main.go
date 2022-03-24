package main

import (
	"log"
	"net/http"
	_ "net/http/pprof" // 会自动注册 handler 到 http server，方便通过 http 接口获取程序运行采样报告
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(1)
	runtime.SetMutexProfileFraction(1) // 开启对锁调用的跟踪
	runtime.SetBlockProfileRate(1)     // 开启对阻塞操作的跟踪

	// http://127.0.0.1:6060/debug/pprof/
	if err := http.ListenAndServe(":6060", nil); err != nil {
		log.Fatal(err)
	}
}
