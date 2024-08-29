package test

import (
	"fmt"
	"sync"
	"time"
)

func WsClose() {
	// 创建一个无缓冲的通道用于发送信号
	signalChan := make(chan struct{})
	// secondChan := make(chan bool)

	// 使用 sync.WaitGroup 来等待所有的 goroutines 完成
	var wg sync.WaitGroup

	// 启动多个 goroutines
	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			fmt.Printf("Goroutine %d is waiting for the signal...\n", id)
			// 每个 goroutine 都会阻塞在这里等待信号

			for {
				select {
				case <-signalChan:
					fmt.Printf("++++++++++++++++++Goroutine %d received the signal.\n", id)
					return
				default:
					fmt.Printf("-----------------------Goroutine %d is sleeping...\n", id)
					time.Sleep(1 * time.Second)
				}
			}
		}(i)
	}

	// 主 goroutine 暂停一段时间，让其他 goroutines 进入阻塞状态
	// 发送信号给所有等待的 goroutines
	fmt.Println("Sending signal to all goroutines...")
	time.Sleep(5 * time.Second)
	close(signalChan)

	// 等待所有 goroutines 完成
	wg.Wait()
	fmt.Println("All goroutines have finished.")
}
