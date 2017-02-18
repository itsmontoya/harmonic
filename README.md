# Seeker
Seeker is a simple sorted flat data store. This is still quite a work in progress, much improvement is yet to be made.

## Benchmarks
```
# Seeker
BenchmarkSeekerSortedGet-4    	     500	   2726619 ns/op	       0 B/op	       0 allocs/op
BenchmarkSeekerSortedPut-4    	    2000	    834277 ns/op	  401921 B/op	   10001 allocs/op
BenchmarkSeekerReversePut-4   	      20	  93772086 ns/op	  401920 B/op	   10001 allocs/op
BenchmarkSeekerRandPut-4      	     100	  15090947 ns/op	  401920 B/op	   10001 allocs/op

# Google's BTree
BenchmarkBtreeSortedGet-4     	     200	   8601598 ns/op	  320000 B/op	   10000 allocs/op
BenchmarkBtreeSortedPut-4     	     100	  13320606 ns/op	 1201064 B/op	   35015 allocs/op
BenchmarkBtreeReversePut-4    	     100	  16879268 ns/op	 2317384 B/op	   59932 allocs/op
BenchmarkBtreeRandPut-4       	     100	  13962494 ns/op	 1223400 B/op	   33175 allocs/op

# Stdlib Hashmap - This is just for a speed reference, map's aren't sorted
BenchmarkMapSortedGet-4       	    3000	    420158 ns/op	       0 B/op	       0 allocs/op
BenchmarkMapSortedPut-4       	    2000	   1026423 ns/op	  451840 B/op	     125 allocs/op
BenchmarkMapReversePut-4      	    2000	   1023698 ns/op	  451838 B/op	     125 allocs/op
BenchmarkMapRandPut-4         	    2000	   1061156 ns/op	  451842 B/op	     125 allocs/op
```