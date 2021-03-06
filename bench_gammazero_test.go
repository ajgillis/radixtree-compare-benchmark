package triebench

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"testing"

	"github.com/gammazero/radixtree"
)

const (
	wordsPath = "test_data/words"
	web2aPath = "test_data/web2a"
	uuidsPath = "test_data/uuids.txt"
	hskPath   = "test_data/hsk_words.txt"
)

//
// Benchmarks
//
func BenchmarkGammazeroWordsPut(b *testing.B) {
	benchmarkPut(wordsPath, b)
}

func BenchmarkGammazeroWordsGet(b *testing.B) {
	benchmarkGet(wordsPath, b)
}

func BenchmarkGammazeroWordsWalk(b *testing.B) {
	benchmarkWalk(wordsPath, b)
}

func BenchmarkGammazeroWordsWalkPath(b *testing.B) {
	benchmarkWalkPath(wordsPath, b)
}

func BenchmarkGammazeroWeb2aPut(b *testing.B) {
	benchmarkPut(web2aPath, b)
}

func BenchmarkGammazeroWeb2aGet(b *testing.B) {
	benchmarkGet(web2aPath, b)
}

func BenchmarkGammazeroWeb2aWalk(b *testing.B) {
	benchmarkWalk(web2aPath, b)
}

func BenchmarkGammazeroWeb2aWalkPath(b *testing.B) {
	benchmarkWalkPath(web2aPath, b)
}

func BenchmarkGammazeroUUIDsPut(b *testing.B) {
	benchmarkPut(uuidsPath, b)
}

func BenchmarkGammazeroUUIDsGet(b *testing.B) {
	benchmarkGet(uuidsPath, b)
}

func BenchmarkGammazeroUUIDsWalk(b *testing.B) {
	benchmarkWalk(uuidsPath, b)
}

func BenchmarkGammazeroUUIDsWalkPath(b *testing.B) {
	benchmarkWalkPath(uuidsPath, b)
}

func BenchmarkGammazeroHSKPut(b *testing.B) {
	benchmarkPut(hskPath, b)
}

func BenchmarkGammazeroHSKGet(b *testing.B) {
	benchmarkGet(hskPath, b)
}

func BenchmarkGammazeroHSKWalk(b *testing.B) {
	benchmarkWalk(hskPath, b)
}

func BenchmarkGammazeroHSKWalkPath(b *testing.B) {
	benchmarkWalkPath(hskPath, b)
}

func BenchmarkGammazeroPutWithExisting(b *testing.B) {
	tree := radixtree.New()
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

func benchmarkPut(filePath string, b *testing.B) {
	words := loadWords(filePath)
	b.ResetTimer()
	b.ReportAllocs()
	for n := 0; n < b.N; n++ {
		tree := radixtree.New()
		for _, w := range words {
			tree.Put(w, w)
		}
	}
}

func benchmarkGet(filePath string, b *testing.B) {
	words := loadWords(filePath)
	tree := radixtree.New()
	for _, w := range words {
		tree.Put(w, w)
	}
	b.ResetTimer()
	b.ReportAllocs()
	for n := 0; n < b.N; n++ {
		for _, w := range words {
			tree.Get(w)
		}
	}
}

func benchmarkWalk(filePath string, b *testing.B) {
	words := loadWords(filePath)
	tree := radixtree.New()
	for _, w := range words {
		tree.Put(w, w)
	}
	b.ResetTimer()
	b.ReportAllocs()
	var count int
	for n := 0; n < b.N; n++ {
		count = 0
		tree.Walk("", func(key string, value interface{}) bool {
			count++
			return false
		})
	}
	if count != len(words) {
		panic("wrong count")
	}
}

func benchmarkWalkPath(filePath string, b *testing.B) {
	words := loadWords(filePath)
	tree := radixtree.New()
	for _, w := range words {
		tree.Put(w, w)
	}
	b.ResetTimer()
	b.ReportAllocs()
	var count int
	for n := 0; n < b.N; n++ {
		count = 0
		for _, w := range words {
			tree.WalkPath(w, func(key string, value interface{}) bool {
				count++
				return false
			})
		}
	}
	if count < len(words) {
		panic("wrong count")
	}
}

func loadWords(wordsFile string) []string {
	f, err := os.Open(wordsFile)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	var word string
	var words []string

	// Scan through line-dilimited words.
	for scanner.Scan() {
		word = scanner.Text()
		words = append(words, word)
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	if len(words) == 0 {
		panic("did not load words")
	}

	return words
}
