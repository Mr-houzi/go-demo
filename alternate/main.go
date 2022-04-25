package main

import (
	"fmt"
	"strings"
	"sync"
)

//交替打印数字和字⺟
//问题描述
//使⽤两个  goroutine 交替打印序列，⼀个  goroutine 打印数字， 另外⼀
//个  goroutine 打印字⺟， 最终效果如下：
//解题思路
//问题很简单，使⽤ channel 来控制打印的进度。使⽤两个 channel ，来分别控制数字和
//字⺟的打印序列， 数字打印完成后通过 channel 通知字⺟打印, 字⺟打印完成后通知数
//字打印，然后周⽽复始的⼯作。

func main()  {
	letterchan, numberchan := make(chan bool), make(chan bool)
	go func() {
		i := 1
		for  {
			select {
			case <-numberchan:
				fmt.Print(i)
				i++
				fmt.Print(i)
				i++
				letterchan <- true
			default:
				break
			}
		}
	}()

	wait := sync.WaitGroup{}
	wait.Add(1)
	go func(wait *sync.WaitGroup) {
		str := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
		i := 0
		for {
			select {
			case <-letterchan:
				if i >= strings.Count(str, "")-1 {
					wait.Done()
					return
				}
				
				fmt.Print(str[i:i+1])
				i++
				fmt.Print(str[i:i+1])
				i++
				numberchan <- true
			default:
				break
			}
		}
	}(&wait)

	numberchan <- true

	wait.Wait()
}