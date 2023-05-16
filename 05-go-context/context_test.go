package go_context

import (
	"context"
	"fmt"
	"runtime"
	"testing"
	"time"
)

func TestContext(t *testing.T) {
	background := context.Background()
	fmt.Println(background)

	todo := context.TODO()
	fmt.Println(todo)
}

func TestContextWithValue(t *testing.T) {
	contextA := context.Background()

	contextB := context.WithValue(contextA, "b", "B")
	contextC := context.WithValue(contextA, "c", "C")

	contextD := context.WithValue(contextB, "d", "D")
	contextE := context.WithValue(contextB, "e", "E")

	contextF := context.WithValue(contextC, "f", "F")

	fmt.Println(contextA)
	fmt.Println(contextB)
	fmt.Println(contextC)
	fmt.Println(contextD)
	fmt.Println(contextE)
	fmt.Println(contextF)

	fmt.Println(contextF.Value("f"))
	fmt.Println(contextF.Value("c"))
	fmt.Println(contextF.Value("b"))
	fmt.Println(contextF.Value("a"))

	fmt.Println(contextE.Value("e"))
	fmt.Println(contextE.Value("b"))
	fmt.Println(contextE.Value("a"))

	fmt.Println(contextD.Value("d"))
	fmt.Println(contextD.Value("b"))
	fmt.Println(contextD.Value("a"))

	fmt.Println(contextC.Value("c"))
	fmt.Println(contextC.Value("a"))

	fmt.Println(contextB.Value("b"))
	fmt.Println(contextB.Value("a"))

	fmt.Println(contextA.Value("a"))
}

func CreateCounter(ctx context.Context) chan int {
	desination := make(chan int)

	go func() {
		defer close(desination)
		counter := 1
		for {
			select {
			case <-ctx.Done():
				return
			default:
				desination <- counter
				counter++
			}
		}
	}()

	return desination
}

func TestContextWithCancel(t *testing.T) {
	fmt.Println("Total Goroutine", runtime.NumGoroutine())
	parent := context.Background()
	ctx, cancel := context.WithCancel(parent)

	destination := CreateCounter(ctx)

	fmt.Println("Total Goroutine", runtime.NumGoroutine())

	for n := range destination {
		fmt.Println("Counter", n)

		if n == 10 {
			break
		}
	}

	cancel()

	time.Sleep(2 * time.Second)

	fmt.Println("Total Goroutine", runtime.NumGoroutine())
}
