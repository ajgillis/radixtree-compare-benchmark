package triebench

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/dghubble/trie"
)

//
// Benchmarks
//
func BenchmarkDghubbleWordsPut(b *testing.B) {
	benchmarkTriePut(wordsPath, b)
}

func BenchmarkDghubbleWordsGet(b *testing.B) {
	benchmarkTrieGet(wordsPath, b)
}

func BenchmarkDghubbleWordsWalk(b *testing.B) {
	benchmarkTrieWalk(wordsPath, b)
}

func BenchmarkDghubbleWordsWalkPath(b *testing.B) {
	benchmarkTrieWalkPath(wordsPath, b)
}

func BenchmarkDghubbleWeb2aPut(b *testing.B) {
	benchmarkTriePut(web2aPath, b)
}

func BenchmarkDghubbleWeb2aGet(b *testing.B) {
	benchmarkTrieGet(web2aPath, b)
}

func BenchmarkDghubbleWeb2aWalk(b *testing.B) {
	benchmarkTrieWalk(web2aPath, b)
}

func BenchmarkDghubbleWeb2aWalkPath(b *testing.B) {
	benchmarkTrieWalkPath(web2aPath, b)
}

func BenchmarkDghubbleUUIDsPut(b *testing.B) {
	benchmarkTriePut(uuidsPath, b)
}

func BenchmarkDghubbleUUIDsGet(b *testing.B) {
	benchmarkTrieGet(uuidsPath, b)
}

func BenchmarkDghubbleUUIDsWalk(b *testing.B) {
	benchmarkTrieWalk(uuidsPath, b)
}

func BenchmarkDghubbleUUIDsWalkPath(b *testing.B) {
	benchmarkTrieWalkPath(uuidsPath, b)
}

func BenchmarkDghubbleHSKPut(b *testing.B) {
	benchmarkTriePut(hskPath, b)
}

func BenchmarkDghubbleHSKGet(b *testing.B) {
	benchmarkTrieGet(hskPath, b)
}

func BenchmarkDghubbleHSKWalk(b *testing.B) {
	benchmarkTrieWalk(hskPath, b)
}

func BenchmarkDghubbleHSKWalkPath(b *testing.B) {
	benchmarkTrieWalkPath(hskPath, b)
}

func BenchmarkDghubblePutWithExisting(b *testing.B) {
	tree := trie.NewRuneTrie()
	for i := 0; i < 10000; i++ {
		tree.Put(fmt.Sprintf("somekey%d", i), true)
	}
	b.ResetTimer()
	b.ReportAllocs()
	for n := 0; n < b.N; n++ {
		if !tree.Put(strconv.Itoa(n), true) {
			b.Fatal("value was overwritten")
		}
	}
}

func benchmarkTriePut(filePath string, b *testing.B) {
	words := loadWords(filePath)
	b.ResetTimer()
	b.ReportAllocs()
	for n := 0; n < b.N; n++ {
		tree := trie.NewRuneTrie()
		for _, w := range words {
			tree.Put(w, w)
		}
	}
}

func benchmarkTrieGet(filePath string, b *testing.B) {
	words := loadWords(filePath)
	tree := trie.NewRuneTrie()
	for _, w := range words {
		tree.Put(w, w)
	}
	b.ReportAllocs()
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		for _, w := range words {
			tree.Get(w)
		}
	}
}

func benchmarkTrieWalk(filePath string, b *testing.B) {
	words := loadWords(filePath)
	tree := trie.NewRuneTrie()
	for _, w := range words {
		tree.Put(w, w)
	}
	b.ResetTimer()
	b.ReportAllocs()

	var count int
	for n := 0; n < b.N; n++ {
		count = 0
		tree.Walk(func(key string, value interface{}) error {
			count++
			return nil
		})
	}
	if count != len(words) {
		panic("wrong count")
	}
}

func benchmarkTrieWalkPath(filePath string, b *testing.B) {
	words := loadWords(filePath)
	tree := trie.NewRuneTrie()
	for _, w := range words {
		tree.Put(w, w)
	}
	b.ResetTimer()
	b.ReportAllocs()
	var count int
	for n := 0; n < b.N; n++ {
		count = 0
		for _, w := range words {
			tree.WalkPath(w, func(key string, value interface{}) error {
				count++
				return nil
			})
		}
	}
	if count < len(words) {
		panic("wrong count")
	}
}
