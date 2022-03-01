package problem692

import "container/heap"

type Words struct {
	Word string
	Count int
}

type maxHeap []Words

func(h maxHeap) Len() int  {return len(h)}
func(h maxHeap) Less(i,j int) bool {

	if h[i].Count == h[j].Count {
		return h[i].Word < h[j].Word
	}

	return h[i].Count > h[j].Count
}


func(h maxHeap) Swap(i,j int) {
	h[i],h[j] = h[j],h[i]

}

func(h *maxHeap) Push(x interface{}) { *h = append(*h, x.(Words))}

func(h *maxHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}



func topKFrequent(words []string, k int) []string {
	res := []string{}

	m := make(map[string]int)

	for _,v := range words {
		m[v]++
	}

	var h maxHeap

	for i, v := range m {
		heap.Push(&h,Words{
			Word: i,
			Count: v,
		})
	}

	for k > 0 {
		res = append(res, heap.Pop(&h).(Words).Word)
		k--
	}

	return res
}