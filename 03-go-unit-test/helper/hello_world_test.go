package helper

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMain(m *testing.M) {
	fmt.Println("Start Testing Hello World")
	m.Run()
	fmt.Println("End Testing Hello World")
}

func TestSubTest(t *testing.T) {
	t.Run("Yoga", func(t *testing.T) {
		result := HelloWorld("Yoga Meleniawan Pamungkas")
		require.Equal(t, "Hello Yoga Meleniawan Pamungkas", result, "The two words should be the same.")
	})

	t.Run("Meleniawan", func(t *testing.T) {
		result := HelloWorld("Meleniawan")
		require.Equal(t, "Hello Yoga Meleniawan Pamungkas", result, "The two words should be the same.")
	})
}

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

func TestHelloWorldTable(t *testing.T) {
	tests := []struct {
		name     string
		request  string
		expected string
	}{
		{
			name:     "Yoga",
			request:  "Yoga Meleniawan Pamungkas",
			expected: "Hello Yoga Meleniawan Pamungkas",
		},
		{
			name:     "Meleniawan",
			request:  "Yoga Meleniawan Pamungkas",
			expected: "Hello Yoga Meleniawan Pamungkas",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := HelloWorld(test.request)
			assert.Equal(t, test.expected, result, "The two words should be the same.")
		})
	}
}
