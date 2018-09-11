package main

import "testing"

func BenchmarkGreyScaleFlower(b *testing.B) {
	for n := 0; n < b.N; n++ {
		convert2("flower.jpg")
	}
}
