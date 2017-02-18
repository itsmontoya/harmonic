package seeker

import (
	"fmt"

	"math/rand"
	"strconv"
	"testing"

	"github.com/google/btree"
	"os"
	"sort"
)

var sorted = getSorted(10000)
var reversed = getReversed(sorted)
var randomized = getRandomized(sorted)

var testVal int
var testSeeker = populateSeeker(sorted)
var testBTree = populateBtree(sorted)
var testMap = populateMap(sorted)

func TestMain(m *testing.M) {
	testSeeker = New(10000)
	m.Run()
	os.Exit(0)
}

func TestGet(t *testing.T) {
	s := New(100)
	s.Put("a", 1)
	s.Put("b", 2)
	s.Put("c", 3)
	fmt.Println(s.Get("b"))
}

func BenchmarkSeekerSortedGet(b *testing.B) {
	benchmarkSeekerGet(b, testSeeker)
}

func BenchmarkSeekerSortedPut(b *testing.B) {
	benchmarkSeekerPut(b, sorted)
}

func BenchmarkBtreeSortedGet(b *testing.B) {
	benchmarkBtreeGet(b, testBTree)
}

func BenchmarkBtreeSortedPut(b *testing.B) {
	benchmarkBtreePut(b, sorted)
}

func BenchmarkMapSortedGet(b *testing.B) {
	benchmarkMapGet(b, testMap)
}

func BenchmarkMapSortedPut(b *testing.B) {
	benchmarkMapPut(b, sorted)
}

func BenchmarkSeekerReversePut(b *testing.B) {
	benchmarkSeekerPut(b, reversed)
}

func BenchmarkBtreeReversePut(b *testing.B) {
	benchmarkBtreePut(b, reversed)
}

func BenchmarkMapReversePut(b *testing.B) {
	benchmarkMapPut(b, reversed)
}

func BenchmarkSeekerRandPut(b *testing.B) {
	benchmarkSeekerPut(b, randomized)
}

func BenchmarkBtreeRandPut(b *testing.B) {
	benchmarkBtreePut(b, randomized)
}

func BenchmarkMapRandPut(b *testing.B) {
	benchmarkMapPut(b, randomized)
}

func benchmarkSeekerGet(b *testing.B, s *Seeker) {
	for i := 0; i < b.N; i++ {
		for j := 0; j < len(sorted); j++ {
			testVal, _ = s.Get(sorted[j])
		}
	}

	b.ReportAllocs()
}

func benchmarkBtreeGet(b *testing.B, bt *btree.BTree) {
	for i := 0; i < b.N; i++ {
		for j := 0; j < len(sorted); j++ {
			item := bt.Get(btreeItem{key: sorted[j]}).(btreeItem)
			testVal = item.val
		}
	}

	b.ReportAllocs()
}

func benchmarkMapGet(b *testing.B, m map[string]int) {
	for i := 0; i < b.N; i++ {
		for j := 0; j < len(sorted); j++ {
			testVal = m[sorted[j]]
		}
	}

	b.ReportAllocs()
}

func benchmarkSeekerPut(b *testing.B, list []string) {
	for i := 0; i < b.N; i++ {
		s := New(10000)

		for j := 0; j < 10000; j++ {
			s.Put(list[j], j)
		}
	}

	b.ReportAllocs()
}

func benchmarkBtreePut(b *testing.B, list []string) {
	for i := 0; i < b.N; i++ {
		bt := btree.New(2)

		for j := 0; j < 10000; j++ {
			bt.ReplaceOrInsert(btreeItem{list[j], j})
		}
	}

	b.ReportAllocs()
}

func benchmarkMapPut(b *testing.B, list []string) {
	for i := 0; i < b.N; i++ {
		m := make(map[string]int, 10000)

		for j := 0; j < 10000; j++ {
			m[list[j]] = i
		}
	}

	b.ReportAllocs()
}

func getSorted(n int) (s []string) {
	for i := 0; i < n; i++ {
		s = append(s, strconv.Itoa(i))
	}

	sort.Strings(s)
	return
}

func getReversed(sorted []string) (rev []string) {
	rev = make([]string, len(sorted))
	cap := len(sorted)
	for i := 0; i < cap; i++ {
		j := cap - i - 1
		rev[j] = sorted[i]
	}

	return
}

func getRandomized(sorted []string) (unsorted []string) {
	unsorted = make([]string, len(sorted))
	perm := rand.Perm(len(sorted))
	for i, v := range perm {
		unsorted[i] = sorted[v]
	}

	return
}

type btreeItem struct {
	key string
	val int
}

func (bi btreeItem) Less(item btree.Item) bool {
	nbi := item.(btreeItem)
	return nbi.key < bi.key
}

func populateSeeker(list []string) (s *Seeker) {
	s = New(len(list))
	for i := 0; i < len(list); i++ {
		s.Put(list[i], i)
	}

	return
}

func populateBtree(list []string) (bt *btree.BTree) {
	bt = btree.New(2)
	for i := 0; i < len(list); i++ {
		bt.ReplaceOrInsert(btreeItem{list[i], i})
	}

	return
}

func populateMap(list []string) (m map[string]int) {
	m = make(map[string]int, len(list))
	for i := 0; i < len(list); i++ {
		m[list[i]] = i
	}

	return
}
