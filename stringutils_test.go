package stringutils_test

import (
	"testing"

	"github.com/earljohn004/stringutils"
)

func benchStringUtilsToUpper(b *testing.B, n int) {
	for i := 0; i < b.N; i++ {
		stringutils.ToUpper("Hello, World!")
	}
}

func benchStringUtilsToLower(b *testing.B, n int) {
	for i := 0; i < b.N; i++ {
		stringutils.ToLower("Hello, World!")
	}
}

func BenchmarkToUpper10(b *testing.B) {
	benchStringUtilsToUpper(b, 10)
}

func BenchmarkToUpper1000(b *testing.B) {
	benchStringUtilsToUpper(b, 1000)
}

func BenchmarkToLower10(b *testing.B) {
	benchStringUtilsToLower(b, 10)
}

func BenchmarkToLower1000(b *testing.B) {
	benchStringUtilsToLower(b, 1000)
}
