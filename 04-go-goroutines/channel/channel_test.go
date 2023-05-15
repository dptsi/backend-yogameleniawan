package channel

import (
	"fmt"
	"testing"
	"time"
)

func TestCreateChannel(t *testing.T) {
	channel := make(chan string)

	go func() {
		time.Sleep(2 * time.Second)
		channel <- "Hello World"
		fmt.Println("Finishing goroutine")
	}()

	data := <-channel
	fmt.Println(data)

	time.Sleep(5 * time.Second)
}
