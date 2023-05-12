package helper

import "testing"

func TestHelloWorld(t *testing.T) {
	result := HelloWorld("Yoga Meleniawan Pamungkas")
	if result != "Hello Yoga Meleniawan Pamungkas" {
		t.Errorf("HelloWorld() = %s; want Hello Yoga Meleniawan Pamungkas", result)
	}
}
