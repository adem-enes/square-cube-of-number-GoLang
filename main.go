package main

import (
	"fmt"
	"sync"
)

/*
=> 1 ile 10 arasındaki sayıların karesini ve küpünü eş zamanlı hesaplayacak goroutine fonksiyonları
=> sonrasında tüm çıktıları bir channel aracılığı ile toplayıp yazdıracak ek bir fonksiyon yazınız
*/

// var wg sync.WaitGroup

func main() {
	wg := new(sync.WaitGroup)
	sqChannel := make(chan int, 10)
	cbChannel := make(chan int, 10)
	// cubeChannel := make(chan int)

	for i := 1; i <= 10; i++ {
		go square(i, wg, sqChannel)
		go cube(i, wg, cbChannel)
		wg.Wait()
	}

	// for v := range sqChannel {
	// 	fmt.Println("sq:", v)
	// }
	// close(sqChannel)
	writeChannel(sqChannel)
	// fmt.Println("sq: ", len(sqChannel))
	fmt.Println("sq: ", cap(sqChannel))
	fmt.Println("----------------")
	// fmt.Println("cb: ", len(cbChannel))
	fmt.Println("cb: ", cap(cbChannel))
	writeChannel(cbChannel)
	// for v := range cbChannel {
	// 	fmt.Println("sq:", v)
	// }
}

func square(number int, wg *sync.WaitGroup, chan1 chan int) {
	wg.Add(1)

	sq := number * number
	chan1 <- sq
	// fmt.Println(cb)

	wg.Done()

}

func cube(number int, wg *sync.WaitGroup, chan1 chan int) {
	wg.Add(1)
	cb := number * number * number
	chan1 <- cb
	// fmt.Println(cb)

	wg.Done()
}

func writeChannel(channel chan int) {
	// for v := range <-channel {
	// 	fmt.Println(v)
	// }

	fmt.Println(<-channel)
}
