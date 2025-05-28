package main

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

// 错误处理
func err() (int32, error) {
	var err error = nil
	return 1, err
}

// 错误处理

// 闭包
func get() func() int {
	i := 0
	return func() int {
		i++
		// 不要返回i++, go自增不回返值
		return i
	}
}

// 闭包
// 方法
type Circle struct {
	radius float64
}

func (c Circle) getAera() float64 {
	return c.radius * 3.14 * c.radius
}

func (c *Circle) setAera(radius float64) {
	c.radius = radius
}

// 方法

// 接口
type myInterface interface {
	methodA()
}

type s struct {
}

type s2 struct {
}

func (s s) methodA() {
	fmt.Println("hello method")
}

func (s2 s2) methodA() {
	fmt.Println("hello method s2")
}

func connectMySQL() {
	pool, err := sql.Open("mysql", "root:123456@/ksp")
	if err != nil {
		panic(err)
	}

	pool.SetConnMaxLifetime(time.Minute * 3)
	pool.SetMaxOpenConns(10)
	pool.SetMaxIdleConns(10)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	var name string
	pool.QueryRowContext(ctx, "SELECT name FROM USER").Scan(&name)
	fmt.Println(name)

	time.Sleep(10 * time.Second)
	pool.Close()
}

func main() {
	connectMySQL()

	// var i interface{} = 42
	// // fmt.Printf("Dynamic type: %T, Dynamic value: %v\n", i, i)
	// map01 := make(map[any]any, 5)
	// map01["key"] = "value"
	// // map01[i.(type)] = i  报错
	// map01[reflect.TypeOf(i).String()] = i
	// fmt.Println(map01["int"])
	// // 类型
	// var i1 interface{} = s{}
	// // var i2 interface{} = s2{}

	// ii1 := i1.(s)
	// ii1.methodA()

	// // 通道阻塞
	// ch := make(chan int)
	// ch <- 1
	// ch <- 2
	// close(ch)
	// fmt.Println("chan阻塞")

	// // Context
	// v := context.WithValue(context.Background(), "key", "hello")
	// v.Value("key")
	// ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	// go func() {
	// 	for range 10 {
	// 		time.Sleep(time.Second)
	// 		select {
	// 		case <-ctx.Done():
	// 			fmt.Println(ctx.Err())
	// 			return
	// 		case t := <-time.After(1):
	// 			fmt.Println(t)
	// 		}
	// 	}
	// }()
	// cancel()
	// time.Sleep(6 * time.Second)

	// Once
	// once := sync.Once{}
	// for range 10 {
	// 	go func() {
	// 		time.Sleep(5 * time.Second)
	// 		once.Do(func() { fmt.Println("hello do") })
	// 	}()
	// }

	// // 主线程必须慢于子线程结束
	// time.Sleep(10 * time.Second)

	// // 锁
	// mx := sync.Mutex{}
	// wg := sync.WaitGroup{}
	// wg.Add(10)
	// ii := 0
	// for range 10 {
	// 	go func() {
	// 		defer wg.Done()
	// 		mx.Lock()
	// 		defer mx.Unlock()
	// 		for j := 0; j < 1000; j++ {
	// 			ii++
	// 		}
	// 	}()
	// }

	// wg.Wait()
	// fmt.Println(ii)

	// waitGroup
	// wg := sync.WaitGroup{}

	// var ch01 = make(chan int32, 1)
	// var ch02 = make(chan int32, 1)
	// var ch03 = make(chan int32, 1)

	// wg.Add(3)

	// // 开启携程, 并传入只读通道
	// go func(ch chan<- int32) {
	// 	defer wg.Done()
	// 	defer close(ch)
	// 	time.Sleep(3 * time.Second)
	// 	ch <- 1
	// }(ch01)

	// go func(ch chan<- int32) {
	// 	defer wg.Done()
	// 	defer close(ch)
	// 	time.Sleep(5 * time.Second)
	// 	ch <- 2
	// }(ch02)

	// go func(ch chan<- int32) {
	// 	defer wg.Done()
	// 	defer close(ch)
	// 	time.Sleep(8 * time.Second)
	// 	ch <- 3
	// }(ch03)

	// // 等待通道有值, 否则阻塞
	// wg.Wait()
	// select {
	// case val01 := <-ch01:
	// 	fmt.Print(val01)
	// case val02 := <-ch02:
	// 	fmt.Print(val02)
	// case val03 := <-ch03:
	// 	fmt.Print(val03)
	// }

	// 接口(实现泛型)
	// var i myInterface = s{}
	// i.methodA()

	// var ii interface{} = s{}
	// ii.methodA()

	// 匿名函数
	// fn := func(a int32, b int32) int32 {
	// 	return a + b
	// }

	// type cb func(int32, int32) int32
	// useCb := func(callback cb, a int32, b int32) int32 {
	// 	return callback(a, b)
	// }
	// 匿名函数

	// var fnn = get()

	// message, err := err()
	// fmt.Println(fnn())
	// fmt.Println(fnn())
	// fmt.Println(fnn())

	// var c Circle
	// fmt.Println(c.getAera())
	// c.setAera(10)
	// fmt.Println(c.getAera())

	// var arr = [5]int{111, 12, 333}
	// var arr02 = [...]int{111, 12, 333, 555}
	// var arr03 = []int{111, 12, 333, 555}
	// var arr04 = make([]int, 4, 10)

	// fmt.Println(arr)
	// fmt.Println(len(arr))
	// fmt.Println(unsafe.Sizeof(arr))

	// fmt.Println(arr02)
	// fmt.Println(len(arr02))
	// fmt.Println(cap(arr02))
	// fmt.Println(unsafe.Sizeof(arr02))

	// fmt.Println(arr03)
	// fmt.Println(len(arr03))
	// fmt.Println(cap(arr03))
	// fmt.Println(unsafe.Sizeof(arr03))

	// fmt.Println(arr04)
	// fmt.Println(len(arr04))
	// fmt.Println(cap(arr04))
	// fmt.Println(unsafe.Sizeof(arr04))

	// arr05 := append(arr03, 1, 2)
	// fmt.Println(arr05)
	// fmt.Println(len(arr05))
	// fmt.Println(cap(arr05))
	// fmt.Println(unsafe.Sizeof(arr05))

	// fmt.Println(quote.Go())

	// // 一定使用make
	// var ch01 = make(chan int32, 1)
	// // var ch03 = make(<-chan int32, 1) // 只写通道

	// // 开启携程, 并传入只读通道
	// go func(ch chan<- int32) {
	// 	time.Sleep(3 * time.Second)
	// 	ch <- 1
	// 	close(ch)
	// }(ch01)

	// // 等待通道有值, 否则阻塞
	// val := <-ch01
	// fmt.Println(val)
}
