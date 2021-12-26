package main

import (
	"fmt"
	"sync"
)

//(b)
func main1() {
	ch := make(chan bool)
	go func() {
		fmt.Println("下山的路又堵起了")
		ch <- true
	}()
	<-ch
}

//(C)
//var wg sync.WaitGroup

func print1() {
	fmt.Println("一行话")
	wg.Done()
}
func print2() {
	fmt.Println("一行话话")
	wg.Done()
}
func print3() {
	fmt.Println("一行话话")
	wg.Done()
}
func print4() {
	fmt.Println("一行话话")
	wg.Done()
}
func print5() {
	fmt.Println("一行话话")
	wg.Done()
}
func print6() {
	fmt.Println("一行话话")
	wg.Done()
}
func print7() {
	fmt.Println("一行话话")
	wg.Done()
}
func print8() {
	fmt.Println("一行话话")
	wg.Done()
}
func print9() {
	fmt.Println("一行话话")
	wg.Done()
}
func print10() {
	fmt.Println("一行话话")
	wg.Done()
}
func main2() {
	wg.Add(10)
	go print1()
	go print2()
	go print3()
	go print4()
	go print5()
	go print6()
	go print7()
	go print8()
	go print9()
	go print10()
	wg.Wait()
}

//(d)
//四个协程分别均分去查找一百万以内的素数
var wg sync.WaitGroup

func road(n int) {
	for num := (n-1)*250000 + 1; num <= n*250000; num++ {
		if num > 1 {
			var flag = true
			for i := 2; i < num; i++ {
				if num%i == 0 {
					flag = false
					break
				}
			}
			if flag {
				fmt.Println("num是素数，是：", num)
			}
		}
	}
	wg.Done()
}
func main3() {
	for i := 1; i <= 4; i++ {
		wg.Add(1)
		go road(i)
	}
	wg.Wait()
	fmt.Println("执行完毕")
}

//编码能力题

func main() {
	var N, M, P, Q, Z int
	var U = 0
	//var room1 []int
	//var floor1 []int
	var roomsignal []int
	fmt.Scanln(&N, &M) //输入第一行，N表示除了顶层共N层楼，M表示除顶层外每层楼有M个房间
	//for i := 1; i <= N; i++ {
	//	floor1 = append(floor1, i)
	//}
	//for i := 0; i <= M; i++ {
	//	room1 = append(room1, i)
	//}
	for i := 1; i <= N*M; i++ {
		fmt.Scanln(&P, &Q) //输入后边的N*M行，P代表是否有楼梯，Q代表房间指示牌上的数字
		if P == 1 {
			roomsignal = append(roomsignal, Q)
		}
		U = roomsignal[Z+Q] + U //使得U的值为从第一层进入有楼梯的房间的指示牌的数字后上楼层之后加上新号码的和
	}

	fmt.Scanln(&Z) //输入的最后一行，表示从第一层的几号房间开始进入
	fmt.Println(U)
}
