package main

/*
• 编写一个HTTP服务器，此练习为正式作业需要提交并批改
• 鼓励群里讨论，但不建议学习委员和课代表发满分答案，给大家留一点思考空间
• 大家视个人不同情况决定完成到哪个环节，但尽量把1都做完
1.接收客户端 request，并将 request 中带的 header 写入 response header
2.读取当前系统的环境变量中的 VERSION 配置，并写入 response header
3.Server 端记录访问日志包括客户端 IP，HTTP 返回码，输出到 server 端的标准输出 4.当访问 localhost/healthz 时，应返回200
*/
import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
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
	HttpServeer()
}

func QueueServer() {
	f := fibonacci()
	//生成者1
	go func() {
		for true {
			time.Sleep(10 * time.Second)
			fmt.Println("生成者1成数字:", q.Push(f()))
		}
	}()
	//生成者2
	go func() {
		for true {
			time.Sleep(1 * time.Second)
			fmt.Println("生成者2成数字:", q.Push(f()))
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
func HttpServeer() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		flusher, ok := w.(http.Flusher)
		if !ok {
			panic("expected http.ResponseWriter to be an http.Flusher")
		}
		w.Header().Set("X-Content-Type-Options", "nosniff")
		w.Header().Set("VERSION", os.Getenv("GOVERSION"))
		for true {
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(fmt.Sprintf("客户端ip:%s 斐波那契数列 消费： %v\n", r.RemoteAddr, q.Pop())))
			flusher.Flush() // Trigger "chunked" encoding and send a chunk...
		}
	})
	log.Print("Listening on localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
