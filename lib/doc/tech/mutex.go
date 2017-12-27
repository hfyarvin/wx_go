package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"sync"
	"sync/atomic"
	"time"
)

func main() {
	state := make(map[int]int)

	//互斥锁
	mutex := &sync.Mutex{}

	//计数
	ops := int64(0)

	//启动多个协程不读读取这个状态
	for i := 0; i < 20; i++ {
		go func() {
			total := 0
			for {
				// 对于每次读取，我们选取一个key来访问，
				// mutex的`Lock`函数用来保证对状态的
				// 唯一性访问，访问结束后，使用`Unlock`
				// 来解锁，然后增加ops计数器
				key := rand.Intn(5)
				fmt.Println("key:", key)
				mutex.Lock() //先锁
				total += state[key]
				mutex.Unlock()
				atomic.AddInt64(&ops, 1)

				// 为了保证这个协程不会让调度器出于饥饿状态，
				// 我们显式地使用`runtime.Gosched`释放了资源
				// 控制权，这种控制权会在通道操作结束或者
				// time.Sleep结束后自动释放。但是这里我们需要
				// 手动地释放资源控制权
				runtime.Gosched()
			}
		}()
	}

	// 同样我们使用10个协程来模拟写状态
	for i := 0; i < 5; i++ {
		go func() {
			for {
				key := rand.Intn(5)
				val := rand.Intn(100)
				mutex.Lock()
				state[key] = val
				mutex.Unlock()
				atomic.AddInt64(&ops, 1)
				runtime.Gosched()
			}
		}()
	}

	time.Sleep(time.Second)

	opsFinal := atomic.LoadInt64(&ops)
	fmt.Println("ops:", opsFinal)

	//最后锁定并输出状态
	mutex.Lock()
	fmt.Println("state: ", state)
	mutex.Unlock()
}
