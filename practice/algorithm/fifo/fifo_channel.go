package fifo

// 消费者跟生产者
import (
	"fmt"
	"sync"
)

// 任务结构体
type Task struct {
	id int
}

// 生产者结构体
type Producer struct {
	out chan *Task // 输出通道，用于向消费者发送任务
}

// 生产方法，不断地向输出通道发送新的任务
func (p *Producer) Produce() {
	id := 0
	for {
		id++
		task := &Task{id: id}
		fmt.Println("Producer:", task.id)
		p.out <- task // 发送任务到输出通道
	}
}

// 消费者结构体
type Consumer struct {
	in chan *Task // 输入通道，用于从生产者接收任务
}

// 消费方法，在无限循环中不断地从输入通道接收新的任务，并执行它们。
func (c *Consumer) Consume(wg *sync.WaitGroup) {
	defer wg.Done() // 在退出时通知 WaitGroup

	for task := range c.in { // 从输入通道接收新的任务（如果没有则阻塞）
		fmt.Println("Consumer:", task.id)
	}
}

func Test_fifo_channel() {
	ch := make(chan *Task) // 创建一个共享的无缓冲通道

	p := &Producer{out: ch} // 创建一个生产者，并将输出通道设置为共享的通道

	var wg sync.WaitGroup // 创建一个 WaitGroup，用于等待所有消费者结束

	for i := 0; i < 3; i++ { // 创建三个消费者，并将输入通道设置为共享的通道
		c := &Consumer{in: ch}
		wg.Add(1)         // 增加 WaitGroup 的计数器
		go c.Consume(&wg) // 启动消费方法（goroutine）
	}

	go p.Produce() // 启动生产方法（goroutine）

	wg.Wait() // 等待所有消费者结束（主 goroutine 阻塞）
}
