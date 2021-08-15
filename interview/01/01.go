package main

import (
	"sync"
)

// 交替打印数字和字符串

var Wait = sync.WaitGroup{}

func main() {
	strCh := make(chan bool, 0)
	numCh := make(chan bool, 0)
	go writeNum(numCh, strCh)
	go writeStr(strCh, numCh)
	numCh <- true
	Wait.Add(1)
	Wait.Wait()
	return
}

func writeStr(ch chan bool, call chan bool) {
	num := 0
	i := 'A'
	for {
		select {
		case <-ch:
			print(string(i + rune(num)))
			num++
			print(string(i + rune(num)))
			num++
			if num == 26 {
				Wait.Done()
			} else {
				call <- true
			}
		}
	}
}

func writeNum(ch chan bool, call chan bool) {
	num := 0
	for {
		select {
		case <-ch:
			print(num + 1)
			num++
			print(num + 1)
			num++
			call <- true
		}
	}
}
