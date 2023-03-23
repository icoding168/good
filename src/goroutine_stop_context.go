package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	ctx, cancel := context.WithCancel(context.Background())

	for i := 0; i < 5; i++ { // 启动 5 个监控狗协程
		wg.Add(1)
		go func(name string) {
			defer wg.Done()
			watchDog(ctx, name)
		}(fmt.Sprintf("【监控狗%d】", i+1))
	}

	time.Sleep(5 * time.Second) // 先让监控狗监控 5 秒

	// 发送停止指令到所有监控狗的 context 中
	cancel()

	wg.Wait()
}

func watchDog(ctx context.Context, name string) {
	// 开启 for select 循环，一直后台监控
	for {
		select {
		case <-ctx.Done():
			fmt.Println(name, "停止指令已收到，马上停止")
			return
		default:
			fmt.Println(name, "正在监控……")
		}
		time.Sleep(1 * time.Second)
	}
}
