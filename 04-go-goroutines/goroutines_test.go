package go_routines

import (
	"fmt"
	"testing"
	"time"
)

func RunHelloWorld() {
	fmt.Println("Hello World")
}

func TestCreateGoRoutines(t *testing.T) {
	go RunHelloWorld()
	fmt.Println("Test Done")

	time.Sleep(1 * time.Second)
}

func DisplayNumber(number int) {
	fmt.Println("Dispaly", number)
}

func TestManyGoroutines(t *testing.T) {
	for i := 0; i < 100000; i++ {
		go DisplayNumber(i)
	}

	time.Sleep(5 * time.Second)
}
