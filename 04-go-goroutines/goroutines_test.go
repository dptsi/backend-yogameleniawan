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
