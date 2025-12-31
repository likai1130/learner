package zuoyebang

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"testing"
	"time"
)

/*
*
输入一个数，输出小于它的所有素数 go实现
除了1和它本身整除
*/
func TestPrime(t *testing.T) {
	//var num int
	f2(10)
}

func isPrime(num int) bool {
	//只能被1和本身整除
	for i := 2; i*i <= num; i++ {
		if num%i == 0 {
			return false
		}
	}
	return true
}

func primesLessThan(n int) []int {
	primes := make([]int, 0)
	if n == 1 {
		return primes
	}
	for i := 0; i < n; i++ {
		if isPrime(i) {
			primes = append(primes, i)
		}
	}
	return primes
}

func f2(num int) {
	than := primesLessThan(num)
	fmt.Printf("小于%d所有素数是%v\n", num, than)
}

/***********************************************/

/*
*
开根号 go 实现 不使用系统函数
*/
func TestSqrtNewtonRaphson(t *testing.T) {
	num := 16.0
	sqrtNum := Sqrt(num) // 初始猜测值设为num的一半
	fmt.Printf("The square root of %.2f is approximately %.6f\n", num, sqrtNum)
}

func Sqrt(x float64) float64 {
	if x == 0 {
		return 0
	}
	z := 1.0
	for i := 0; i < 10; i++ {
		z = z - ((z*z - x) / (2 * z))
	}
	return z
}

/*****************/
/**
go 实现等待10个协程完成任务，并获取每个协程的结果
*/
func TestGoroutine(t *testing.T) {
	workNum := 10
	var wg sync.WaitGroup
	workResultCh := make(chan string, workNum)

	for i := 0; i < workNum; i++ {
		wg.Add(1)
		go worker(workResultCh, &wg, i)
	}
	wg.Wait()
	close(workResultCh)
	for ch := range workResultCh {
		fmt.Println(ch)
	}
	fmt.Println("Done!!!")

}

func worker(result chan string, wg *sync.WaitGroup, i int) {
	defer wg.Done()
	result <- fmt.Sprintf("第%d协程完成任务", i)
}

/*********/

/*
*
一个server，多个client
server发送一个数据，所有client都能收到
任意client发送数据，server都能收到
请用channel，goroutine实现
*/
type Client struct {
	Id  int
	Msg chan string
}

func TestCS(t *testing.T) {
	// 注册接收信号的 channel
	ctx, cancelFunc := context.WithCancel(context.Background())
	clients := make([]Client, 0, 3)
	ch := make(chan string)
	var wg sync.WaitGroup
	wg.Add(1)
	go server(ctx, &wg, ch, &clients)
	for i := 1; i <= 3; i++ {
		wg.Add(1)
		cliCh := Client{
			Id:  i,
			Msg: make(chan string),
		}
		clients = append(clients, cliCh)
		go client(ctx, &cliCh, &wg, ch, fmt.Sprintf("我是客户端%d,", i))
	}
	wg.Wait()
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	cancelFunc()
}

func client(ctx context.Context, cli *Client, wg *sync.WaitGroup, ch chan string, msg string) {
	defer wg.Done()
	// 给server发数据
	ch <- msg
	// 接受server数据
	go func(ctx context.Context) {
		for {
			select {
			case result, ok := <-cli.Msg:
				if ok {
					log.Printf("客户端id=%d,收到信息: %s", cli.Id, result)
				}
			case <-ctx.Done():
				log.Printf("客户端id=%d程序结束\n", cli.Id)
				return
			}
		}
	}(ctx)

}

func server(ctx context.Context, wg *sync.WaitGroup, ch chan string, clients *[]Client) {
	defer wg.Done()

	go func(ctx context.Context) {
		for {
			select {
			case <-ctx.Done():
				log.Println("服务端程序结束")
				return
			case msg, ok := <-ch:
				if ok {
					log.Printf("server端收到: %s \n", msg)
				}
			}
		}
	}(ctx)

	for len(*clients) > 0 {
		for i := 0; i < 3; i++ {
			for _, c := range *clients {
				c.Msg <- fmt.Sprintf("大家好我是服务端")
			}
			time.Sleep(1 * time.Second)
		}
		return
	}
}

/*************/
/**
编程题 golang写个多生产和多消费 （纯纯自由发挥）
*/
/*var (
	queue        = make(chan int, 5) // 创建一个容量为5的缓冲区
	wg           sync.WaitGroup      // 用于等待所有goroutines完成
	numProducers = 2                 // 生产者数量
	numConsumers = 3                 // 消费者数量
)

func TestConsumerAndProducer(t *testing.T) {
	// 启动生产者
	wg.Add(numProducers)
	for i := 0; i < numProducers; i++ {
		go producer(i)
	}

	// 启动消费者
	wg.Add(numConsumers)
	for i := 0; i < numConsumers; i++ {
		go consumer(i)
	}

	// 等待所有goroutines完成
	wg.Wait()
}

// 生产者函数
func producer(id int) {
	defer wg.Done()
	for {
		item := rand.Intn(1000) // 生产一个随机数作为商品
		queue <- item           // 将商品放入队列
		fmt.Printf("Producer %d produced: %d\n", id, item)
		time.Sleep(time.Millisecond * 500) // 生产者休眠一段时间
	}
}

// 消费者函数
func consumer(id int) {
	defer wg.Done()
	for {
		item := <-queue // 从队列中取出商品
		fmt.Printf("Consumer %d consumed: %d\n", id, item)
		time.Sleep(time.Millisecond * 1000) // 消费者休眠一段时间
	}
}*/

// golang 两（三）线程交替打印0-100

func TestPrint(t *testing.T) {
	ch := make(chan int)
	ch2 := make(chan int)
	var wg sync.WaitGroup
	wg.Add(1)
	go goprint(&wg, ch, ch2, 1)
	go goprint(&wg, ch2, ch, 2)
	ch <- 0
	wg.Wait()
	close(ch)
	log.Println("Done!!!")
}

func goprint(wg *sync.WaitGroup, ch, ch2 chan int, i int) {
	defer wg.Done()
	for {
		c := <-ch
		if c > 100 {
			return
		}
		log.Printf("线程%d打印：%d\n", i, c)
		c += 1
		ch2 <- c
		time.Sleep(20 * time.Millisecond)
	}
}

func bubbleSort(arr []int) {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		// 每一轮冒泡将当前最大的元素移动到末尾
		for j := 0; j < n-i-1; j++ {
			if arr[j] > arr[j+1] {
				// 交换相邻元素，使较大的元素向后移动
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
}

func TestName(t *testing.T) {
	// 测试冒泡排序
	arr := []int{64, 34, 25, 12, 22, 11, 90}
	fmt.Println("原始数组:", arr)
	bubbleSort(arr)
	fmt.Println("排序结果:", arr)
}
