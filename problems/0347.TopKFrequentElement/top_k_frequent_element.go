package problem347

import "container/heap"

type HeapNode []Node

type  Node struct{
    Key int
    Value int
}

func topKFrequent(nums []int, k int) []int {
    
    res := []int{}
    m := make(map[int]int)
    
    for _,v := range nums {
        m[v]++
    }
    
    nodes := []Node{}
    
    for k,v := range m {
        nodes = append(nodes, Node{Key:k, Value:v})
    } 
    
    var h HeapNode
    for _,v := range nodes {
        heap.Push(&h, v)
    }
    
    for len(h) > 0 && k > 0 {
        n := heap.Pop(&h).(Node)
        k--
        res = append(res, n.Key)
    }
    
    
    return res
}

func (h HeapNode) Len() int{
    return len(h)
}

func (h HeapNode) Swap(i,j int) {
    h[i],h[j] = h[j],h[i]
}
// max heap
func (h HeapNode) Less(i,j int) bool{
    return h[i].Value > h[j].Value
}

func (h *HeapNode) Push(x interface{}) {
    *h = append(*h, x.(Node))
}

func (h *HeapNode) Pop() interface{} {
    old := *h
    n := len(old)
    x := old[n-1]
    *h = old[:n-1]
    
    return x
}

