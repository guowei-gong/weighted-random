package weighted_random

import (
	"fmt"
	"math/rand"
	"sort"
	"strconv"
	"testing"
	"time"
)

func TestWeightedRandomS1(t *testing.T) {
	cnt := make(map[int]int)
	list := []int{1, 2, 4, 8}
	for i := 0; i < 100; i++ {
		v := weightedRandomS1(list)
		cnt[v]++
	}

	for k, v := range cnt {
		fmt.Println(k, v)
	}
}

func weightedRandomS1(weights []int) int {
	if len(weights) == 0 {
		return 0
	}

	var indexList []int

	for i, weight := range weights {
		cnt := 0
		for weight > cnt {
			indexList = append(indexList, i)
			cnt++
		}
	}

	rand.Seed(time.Now().UnixNano())
	return indexList[rand.Intn(len(indexList))]
}

func TestWeightedRandomS2(t *testing.T) {
	cnt := make(map[int]int)
	for i := 0; i < 100; i++ {
		v := weightedRandomS2()
		cnt[v]++
	}

	for k, v := range cnt {
		fmt.Println(k, v)
	}
}

func weightedRandomS2() int {
	rand.Seed(time.Now().UnixNano())
	r := rand.Intn(15)
	if r <= 1 {
		return 0
	} else if 1 < r && r <= 3 {
		return 1
	} else if 3 < r && r <= 7 {
		return 2
	} else {
		return 3
	}
}

func TestWeightedRandomS3(t *testing.T) {
	cnt := make(map[int]int)
	list := []int{1, 2, 4, 8}
	for i := 0; i < 100; i++ {
		v := weightedRandomS3(list)
		cnt[v]++
	}

	for k, v := range cnt {
		fmt.Println(k, v)
	}
}

func weightedRandomS3(weights []int) int {
	rand.Seed(time.Now().UnixNano())
	r := rand.Intn(15)
	for i, v := range weights {
		r = r - v
		if r <= 0 {
			return i
		}
	}
	return len(weights) - 1
}

func TestWeightedRandomS4(t *testing.T) {
	cnt := make(map[int]int)
	list := []int{1, 2, 4, 8}
	for i := 0; i < 100; i++ {
		v := weightedRandomS4(list)
		cnt[v]++
	}

	for k, v := range cnt {
		fmt.Println(k, v)
	}
}

func weightedRandomS4(weights []int) int {
	sort.Sort(sort.Reverse(sort.IntSlice(weights)))
	rand.Seed(time.Now().UnixNano())
	r := rand.Intn(15)
	for i, v := range weights {
		r = r - v
		if r <= 0 {
			return i
		}
	}
	return len(weights) - 1
}

func weightedRandomS5(weights []int) int {
	rand.Seed(time.Now().UnixNano())
	sum := 0
	var sumWeight []int
	for _, v := range weights {
		sum += v
		sumWeight = append(sumWeight, sum)
	}
	r := rand.Intn(sum)
	idx := sort.SearchInts(sumWeight, r)
	return weights[idx]
}

func TestWeightedRandomS6(t *testing.T) {
	cnt := make(map[int]int)
	list := []int{1, 2, 4, 8}
	for i := 0; i < 100; i++ {
		v := weightedRandomS6(list)
		cnt[v]++
	}

	for k, v := range cnt {
		fmt.Println(k, v)
	}
}

func weightedRandomS6(weights []int) int {
	rand.Seed(time.Now().UnixNano())
	sum := 0
	var sumWeight []int
	for _, v := range weights {
		sum += v
		sumWeight = append(sumWeight, sum)
	}
	r := rand.Intn(sum)
	idx := searchInts(sumWeight, r)
	return weights[idx]
}

func searchInts(a []int, x int) int {
	i, j := 0, len(a)
	for i < j {
		h := int(uint(i+j) >> 1)
		if a[h] < x {
			i = h + 1
		} else {
			j = h
		}
	}
	return i
}

const BMMinWeights = 10
const BMMaxWeights = 1000000

func BenchmarkWeightedRandomS5(b *testing.B) {
	for n := BMMinWeights; n <= BMMaxWeights; n *= 10 {
		b.Run(strconv.Itoa(n), func(b *testing.B) {
			weights := mockWeights(n)
			b.ResetTimer()
			b.RunParallel(func(pb *testing.PB) {
				for pb.Next() {
					weightedRandomS5(weights)
				}
			})
		})
	}
}

func BenchmarkWeightedRandomS6(b *testing.B) {
	for n := BMMinWeights; n <= BMMaxWeights; n *= 10 {
		b.Run(strconv.Itoa(n), func(b *testing.B) {
			weights := mockWeights(n)
			b.ResetTimer()
			b.RunParallel(func(pb *testing.PB) {
				for pb.Next() {
					weightedRandomS6(weights)
				}
			})
		})
	}
}

func mockWeights(n int) []int {
	weights := make([]int, 0, n)
	for i := 0; i < n; i++ {
		w := rand.Intn(10)
		weights = append(weights, w)
	}
	return weights
}

func TestWeightedRandomS7(t *testing.T) {
	cnt := make(map[int]int)
	list := []float64{1, 2, 4, 8}
	for i := 0; i < 100; i++ {
		v := weightedRandomS7(list)
		cnt[v]++
	}

	for k, v := range cnt {
		fmt.Println(k, v)
	}
}

func weightedRandomS7(weights []float64) int {
	var sum float64
	var winner int
	rand.Seed(time.Now().UnixNano())
	for i, v := range weights {
		sum += v
		f := rand.Float64()
		if f*sum < v {
			winner = i
		}
	}
	return winner
}
