package main

/*
• 基于Channel编写一个简单的单线程生产者消费者模型
• 队列:
队列长度10，队列元素类型为 int
• 生产者:
每1秒往队列中放入一个类型为 int 的元素，队列满时生产者可以阻塞 • 消费者: 每一秒从队列中获取一个元素并打印，队列为空时消费者阻塞
*/
import (
	"fmt"
	"time"
)

func produce(ch chan<- int) {
	for i := 0; i < 10; i++ {
		ch <- i
		time.Sleep(1 * time.Second)
		fmt.Printf("producing data: %d\n", i)
	}
	close(ch)
}

func consumer(ch <-chan int) {
	for k := range ch {
		fmt.Printf("Get data: %d\n", k)
	}
}

func main() {
	ch := make(chan int, 5)
	go produce(ch)
	consumer(ch)
}
