package main

/*
• 将练习1.2中的生产者消费者模型修改成为多个生产者和多个消费者模式
*/
import (
	"errors"
	"fmt"
	"time"
)

type Queue struct {
	buf chan int
}

func (q *Queue) Push(i int) interface{} {
	if len(q.buf) >= cap(q.buf) {
		return errors.New("队列已满")
	}
	q.buf <- i
	return i
}
func (q *Queue) Pop() interface{} {
	if len(q.buf) == 0 {
		return errors.New("队列为空")
	}
	return <-q.buf
}

var (
	q = Queue{buf: make(chan int, 10)}
)

func fibonacci() func() int {
	var i, y = 1, 0
	return func() int {
		i, y = y, i+y
		return y
	}
}

func main() {
	QueueServer()
	//HttpServeer()
}

func QueueServer() {
	f := fibonacci()
	//生产者1
	go func() {
		for true {
			time.Sleep(10 * time.Second)
			fmt.Println("生产者1成数字:", q.Push(f()))
		}
	}()
	//生产者2
	go func() {
		for true {
			time.Sleep(1 * time.Second)
			fmt.Println("生产者2成数字:", q.Push(f()))
		}
	}()
	//消费者1
	go func() {
		for true {
			fmt.Println("消费者1消费:", q.Pop())
			time.Sleep(10 * time.Second)
		}
	}()

	//消费者2
	go func() {
		for true {
			time.Sleep(2 * time.Second)
			fmt.Println("消费者2消费:", q.Pop())
		}
	}()
	<-time.After(60 * time.Second)
}
