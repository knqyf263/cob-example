package main

import (
    "fmt"
    "testing"
)

func BenchmarkAppend_Allocate(b *testing.B) {
    base := make([]string, b.N)
    b.ResetTimer()
    for i := 0; i < b.N; i++ {
        base[i] = fmt.Sprintf("no%d", i)
    }
}

type Book struct {
	Title string
}

func Call(b *Book){}

func BenchmarkCall(b *testing.B) {
	book := Book{"Golang"}
	for i := 0; i < b.N; i++ {
		Call(&book)
	}
}
