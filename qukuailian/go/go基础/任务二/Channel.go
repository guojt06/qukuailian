package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

// 题目 ：编写一个Go程序，定义一个函数，该函数接收一个整数指针作为参数，在函数内部将该指针指向的值增加10，然后在主函数中调用该函数并输出修改后的值。
// 考察点 ：指针的使用、值传递与引用传递的区别。
// 题目 ：实现一个函数，接收一个整数切片的指针，将切片中的每个元素乘以2。
// 考察点 ：指针运算、切片操作。
// ✅
// Goroutine
// 题目 ：编写一个程序，使用 go 关键字启动两个协程，一个协程打印从1到10的奇数，另一个协程打印从2到10的偶数。
// 考察点 ： go 关键字的使用、协程的并发执行。
// 题目 ：设计一个任务调度器，接收一组任务（可以用函数表示），并使用协程并发执行这些任务，同时统计每个任务的执行时间。
// 考察点 ：协程原理、并发任务调度。
// ✅
// 面向对象
// 题目 ：定义一个 Shape 接口，包含 Area() 和 Perimeter() 两个方法。然后创建 Rectangle 和 Circle 结构体，实现 Shape 接口。在主函数中，创建这两个结构体的实例，并调用它们的 Area() 和 Perimeter() 方法。
// 考察点 ：接口的定义与实现、面向对象编程风格。
// 题目 ：使用组合的方式创建一个 Person 结构体，包含 Name 和 Age 字段，再创建一个 Employee 结构体，组合 Person 结构体并添加 EmployeeID 字段。为 Employee 结构体实现一个 PrintInfo() 方法，输出员工的信息。
// 考察点 ：组合的使用、方法接收者。
// ✅
// Channel
// 题目 ：编写一个程序，使用通道实现两个协程之间的通信。一个协程生成从1到10的整数，并将这些整数发送到通道中，另一个协程从通道中接收这些整数并打印出来。
// 考察点 ：通道的基本使用、协程间通信。
// 题目 ：实现一个带有缓冲的通道，生产者协程向通道中发送100个整数，消费者协程从通道中接收这些整数并打印。
// 考察点 ：通道的缓冲机制。
// ✅
// 锁机制
// 题目 ：编写一个程序，使用 sync.Mutex 来保护一个共享的计数器。启动10个协程，每个协程对计数器进行1000次递增操作，最后输出计数器的值。
// 考察点 ： sync.Mutex 的使用、并发数据安全。
// 题目 ：使用原子操作（ sync/atomic 包）实现一个无锁的计数器。启动10个协程，每个协程对计数器进行1000次递增操作，最后输出计数器的值。
// 考察点 ：原子操作、并发数据安全。

// 题目 ：编写一个Go程序，定义一个函数，该函数接收一个整数指针作为参数，在函数内部将该指针指向的值增加10，然后在主函数中调用该函数并输出修改后的值。
// 考察点 ：指针的使用、值传递与引用传递的区别。
// 题目 ：实现一个函数，接收一个整数切片的指针，将切片中的每个元素乘以2。
// 考察点 ：指针运算、切片操作。

// 题目 ：编写一个Go程序，定义一个函数，该函数接收一个整数指针作为参数，在函数内部将该指针指向的值增加10，然后在主函数中调用该函数并输出修改后的值。
// 考察点 ：指针的使用、值传递与引用传递的区别。
func increase(num *int) {
	*num += 10
}

// 题目 ：实现一个函数，接收一个整数切片的指针，将切片中的每个元素乘以2。
// 考察点 ：指针运算、切片操作。
func double(num *[]int) {
	for i := 0; i < len(*num); i++ {
		(*num)[i] = (*num)[i] * 2
	}
}

// 题目 ：编写一个程序，使用 go 关键字启动两个协程，一个协程打印从1到10的奇数，另一个协程打印从2到10的偶数。
// 考察点 ： go 关键字的使用、协程的并发执行。
func toNmu() {
	for i := 0; i < 11; i++ {
		if i%2 == 0 {
			fmt.Printf("协程一：%v\n", i)
		}
	}
}

func toNmutwo() {
	for i := 0; i < 11; i++ {
		if i%2 == 1 {
			fmt.Printf("协程二：%v\n", i)
		}
	}
}

// 题目 ：设计一个任务调度器，接收一组任务（可以用函数表示），并使用协程并发执行这些任务，同时统计每个任务的执行时间。
// 考察点 ：协程原理、并发任务调度。
// Task 是一个函数类型，表示任务
type Task func()

// TaskScheduler 是任务调度器
type TaskScheduler struct {
	tasks []Task
}

// AddTask 添加一个任务到调度器
func (scheduler *TaskScheduler) AddTask(task Task) {
	scheduler.tasks = append(scheduler.tasks, task)
}

// Run 并发执行所有任务，并统计每个任务的执行时间
func (scheduler *TaskScheduler) Run() {
	var wg sync.WaitGroup
	for _, task := range scheduler.tasks {
		wg.Add(1)
		go func(task Task) {
			defer wg.Done()
			start := time.Now()
			task()
			duration := time.Since(start)
			fmt.Printf("任务执行时间: %v\n", duration)
		}(task)
	}
	wg.Wait()
}

// 示例任务函数
func taskOne() {
	fmt.Println("任务一：开始")
	time.Sleep(2 * time.Second) // 模拟任务需要2秒
	fmt.Println("任务一：结束")
}
func taskTwo() {
	fmt.Println("任务二：开始")
	time.Sleep(3 * time.Second) // 模拟任务需要3秒
	fmt.Println("任务二：结束")
}

// // 面向对象
// 题目 ：定义一个 Shape 接口，包含 Area() 和 Perimeter() 两个方法。然后创建 Rectangle 和 Circle 结构体，
// 实现 Shape 接口。在主函数中，创建这两个结构体的实例，并调用它们的 Area() 和 Perimeter() 方法。
// 考察点 ：接口的定义与实现、面向对象编程风格。
type Shape interface {
	Area()
	Perimeter()
}

type Rectangle struct {
}

type Circle struct {
}

func (r *Rectangle) Area() {
	fmt.Println("矩形的面积")
}
func (r *Rectangle) Perimeter() {
	fmt.Println("矩形的周长")
}

func (c *Circle) Area() {
	fmt.Println("圆的面积")
}

func (c *Circle) Perimeter() {
	fmt.Println("圆的周长")
}

//// 题目 ：使用组合的方式创建一个 Person 结构体，包含 Name 和 Age 字段，再创建一个 Employee 结构体，
//组合 Person 结构体并添加 EmployeeID 字段。为 Employee 结构体实现一个 PrintInfo() 方法，输出员工的信息。
// 考察点 ：组合的使用、方法接收者。

// Person 结构体
type Person struct {
	Name string
	Age  int
}

type Employee struct {
	Person
	EmployeeID int
}

func (e *Employee) printInfo() {
	fmt.Printf("员工姓名：%s，年龄：%d，员工编号：%d\n", e.Name, e.Age, e.EmployeeID)
}

// 题目 ：编写一个程序，使用通道实现两个协程之间的通信。一个协程生成从1到10的整数，并将这些整数发送到通道中，另一个协程从通道中接收这些整数并打印出来。
// 考察点 ：通道的基本使用、协程间通信。
func setchannel(ch chan<- int, wg *sync.WaitGroup) {
	for i := 0; i < 11; i++ {
		ch <- i
		fmt.Printf("生产: %d\n", i)
	}
	close(ch) // 关闭通道
	wg.Done()
	fmt.Println("生产者: 完成所有发送")

}

func getchannel(chi <-chan int, wg *sync.WaitGroup) {
	for num := range chi {
		fmt.Printf("接收: %d\n", num)
	}
	wg.Done()
	fmt.Println("消费者: 消费完成")

}

// 题目 ：实现一个带有缓冲的通道，生产者协程向通道中发送100个整数，消费者协程从通道中接收这些整数并打印。
// 考察点 ：通道的缓冲机制。
func setchannelbuffer(ch chan<- int, wg *sync.WaitGroup) {
	for i := 0; i < 100; i++ {
		ch <- i
		fmt.Printf("生产: %d\n", i)
	}
	close(ch) // 关闭通道
	wg.Done()
}

func getchannelbuffer(chi <-chan int, wg *sync.WaitGroup) {
	for v := range chi {
		fmt.Printf("接收: %d\n", v)
	}
	wg.Done()
}

// 锁机制
// 题目 ：编写一个程序，使用 sync.Mutex 来保护一个共享的计数器。启动10个协程，每个协程对计数器进行1000次递增操作，最后输出计数器的值。
// 考察点 ： sync.Mutex 的使用、并发数据安全。
// 题目 ：使用原子操作（ sync/atomic 包）实现一个无锁的计数器。启动10个协程，每个协程对计数器进行1000次递增操作，最后输出计数器的值。
// 考察点 ：原子操作、并发数据安全。

func synchronize(counter int, mutex *sync.Mutex, wg *sync.WaitGroup) {
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < 1000; j++ {
				//加锁
				mutex.Lock()
				counter++
				//解锁
				mutex.Unlock()

			}
			// 等待所有goroutine完成
			fmt.Printf("计数器的值为：%d\n", counter)
		}()
		wg.Wait()
	}
	// 等待所有goroutine完成
	fmt.Printf("计数器的值为：%d\n", counter)
}

func synchronize1(counter int64, mutex *sync.Mutex, wg *sync.WaitGroup) {
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < 1000; j++ {
				atomic.AddInt64(&counter, 1)

			}
			// 等待所有goroutine完成
			fmt.Printf("计数器的值为：%d\n", counter)
		}()
		wg.Wait()
	}
	// 等待所有goroutine完成
	fmt.Printf("计数器的值为：%d\n", counter)
}

// // 题目 ：编写一个Go程序，定义一个函数，该函数接收一个整数指针作为参数，在函数内部将该指针指向的值增加10，然后在主函数中调用该函数并输出修改后的值。
// 考察点 ：指针的使用、值传递与引用传递的区别。
// 题目 ：实现一个函数，接收一个整数切片的指针，将切片中的每个元素乘以2。
// 考察点 ：指针运算、切片操作。
func increase1(num *int) {
	*num += 10
}

func double1(num *[]int) {}

func main() {
	num := 15
	// increase(&num)
	// fmt.Printf("修改后的值为：%d\n", num)

	// list := []int{1, 2, 3, 4, 54, 35, 5}
	// double(&list)
	// fmt.Printf("修改后的值为：%v\n", list)

	// go toNmu()
	// go toNmutwo()

	// time.Sleep(1 * time.Second)

	// i := []int{}
	// i = append(i, 123, 43535)
	// fmt.Println(i)

	// scheduler := &TaskScheduler{}
	// // 添加任务
	// scheduler.AddTask(taskOne)
	// scheduler.AddTask(taskTwo)
	// // 运行任务
	// scheduler.Run()

	// shape := &Rectangle{}
	// shape.Area()
	// shape.Perimeter()

	// shape1 := &Circle{}
	// shape1.Area()
	// shape1.Perimeter()

	// 创建 Employee 实例
	// emp := Employee{
	// 	Person: Person{
	// 		Name: "张三",
	// 		Age:  30,
	// 	},
	// 	EmployeeID: 45645,
	// }

	// emp.printInfo()

	// //创建管道
	// a := make(chan int)
	// //创建计数器
	// var wg sync.WaitGroup
	// wg.Add(2)
	// //启动生产者
	// go setchannel(a, &wg)
	// //启动消费者
	// go getchannel(a, &wg)
	// //等待执行完成
	// wg.Wait()
	// fmt.Println("完成")

	//创建一个缓存管道
	// wr := make(chan int, 10)
	// //创建计数器
	// var wg sync.WaitGroup
	// wg.Add(2)
	// go setchannelbuffer(wr, &wg)
	// go getchannelbuffer(wr, &wg)
	// wg.Wait()
	// fmt.Println("完成")

	// synchronize(0, &sync.Mutex{}, &sync.WaitGroup{})
	// synchronize1(0, &sync.Mutex{}, &sync.WaitGroup{})

	increase1(&num)
	fmt.Printf("计数器的值为：%d\n", num)
}
