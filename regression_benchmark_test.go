package main

import (
	"os"
	"testing"
)

func SuppressOutput() {
	os.Stdout, _ = os.OpenFile(os.DevNull, os.O_WRONLY, 0)
}

func BenchmarkSerialRegression(b *testing.B) {
	SuppressOutput()

	for i := 0; i < 100; i++ {
		SerialRegression()
	}
}

func BenchmarkConcurrentRegression(b *testing.B) {
	SuppressOutput()

	for i := 0; i < 100; i++ {
		ConcurrentRegression()
	}
}

