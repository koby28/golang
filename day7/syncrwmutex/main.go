package main

import (
	"fmt"
	"sync"
	"runtime"
)

func clickWithMutex(total *int, m *sync.RWMutex, ch chan int) {
	for i := 0; i < 1000; i++ {
		m.Lock()
		*total += 1
		m.Unlock()
		//这里是写 下面是读，外层还有线程的竞争
		if i == 500 {
			m.RLock()
			fmt.Println(*total)
			m.RUnlock()
		}
	}
	ch <- 1
}

func main() {

	runtime.GOMAXPROCS(4) //使用多个处理器，不然都是顺序执行。

	m := new(sync.RWMutex)
	count := 0

	ch := make(chan int, 10) //保证输出时count完了

	for i := 0; i < 10; i++ {
		go clickWithMutex(&count, m, ch)
	}

	for i := 0; i < 10; i++ {
		<-ch
	}

	fmt.Printf("count:%d\n", count)
}
