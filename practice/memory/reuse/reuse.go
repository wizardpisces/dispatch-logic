package main

import (
	"fmt"
	"sync"
)

type Person struct {
	Name string
	Age  int
}

func main() {
	pool := &sync.Pool{
		New: func() interface{} {
			// 当池为空时，调用该函数创建新实例
			return &Person{"Li Lei", 11} // 创建一个新对象
		},
	}
	// 从对象池中获取一个实例
	p1 := pool.Get().(*Person)
	fmt.Println(p1) // &{Li Lei 11}
	p1.Name = "Han Meimei"
	p1.Age = 12
	fmt.Println(p1) // &{Han Meimei 12}
	// 将该实例放回对象池中
	pool.Put(p1)
	// 再次从对象池中获取一个实例
	p2 := pool.Get().(*Person)
	fmt.Println(p2) // &{Han Meimei 12}
}
