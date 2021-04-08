package concurrency

import (
	"bytes"
	"fmt"
	"math/rand"
	"time"
)

func announce(message string,delay time.Duration)  {
	c := make(chan int)
	go func() {
		time.Sleep(delay)
		fmt.Println(message)
		// fmt.Println(<-c )
		// c <- 10
		// c <- 20
	}()
	c <- 10
	// fmt.Println(<- c)

	// fmt.Println(<- c)
}

func test01() {
	var sem = make(chan int,10)
	var queue = make(chan int,5)
	var handle = func (r int) {
		sem <- 1
		time.Sleep(time.Duration(rand.Intn(2)) * time.Second)
		fmt.Println("receiver:",r)
		// make something
		<- sem
	}
	go func() {
		for i := 1;i < 100;i ++ {
			queue <- i * 10
		}
		close(queue)
	}()
	for r := range queue {
		go handle(r)
	}
}
func test02() {
	var freeList = make(chan *bytes.Buffer, 100)
	var serveChan = make(chan *bytes.Buffer)
	var client = func() {
		for {
			var b *bytes.Buffer
			select {
				case b = <- freeList:
					fmt.Println("获取到资源.")
				default:
					b = new(bytes.Buffer)
			}
			//
			serveChan <- b
		}
	}

	var serve = func (){
	}
	serve()
	client()


}
func Do() {
	// test02()
	// test01()
	// announce("whx",time.Second / 2)
}