/*
 暂时未弄清楚原子操作
*/

package main

import (
	"fmt"
	"runtime"
	"sync/atomic"
	"time"
)

func main() {
	ops := uint64(0)

	// 为了模拟并行更新，我们使用50个协程来每隔1毫秒来
	// 增加一下counter值，注意这里的50协程里面的for循环，
	// 也就是说如果主协程不退出，这些协程将永远运行下去
	// 所以这个程序每次输出的值有可能不一样
	for i := 0; i < 1; i++ {
		go func() {
			for {
				// 为了能够保证counter值增加的原子性，我们使用
				// atomic包中的AddUint64方法，将counter的地址和
				// 需要增加的值传递给函数即可
				fmt.Println("ops: ", ops)
				atomic.AddUint64(&ops, 1)

				//允许其他协程来处理
				runtime.Gosched()
			}
		}()
	}

	time.Sleep(time.Second)

	// 为了能够在counter仍被其他协程更新值的同时安全访问counter值，
	// 我们获取一个当前counter值的拷贝，这里就是opsFinal，需要把
	// ops的地址传递给函数`LoadUint64`
	opsFinal := atomic.LoadUint64(&ops)
	fmt.Println("ops:", opsFinal)
}
