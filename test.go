package main

import (

	"io/ioutil"
	"testing"
)

func BenchmarkGG(b *testing.B) {
	for i := 0; i < b.N; i++ {
		GG(ioutil.Discard)
	}
}

//func BenchmarkSolution(b *testing.B) {
//	for i := 0; i < b.N; i++ {
//		Solution(ioutil.Discard)
//	}
//}