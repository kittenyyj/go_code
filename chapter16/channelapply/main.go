package main

import (
	"fmt"
	"time"
)

//write Data
func writeData(intChan chan int) {
	for i := 1; i <= 50; i++ {
		//放入数据
		intChan <- i //
		fmt.Println("writeData ", i)
		//time.Sleep(time.Second)
	}
	close(intChan) //关闭
}

//read data
func readData(intChan chan int, exitChan chan bool) int {

	for {
		v, ok := <-intChan
		if !ok {
			break
		}
		//time.Sleep(time.Second)
		fmt.Printf("readData 读到数据=%v\n", v)
	}
	//readData 读取完数据后，即任务完成
	time.Sleep(time.Second)
	//exitChan <- true
	close(exitChan)
	return 1
}

func main() {

	//创建两个管道
	intChan := make(chan int, 10)
	exitChan := make(chan bool, 1)

	go writeData(intChan)
	go readData(intChan, exitChan)

	//time.Sleep(time.Second * 10)

	k, v := <-exitChan
	fmt.Println(k, v)

}
