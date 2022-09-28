package waitWay

import (
	"fmt"
	"time"
)

func UseChannel() {
	ch := make(chan string)
	go sayTakeCh("Hello ch", ch)

	<-ch
}

func sayTakeCh(s string, c chan string) {
	for i := 0; i < 5; i++ {
		time.Sleep(1000 * time.Millisecond)
		fmt.Println(s)
	}

	c <- "FINISH"
}
