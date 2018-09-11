package main

import "testing"

func BenchmarkGreyScaleRoutine(b *testing.B) {
	for n := 0; n < b.N; n++ {
		convert("flower.jpg")
	}
}
