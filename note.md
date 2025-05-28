[toc]
# 0. 序言

这篇笔记用于快速记录Go语言的特性.

# 1. 包设置相关

参考[史上最全的Go语言模块（Module）管理详解（基于Go1.19）-腾讯云开发者社区-腾讯云](https://cloud.tencent.com/developer/article/2171597)

## 1.1 包网站

pkg.go.dev

## 1.2 代理设置

你懂的, 通过设置环境变量修改

go env -w GOPROXY=goproxy.cn

## 1.3 包下载路径

修改GOPATH环境变量(模块的存储路径)

## 1.4 常用包管理命令

* go mod tidy (本地模块同步到go.mod文件中)
* go mod download (命令下载所有在go.mod文件中声明的依赖)
* go install pkg@version
* go get pkg

go get和go install的区别:

> 用途不同：go get 主要用于管理项目的依赖项，而 go install 主要用于安装可执行文件
> 对 go.mod 文件的影响：go get 会修改 go.mod 文件，而 go install 不会
> 编译行为：go get 不会编译包，而 go install 会编译并安装可执行文件
> 总之，go get 适用于依赖项管理，而 go install 适用于安装命令行工具
> 自 Go 1.16 起，go install 可以接受一个版本后缀，并且可以忽略当前目录的 go.mod 文件

## 1.5 包指令

go的依赖管理机制参考了Java, 非常符合**模块开发的原则**

每个模块**可以属于不同项目, 也独立出来**, 包的名称一般为包的独立地址, 比如github.com/ljwgo/xxx. 但也属于任何使用它的项目 (具有本地路径)

这样做的好处是, 当你的包缺时, 会自动从github.com/ljwgo/xxx网址下载, 如果下载失败, go编译程序可能报错,直到超时.

### 1.5.0 mod init新建指令

当我们要构建自己的包(更准确说是模块), 使用

```go
go mod init example.com/pkg
```

为模块example.com/pkg建立一个依赖文件(类似JavaScript的package.json, .lock文件)

### 1.5.1 replace指令

在go.mod文件中加入

```go
replace github.com/ljwgo/xxx ../ModuleXXX
```

就能为包分配本地路径了, 但这样做的缺点也很明显**每个go.mod文件都要设置, 并且路径根据mod文件所在位置还会不同**.

### 1.5.2 go.work管理文件

使用

```go
go work init ModuleA ModuleB Demo
```

在全局上为用到的每个模块设置本地路径

go.work内容

```go
go 1.19
use (
    ./Demo
    ./ModuleA
    ./ModuleB
)
```

如果修改了模块名(如更换了网址), 则可以使用模块别名

```go
replace github.com/ljwgo/MA v1.2.3  => ./ModuleA 
replace github.com/ljwgo/MB v1.2.3 => ./ModuleB 
```

# 2. vscode设置

为了让vscode显示智能提示, 需要下载go的一个语言服务器(满足语言服务协议)

go get -v golang.org/x/tools/gopls

# 3. go变量
## 3.1 go变量作用域

go中没有class概念, 它是函数式的

当它仍可以类似public, protected的控制变量的作用域

* 当标识符（包括常量、变量、类型、函数名、结构字段等等）以一个大写字母开头，如：Group1，那么使用这种形式的标识符的对象就可以被外部包的代码所使用（客户端程序需要先导入这个包），这被称为导出（像面向对象语言中的 public）；

* 标识符如果以小写字母开头，则对包外是不可见的，但是他们在整个包的内部是可见并且可用的（像面向对象语言中的 protected ）。

## 3.2 变量声明

go中变量类型可以通过编译器进行推导

```go
var i = 10
// 你也可以显示指定类型
var ii int32 = 100
// 声明并赋值
iii := 123
```

## 3.3 特殊变量

### 3.3.1 _垃圾桶

_ 实际上是一个只写变量，你不能得到它的值。这样做是因为 Go 语言中你必须使用所有被声明的变量，但有时你并不需要使用从一个函数得到的所有返回值。

### 3.3.2 常量

常量还可以用作枚举

```go
const (
    Unknown = 0
    Female = 1
    Male = 2
)
```

### 3.3.3 iota特殊常量

iota，特殊常量，可以认为是一个可以被编译器修改的常量。
iota 在 const关键字出现时将被重置为 0(const 内部的第一行之前)，const 中每新增一行常量声明将使 iota 计数一次(iota 可理解为 const 语句块中的行索引)。
iota 可以被用作枚举值：

```go
const (
    a = iota
    b = iota
    c = iota
)
```

第一个 iota 等于 0，每当 iota 在新的一行被使用时，它的值都会自动加 1；所以 a=0, b=1, c=2

### 3.3.4 空值

* nil
* 0
* ""

## 3.4 初始化

数据初始化依旧使用**{}**初始化器

# 4. go控制语句

## 4.1 没有while关键字, 只有for

使用

```go
for true{
}
```

代替

## 4.2 迭代语句

Go 语言中 range 关键字用于 for 循环中迭代数组(array)、切片(slice)、通道(channel)或集合(map)的元素。在数组和切片中它返回元素的索引和索引对应的值，在集合中返回 key-value 对。

```go
// 遍历map
for key, val := range oldMap {
  fmt.Println(key + val)
}

// 遍历通道(直到通道关闭)
ch = make(chan int, 2)
ch <- 1
ch <- 2
close(ch)

for i := range ch {
  fmt.Println(i)
}

// 通过 range 获取参数列表:
package main

import (
    "fmt"
    "os"
)

func main() {
    fmt.Println(len(os.Args))
    for _, arg := range os.Args {
        fmt.Println(arg)
    }
}

//Go 中的中文采用 UTF-8 编码，因此逐个遍历字符时必须采用 for-each 形式
/*
str: hello
0x68 h, 0x65 e, 0x6c l, 0x6c l, 0x6f o,               
0x68, 0x65, 0x6c, 0x6c, 0x6f,                         
                                                      
str: 中国人                                           
0x4e2d 中, 0x56fd 国, 0x4eba 人,                      
0xe4, 0xb8, 0xad, 0xe5, 0x9b, 0xbd, 0xe4, 0xba, 0xba, 
*/
```

# 5. go函数

## 5.1 go函数可以返回多个值 (这也是go解决try-catch不足的地方, 让错误处理流程更明确)

```go
func swap(x, y string) (string, string, error) {
    var err error = nil
    return y, x, err
}

func main(){
    y, x, err := swap(1, 2)
    if (err == nil) {
        // do somthing here!
    }
    else {
        // deal with error here!
    }
}
```

## 5.2 和c不同, go支持匿名函数. 并且函数可以作为实参传递, 类型不需要很明确(整个函数的类型)

```go
// 不能是全局函数
func main() {
	// 匿名函数
	fn := func(a int32, b int32) int32 {
		return a + b
	}

	type cb func(int32, int32) int32
	useCb := func(callback cb, a int32, b int32) int32 {
		return callback(a, b)
	}
}
```

## 5.3 支持闭包(延长局部变量生命周期)

类似装箱拆箱, 闭包变量会从栈逃逸到堆

```go
func get() (func() int) {
  i := 0
  return func() int {
    i++;
    // 不要返回i++, go自增不回返值
    return i
  }
}
```

## 5.4 方法

go是函数式编程的, 它没有类, 但使用struct代替. 可以实现方法(类本质就是给普通函数隐式传递一个指向对象的指针而已)
go的方法在返回值上稍有不同

```go
type Circle struct {
	radius float64
}

func (c Circle) getAera() float64 {
	return c.radius * 3.14 * c.radius
}

// 如果要修改结构体成员, 务必使用指针
func (c *Circle) setAera(radius float64) {
	c.radius = radius
}

// main中调用
var c Circle
fmt.Println(c.getAera())
c.setAera(10)
fmt.Println(c.getAera())
```

# 6. go数据类型

## 6.1 数组类型

其中...含义是**让编译器自行计算arr02的size大小**, **如果不加入..., arr02为切片, 加入后为数组**

```go
var arr = [5]int{111, 12, 333}
var arr02 = [...]int{111, 12, 333}
fmt.Println(arr)
fmt.Println(len(arr))
fmt.Println(unsafe.Sizeof(arr))

fmt.Println(arr02)
fmt.Println(len(arr02))
fmt.Println(unsafe.Sizeof(arr02))
```

## 6.2 slice切片

slice可以理解为更轻量化的vector(c++). 它是可变长的. 在go中, 它是数组的轻量化封装. **数组只要不设置大小就是切片了**.

```go
var arr03 = []int{111, 12, 333, 555}
// 另一种方式是使用make(类型, 元素size, 站内存大小). make自动的介绍挺详细的. 
var arr04 = make([]int, 4, 10)

fmt.Println(arr03)
fmt.Println(len(arr03))
fmt.Println(cap(arr03))
fmt.Println(unsafe.Sizeof(arr03))

fmt.Println(arr04)
fmt.Println(len(arr04))
fmt.Println(cap(arr04))
fmt.Println(unsafe.Sizeof(arr04))

// 改变切片长度
arr05 := append(arr03, 1, 2)
fmt.Println(arr05)
fmt.Println(len(arr05))
fmt.Println(cap(arr05))
fmt.Println(unsafe.Sizeof(arr05))
```

slice结构有三个值: 指向数组的指针, size大小和capacity大小. 因此slice的字节大小固定是24.

- Go 语言的数组与c不同, 是值，其长度是其类型的一部分，作为函数参数时，是 值传递，函数中的修改对调用者不可见
- Go 语言中对数组的处理，一般采用 切片 的方式，切片包含对底层数组内容的引用，作为函数参数时，类似于 指针传递，函数中的修改对调用者可见

slice可以像python一样进行**分割操作**

```go
var arr03 = []int{111, 12, 333, 555}
s := arr03[1:len(arr03)]
```



## 6.3 channel通道

go里经典的名言

```go
Don't communicate by sharing memory, share memory by communicating.
```

就是说在共享内存前, 请进行通信, 保证同步.

而在多个协程中同步使用的就是通道

```go
// 一定使用make
var ch01 = make(chan int32, 1)
// var ch03 = make(<-chan int32, 1) // 只写通道

// 开启协程, 并传入只读通道
go func(ch chan<- int32) {
    time.Sleep(3 * time.Second)
    ch <- 1
    close(ch)  // 记得关闭通道
}(ch01)

// 主协程等待通道有值, 无值阻塞
val := <-ch01
fmt.Println(val)
```

有关并发的, 会专门在并发小节中写

注意, 如果通道缓存满了, 就会阻塞

```go
// 通道阻塞
ch := make(chan int)
ch <- 1
ch <- 2
close(ch)
fmt.Println("chan阻塞")
```



## 6.4 接口类型

1. go的接口类型和传统OOP语言的接口不同, 它更类似于元类(也就是Object). 它是go实现多态的机制. go中, **只要某个结构包含接口的全部定义, 它就隐式继承该接口**
2. 接口变量实际上包含了两个部分：
   动态类型：存储实际的值类型。
   动态值：存储具体的值。

```go
// 接口
type myInterface interface {
	methodA()
}

type s struct {
}

// s结构隐式实现结构myInterface的所有方法
func (s s) methodA() {
	fmt.Println("hello method")
}

func main() {
	// 接口(实现泛型)
	var i myInterface = s{}
	i.methodA()
}
```

```go
package main
import "fmt"
func main() {
  var i interface{} = 42
  fmt.Printf("Dynamic type: %T, Dynamic value: %v\n", i, i)
}
```

2. 接口使用**组合代替继承**

```go
type myInterface interface {
	methodA()
}

type myInterface2 interface {
	methodB()
}

// 组合
type myInterface2 interface {
	methodB()
}
```

据说这样做可以减轻**继承带来的脆弱基类问题(基类冻结)**

3. 空接口

   空接口定位类似Object. 容纳百川.

   ```go
   var ii interface{} = s{}
   ii.methodA()  // 报错
   ```

   **使用interface, 因为涉及运行时类型检查, 因此会影响效率**

# 7. go并发

## 7.1 go协程

go的高并发特征依赖于**协程**, 有关协程这里简单说一下

线程或进程的切换都是依靠操作系统(内核)进行的, 因为存在用户态到内核态的切换, 因此浪费效率. 线程一旦阻塞, 程序会在时间片未用完前就将cpu控制权让出. 

协程可以做到当线程阻塞时, 自动切换运行的协程(新建线程或使用其它旧线程), 依旧把控cpu的控制权.

使用go开启协程, 使用通道进行同步通信

```go
go func(){
  // do something here  
}()
```

## 7.2 go的GMP模型

参考[[Go三关-典藏版\]Golang调度器GPM原理与调度全分析 - 知乎](https://zhuanlan.zhihu.com/p/323271088)

参考[【Golang】协程让出、抢占、监控、调度 - 知乎](https://www.zhihu.com/zvideo/1308438960025776128)

简单总结一下

* go的GMP分别是G协程, M线程, P调度器.

* P看作G的缓存池, 和M直接通信, 非必要M不会访问G全局run queue. P的数量是编译前确定好的, M的数量会因为阻塞而增加(M >= P, 并且自旋M的数量一定为P, 自旋M能保证G被尽快执行). 

* 因此,不难看出, 需要一个总监控线程. 保证定时任务及时无误的触发(挂起协程及时执行). 对于io阻塞, 需要进行轮询. 这是主动让出cpu的情形.

每个协程一次最多运行10ms, 对于不主动让出的cpu, 将采取**抢占**方式, 关于抢占, 上面的视频我看得晕乎乎的.

* P的数量由启动时环境变量 `$GOMAXPROCS` 或者是由 `runtime` 的方法 `GOMAXPROCS()` 决定. 
* 程序运行开始, 会创建一个`m0`线程, 这个M对应的实例会在全局变量`runtime.m0`中, 每个线程都有一个特殊的`g0`线程, 它负责系统调用, 调度其它G，G0不指向任何可执行的函数. m0也有g0线程. 之后是调度器初始化：初始化m0、栈、垃圾回收，以及创建和初始化由GOMAXPROCS个P构成的P列表. `runtime.main`, 代码经过编译后，`runtime.main`会调用`main.main`，程序启动时会为`runtime.main`创建goroutine，称它为main goroutine吧，然后把main goroutine加入到P的本地队列

## 7.3 go并发控制

### 7.3.1 select 语句

select会随机执行一个可运行的case。如果没有case可运行，它将阻塞，直到有case可运行。
select语句只能用于通道操作，每个 case 必须是一个通道操作，要么是发送要么是接收。
select语句会监听所有指定的通道上的操作，一旦其中一个通道准备好就会执行相应的代码块。
如果多个通道都准备好，那么 select 语句会随机选择一个通道执行 (避免某些操作饿死)。存在default时, 如果所有通道都没有准备好，那么执行 default 块中的代码。

```go
select {
  case <- channel1:
    // 执行的代码
  case value := <- channel2:
    // 执行的代码
  case channel3 <- value:
    // 执行的代码
    // 你可以定义任意数量的 case
  default:
    // 所有通道都没有准备好，执行的代码
}
```

### 7.3.2 waitGroups 操作(不能保证原子性)

类似python的join, 它可以让携程等待所有其它携程结束后执行.

```go
wg := sync.WaitGroup{}

var ch01 = make(chan int32, 1)
var ch02 = make(chan int32, 1)
var ch03 = make(chan int32, 1)

wg.Add(3)

// 开启携程, 并传入只读通道
go func(ch chan<- int32) {
    defer wg.Done()
    defer close(ch)
    time.Sleep(3 * time.Second)
    ch <- 1
}(ch01)

go func(ch chan<- int32) {
    defer wg.Done()
    defer close(ch)
    time.Sleep(5 * time.Second)
    ch <- 2
}(ch02)

go func(ch chan<- int32) {
    defer wg.Done()
    defer close(ch)
    time.Sleep(8 * time.Second)
    ch <- 3
}(ch03)

// 等待通道有值, 否则阻塞
wg.Wait()
select {
    case val01 := <-ch01:
    fmt.Print(val01)
    case val02 := <-ch02:
    fmt.Print(val02)
    case val03 := <-ch03:
    fmt.Print(val03)
}
```

### 7.3.3 mutex.lock上锁(互斥锁)

```go
package main

import (
    "fmt"
    "sync"
)

func main() {
	// 锁
	mx := sync.Mutex{}
	wg := sync.WaitGroup{}
	wg.Add(10)
	ii := 0
	for range 10 {
		go func() {
			defer wg.Done()
			mx.Lock()
			defer mx.Unlock()
			for j := 0; j < 1000; j++ {
				ii++
			}
		}()
	}

	wg.Wait()
	fmt.Println(ii)
}
```

### 7.3.4 RWMutex读写锁

* Lock/Unlock：针对写操作。
  不管锁是被reader还是writer持有，这个Lock方法会一直阻塞，Unlock用来释放锁的方法
* RLock/RUnlock：针对读操作
  当锁被reader所有的时候，RLock会直接返回，当锁已经被writer所有，RLock会一直阻塞，直到能获取锁，否则就直接返回，RUnlock用来释放锁的方法

### 7.3.5 sync.Once

确保传入的函数只执行一次

```go
once := sync.Once{}
for range 10 {
    go func() {
        time.Sleep(5 * time.Second)
        once.Do(func() { fmt.Println("hello do") })
    }()
}

// 主线程必须慢于子线程结束
time.Sleep(10 * time.Second)
```

### 7.3.6 context

* context很像react中的createContext.

context作用是**传递控制信号(超时, 取消)**, context可以同时透传给多个协程.

**协程是扁平的, 而context能将多个协程组成成树形结构**, 你完全可以利用通道来取代context的作用. 只是context更符合语义化, 更方便. 通道应该更关注数据的同步, 而不是命令传递.

```go
// Context
ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
defer cancel()
go func() {
    for range 10 {
        time.Sleep(time.Second)
        select {
            case <-ctx.Done():
            fmt.Println(ctx.Err())
            return
            case t := <-time.After(1):
            fmt.Println(t)
        }
    }
}()
time.Sleep(6 * time.Second)
```

当然, 也可以使用context进行传值(当然, 要谨慎使用)

```go
v := context.WithValue(context.Background(), "key", "hello")
v.Value("key")
```

* context使用的注意事项

  context在协程间传输一定要小心, 不要让外层协程提前关闭ctx.

  ```go
  ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
  go func() {
      for range 10 {
          time.Sleep(time.Second)
          // ctx02, cancel02 := context.WithTimeout(context.Background(), 5*time.Second) 可以在这里新建context.
          select {
              case <-ctx.Done():
              fmt.Println(ctx.Err())
              return
              case t := <-time.After(1):
              fmt.Println(t)
          }
      }
  }()
  cancel()
  time.Sleep(6 * time.Second)
  ```

### 7.3.7 还有很多

* atomic原子操作, sync.Pool临时对象池(对象随时可能回收, 但可以减少GC压力)

* 信号机制

# 8. 类型转换

## 8.1 常用类型转换

其中strconv是字符串和常见类型之间进行转换的包.

```go
//将整型转换为浮点型：
var a int = 10
var b float64 = float64(a)

// 字符串类型的转换
package main

import (
    "fmt"
    "strconv"
)
func main() {
    str := "123"
    num, err := strconv.Atoi(str)
    if err != nil {
        fmt.Println("转换错误:", err)
    } else {
        fmt.Printf("字符串 '%s' 转换为整数为：%d\n", str, num)
    }
}
```

## 8.2. 接口类型转换
* 类型断言v.(T). 类型断言用于从接口类型中提取其底层值. **不兼容返回err.** 

```go
type s struct {
}

func (s s) methodA() {
	fmt.Println("hello method")
}

func main() {
	var i1 interface{} = s{}
	ii1 := i1.(s)
	ii1.methodA()
}
```

## 8.3 获取接口类型

### 8.3.1 **特殊语法v.(type).** 获取v变量的类型, 而不是进行断言.

```go
// 类型选择
func f (v interface{}){
  t := v.(type)
  switch(t){
    case int: break;
    case float64: break;
    case StructOne: break;
  }
}

// 接口组合
type Reader interface {
  Read() string
}
type Writer interface {
  Write(data string)
}
type ReadWriter interface {
  Reader
  Writer
}
```

### 8.3.2 reflect.TypeOf

```go
var i interface{} = 42
fmt.Println(reflect.TypeOf(i))
```

# 9. 练习

## 9.1 依赖注入尝试 (类型不能作为键, reflect.TypeOf的可以)

无奈使用字符串

```go
var i interface{} = 42
map01 := make(map[any]any, 5)
map01["key"] = "value"
// map01[i.(type)] = i  报错
map01[reflect.TypeOf(i).String()] = i
fmt.Println(map01["int"])
```

## 9.2 连接数据库

* 下载mysql驱动

```go
go get -u github.com/go-sql-driver/mysql
```

* 导入驱动和sql包

```go
import (
    "database/sql"
    "time"
    _ "github.com/go-sql-driver/mysql"
)
```

* 使用

```sql
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
```

具体参考go的官网吧. 因为sql是标准库

如果想练习socket编程, 可以参考标准库的net模块

# 10. 附录

## 10.1 标准库地址

[Standard library - Go Packages](https://pkg.go.dev/std)

# 11. 小结

1. **简洁高效的语法**
   - 类 C 语法易上手，强制代码格式统一（`gofmt`），减少风格争议。
   - 编译型语言，直接生成机器码，执行效率接近 C/C++（无虚拟机开销）。
2. **原生并发模型**
   - `goroutine`（轻量级线程，内存占用约 2KB）和 `channel` 简化高并发编程，对比 Java/ Python 的线程模型更高效。
   - 典型用例：单机可轻松支撑百万级并发（如消息队列服务）。
3. **编译与跨平台**
   - 编译速度快（依赖分析优化），适合快速迭代。
   - 支持交叉编译，单命令生成多平台二进制文件（如 `GOOS=linux GOARCH=amd64 go build`）。
4. **内置工具链**
   - 自带测试、性能分析（`pprof`）、文档生成（`godoc`）等工具，开箱即用。
5. **内存管理**
   - 自动垃圾回收（GC），经版本优化后延迟通常低于 1ms（Go 1.14+ 引入并发标记改进）。
6. **标准库与生态**
   - 网络、加密、JSON 等标准库丰富，适合 Web 后端开发。
   - 云原生主流工具均用 Go 编写（Docker、Kubernetes、Prometheus 等）。

------

### **缺点**

1. **泛型支持较晚**
   - 泛型（1.18 引入）生态成熟度不如 Java/C#，部分库需适配。
2. **错误处理机制**
   - 显式 `if err != nil` 检查冗长，需结合 `errors` 包或第三方库（如 `pkg/errors`）增强可读性。
3. **包管理历史问题**
   - 早期依赖管理混乱，现由 `Go Modules`（1.11+）解决，但部分旧项目仍需迁移。
4. **灵活性与限制**
   - 无继承和运算符重载，依赖接口和组合，对 OOP 习惯的开发者需适应。
   - 不支持动态加载代码（如插件系统受限）。
5. **领域局限性**
   - GUI、移动端、嵌入式等领域生态较弱（如较少的机器学习库）。
6. **GC 暂停问题**
   - 虽经优化，仍不适合硬实时系统（如高频交易场景）。

# 12. 更多

* 泛型[后端 - Go 1.18 泛型全面讲解：一篇讲清泛型的全部 - 个人文章 - SegmentFault 思否](https://segmentfault.com/a/1190000041634906)