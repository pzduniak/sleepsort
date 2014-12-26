# sleepsort

Sleep sort implementation in Go based on a [4chan discussion](http://dis.4chan.org/read/prog/1295544154).
Go is probably the best language to use for implementation of such ridicolous algorithm.

The package exports two methods:

```
sleepsort.IntSlice(input []int, itemDuration time.Duration) []int
sleepsort.FloatSlice(input []float64, itemDuration time.Duration) []float64
```

`input` is the slice to sort. `itemDuration` is used for determining the sleep time. Each goroutine
sleeps for `item*itemDuration` time. Both methods return the sorted slice.
