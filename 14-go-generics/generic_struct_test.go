package belajar_golang_generics

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Data[T any] struct {
	First  T
	Second T
}

func (d *Data[_]) SayHello(name string) string {
	return "Hello " + name
}

func (d *Data[T]) ChangeFirst(first T) T {
	d.First = first
	return d.First
}

func TestData(t *testing.T) {
	data := Data[string]{
		First:  "Yoga",
		Second: "Meleniawan",
	}

	fmt.Println(data)
}

func TestGenericMethod(t *testing.T) {
	data := Data[string]{
		First:  "Yoga",
		Second: "Meleniawan",
	}

	assert.Equal(t, "Pamungkas", data.ChangeFirst("Pamungkas"))
	assert.Equal(t, "Hello Pamungkas", data.SayHello("Pamungkas"))
}
