package problem480

import (
	"container/heap"
)

type minHeap []int
type maxHeap []int


func(h minHeap) Len() int {return len(h)}
func(h minHeap) Less(a,b int) bool {return h[a] < h[b]}
func(h minHeap) Swap(a,b int) {h[a],h[b] = h[b], h[a]}
func(h *minHeap) Push(x interface{}) { *h = append(*h, x.(int))}
func(h *minHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}


func(h maxHeap) Len() int {return len(h)}

func(h maxHeap) Swap(a,b int) {h[a],h[b] = h[b], h[a]}
func(h maxHeap) Less(a,b int) bool {return h[a]> h[b]}

func(h *maxHeap) Push(x interface{}) { *h = append(*h, x.(int))}
func (h minHeap) Peek() interface{} {return h[0]}

func (h maxHeap) Peek() interface{} {return h[0]}
func(h *maxHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

func medianSlidingWindow(nums []int, k int) []float64 {
	var minH minHeap
	var maxH maxHeap

	res := []float64{}

	for i := 0; i < len(nums); i++ {
		if i >= k {
			remove(nums[i-k], &minH, &maxH)
		}

		heap.Push(&minH,nums[i])
		heap.Push(&maxH,heap.Pop(&minH).(int))

		balance(&minH, &maxH)
		if i >= k-1 {
			res = append(res,getMedian(&maxH, &minH, k))
		}
	}

	return res
}

func remove(removed int, min *minHeap, max *maxHeap) {
	deleted := false

	for i := 0; i < len(*min); i++ {
		if removed == (*min)[i] {
			heap.Remove(min, i)
			deleted = true
		}
	}

	if deleted {
		return
	}

	for i := 0; i < len(*max); i++ {
		if removed == (*max)[i] {
			heap.Remove(max, i)
		}
	}
}

func balance(min *minHeap, max *maxHeap) {
	if len(*max) > len(*min) {
		heap.Push(min, heap.Pop(max).(int))
	}

}

func getMedian(max *maxHeap, min *minHeap, k int) float64 {
	if k % 2 == 0 {
		return (float64(max.Peek().(int)) + float64(min.Peek().(int)))/2.0
	}

	return float64(min.Peek().(int))
}