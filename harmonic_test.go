package harmonic

import (
	"fmt"

	"math/rand"
	"strconv"
	"testing"

	"os"
	"sort"

	"github.com/google/btree"
)

const testCount = 10000

var (
	sorted10      = getSorted(10)
	sorted100     = getSorted(100)
	sorted1000    = getSorted(1000)
	sorted10000   = getSorted(10000)
	sorted100000  = getSorted(100000)
	sorted1000000 = getSorted(1000000)
)

var sorted = getSorted(testCount)
var reversed = getReversed(sorted)
var randomized = getRandomized(sorted)

var testVal int

//var testHarmonic *Harmonic
var testHarmonic = populateHarmonic(sorted)
var testBTree = populateBtree(sorted)
var testMap = populateMap(sorted)

func TestMain(m *testing.M) {
	m.Run()
	os.Exit(0)
}

func TestGet(t *testing.T) {
	fmt.Println(testHarmonic.Get("9999"))
}

func BenchmarkHarmonicSortedGet_10(b *testing.B) {
	benchmarkHarmonicGet(b, testHarmonic, sorted10)
}

func BenchmarkHarmonicSortedGet_100(b *testing.B) {
	benchmarkHarmonicGet(b, testHarmonic, sorted100)
}

func BenchmarkHarmonicSortedGet_1000(b *testing.B) {
	benchmarkHarmonicGet(b, testHarmonic, sorted1000)
}

func BenchmarkHarmonicSortedGet_10000(b *testing.B) {
	benchmarkHarmonicGet(b, testHarmonic, sorted10000)
}

func BenchmarkHarmonicSortedGet_100000(b *testing.B) {
	benchmarkHarmonicGet(b, testHarmonic, sorted100000)
}

func BenchmarkHarmonicSortedGet_1000000(b *testing.B) {
	benchmarkHarmonicGet(b, testHarmonic, sorted1000000)
}

func BenchmarkHarmonicSortedPut(b *testing.B) {
	benchmarkHarmonicPut(b, sorted)
}

func BenchmarkBtreeSortedGet(b *testing.B) {
	benchmarkBtreeGet(b, testBTree)
}

func BenchmarkBtreeSortedPut(b *testing.B) {
	benchmarkBtreePut(b, sorted)
}

func BenchmarkMapSortedGet_10(b *testing.B) {
	benchmarkMapGet(b, testMap, sorted10)
}

func BenchmarkMapSortedGet_100(b *testing.B) {
	benchmarkMapGet(b, testMap, sorted100)
}

func BenchmarkMapSortedGet_1000(b *testing.B) {
	benchmarkMapGet(b, testMap, sorted1000)
}

func BenchmarkMapSortedGet_10000(b *testing.B) {
	benchmarkMapGet(b, testMap, sorted10000)
}

func BenchmarkMapSortedGet_100000(b *testing.B) {
	benchmarkMapGet(b, testMap, sorted100000)
}

func BenchmarkMapSortedGet_1000000(b *testing.B) {
	benchmarkMapGet(b, testMap, sorted1000000)
}

func BenchmarkMapSortedPut(b *testing.B) {
	benchmarkMapPut(b, sorted)
}

func BenchmarkHarmonicReversePut(b *testing.B) {
	benchmarkHarmonicPut(b, reversed)
}

func BenchmarkBtreeReversePut(b *testing.B) {
	benchmarkBtreePut(b, reversed)
}

func BenchmarkMapReversePut(b *testing.B) {
	benchmarkMapPut(b, reversed)
}

func BenchmarkHarmonicRandPut(b *testing.B) {
	benchmarkHarmonicPut(b, randomized)
}

func BenchmarkBtreeRandPut(b *testing.B) {
	benchmarkBtreePut(b, randomized)
}

func BenchmarkMapRandPut(b *testing.B) {
	benchmarkMapPut(b, randomized)
}

func benchmarkHarmonicGet(b *testing.B, h *Harmonic, list []string) {
	for i := 0; i < b.N; i++ {
		for j := 0; j < len(list); j++ {
			testVal, _ = h.Get(list[j])
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

func benchmarkMapGet(b *testing.B, m map[string]int, list []string) {
	for i := 0; i < b.N; i++ {
		for j := 0; j < len(list); j++ {
			testVal = m[list[j]]
		}
	}

	b.ReportAllocs()
}

func benchmarkHarmonicPut(b *testing.B, list []string) {
	for i := 0; i < b.N; i++ {
		h := New(testCount)

		for j := 0; j < testCount; j++ {
			h.Put(list[j], j)
		}
	}

	b.ReportAllocs()
}

func benchmarkBtreePut(b *testing.B, list []string) {
	for i := 0; i < b.N; i++ {
		bt := btree.New(2)

		for j := 0; j < testCount; j++ {
			bt.ReplaceOrInsert(btreeItem{list[j], j})
		}
	}

	b.ReportAllocs()
}

func benchmarkMapPut(b *testing.B, list []string) {
	for i := 0; i < b.N; i++ {
		m := make(map[string]int, testCount)

		for j := 0; j < testCount; j++ {
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

func populateHarmonic(list []string) (h *Harmonic) {
	h = New(len(list))
	for i := 0; i < len(list); i++ {
		h.Put(list[i], i)
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
