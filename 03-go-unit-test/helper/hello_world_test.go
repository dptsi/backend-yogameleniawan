package helper

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHelloWorld(t *testing.T) {
	result := HelloWorld("Yoga Meleniawan Pamungkas")
	if result != "Hello Yoga Meleniawan Pamungkas" {
		t.Errorf("HelloWorld() = %s; want Hello Yoga Meleniawan Pamungkas", result)
	}
}

func TestHelloWorldOther(t *testing.T) {
	result := HelloWorld("Yoga Meleniawan Pamungkas")
	if result != "Hello Yoga Meleniawan Pamungkas" {
		t.Errorf("HelloWorld() = %s; want Hello Yoga Meleniawan Pamungkas", result)
	}
}

func TestHelloWorldAssertion(t *testing.T) {
	result := HelloWorld("Yoga Meleniawan Pamungkas")

	assert.Equal(t, "Hello Yoga Meleniawan Pamungkas", result, "The two words should be the same.")
	fmt.Println("TestHelloWorldAssertion() done")
}
