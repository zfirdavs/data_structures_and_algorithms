package main

import (
	"flag"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const threshold = 1 << 13

func random(n int) []int {
	s := make([]int, n)

	src := rand.NewSource(time.Now().UnixNano())
	rand := rand.New(src)

	for i := 0; i < n; i++ {
		s[i] = rand.Intn(n)
	}

	return s
}

var randLength = flag.Int("randn", 20, "the random slice length")

func main() {
	flag.Parse()

	r := random(*randLength)
	fmt.Printf("before: %#v\n", r)

	parallelMergeSort(r)
	fmt.Printf("after: %#v\n", r)
}

func sequentiallyMergeSort(s []int) {
	if len(s) <= 1 {
		return
	}

	middle := len(s) / 2
	sequentiallyMergeSort(s[:middle])
	sequentiallyMergeSort(s[middle:])

	merge(s, middle)
}

func merge(s []int, middle int) {
	var (
		leftPart, curr int
		rightPart      = middle
		length         = len(s)
		helper         = make([]int, length)
		high           = length - 1
	)

	copy(helper, s)

	for leftPart <= middle-1 && rightPart <= high {
		if helper[leftPart] <= helper[rightPart] {
			s[curr] = helper[leftPart]
			leftPart++
		} else {
			s[curr] = helper[rightPart]
			rightPart++
		}
		curr++
	}

	for leftPart <= middle-1 {
		s[curr] = helper[leftPart]
		curr++
		leftPart++
	}
}

func parallelMergeSort(s []int) {
	len := len(s)

	if len <= 1 {
		return
	}

	if len <= threshold {
		sequentiallyMergeSort(s)
		return
	}

	middle := len / 2
	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		defer wg.Done()
		parallelMergeSort(s[:middle])
	}()

	parallelMergeSort(s[middle:])

	wg.Wait()
	merge(s, middle)
}
