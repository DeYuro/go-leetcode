package meeting_room_2

import (
	"container/heap"
	"sort"
	"strconv"
	"testing"
)

type Interval struct {
	Start int
	End   int
}

// min heap for intervals and sorted by End
type IntervalMinHeap []Interval

func (h IntervalMinHeap) Len() int {
	return len(h)
}

// note: the best practice is to use pointer receiver for all methods if some methods need it
// but the actual code is not as readable.. many (*h) in the code
func (h IntervalMinHeap) Less(i, j int) bool {
	// min heap sort by end time, so the top is always the next meeting that will end
	return h[i].End < h[j].End
}

func (h IntervalMinHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *IntervalMinHeap) Push(x interface{}) {
	*h = append(*h, x.(Interval))
}

// this is method that package heap will use for the internal data
// too add and remove things from heap, use heap.Push and heap.Pop
// since this function is for the heap package to use, this function actually remove and
// return the last element in the slice Len() - 1 instead of "pop head of heap"
func (h *IntervalMinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

func minMeetingRooms(intervals []Interval) int {
	if intervals == nil || len(intervals) == 0 {
		return 0
	}

	intervalsOrderedByStart := make([]Interval, len(intervals))
	copy(intervalsOrderedByStart, intervals)
	sort.Slice(intervalsOrderedByStart, func (i ,j int) bool {
		return intervalsOrderedByStart[i].Start < intervalsOrderedByStart[j].Start
	})

	h := &IntervalMinHeap{}
	heap.Init(h)

	for i := 0; i < len(intervalsOrderedByStart); i++ {
		// note: when the heap is empty, add to heap and you are done
		// remember to add continue, otherwise later logic will add it twice
		if h.Len() == 0 {
			heap.Push(h, intervalsOrderedByStart[i])
			continue
		}

		// the next meeting that about to stop
		m := (*h)[0]
		// note: we need <= instead of < here, because for back to back meetings
		// we don't need additional room
		if m.End <= intervalsOrderedByStart[i].Start {
			// the next meeting that is about to stop can stop now
			heap.Pop(h)
		}
		heap.Push(h, intervalsOrderedByStart[i])
	}

	// whatever number of meeting in the heap is the number of conference room we need
	// because at the time that meeting is about to start, there is no current meeting
	// room available for it, we will need additional meeting room, since we didn't pop
	// in this case, the size of heap increased too.
	return h.Len()
}


func numWaterBottles(numBottles int, numExchange int) int {
	sum, emptyBottles := 0, 0
	for numBottles > 0 || emptyBottles >= numExchange {
		sum += numBottles
		emptyBottles += numBottles
		numBottles = emptyBottles / numExchange
		emptyBottles = emptyBottles % numExchange
	}
	return sum
}

func numWaterBottles(numBottles int, numExchange int) int {
	var res int
	cur := 0
	for numBottles > 0 {
		res += numBottles
		tmp := (cur + numBottles) % numExchange
		numBottles = (cur + numBottles) / numExchange
		cur = tmp
	}
	return res
}


func leastBricks(wall [][]int) int {
	wlen := len(wall)

	// get the width of row
	width := 0
	for i := 0; i < len(wall[0]); i++ {
		width += wall[0][i]
	}

	// store count of borders per index in map
	edgeMap := make(map[int]int)

	min := wlen
	for i := 0; i < wlen; i++ {
		w := wall[i]

		x := 0
		for j := 0; j < len(w); j++ {
			x += w[j]
			if x == width {
				break
			}

			var bricks int
			if cnt, ok := edgeMap[x]; ok {
				edgeMap[x] = cnt + 1
				bricks = wlen - edgeMap[x]
			} else {
				edgeMap[x] = 1
				bricks = wlen - 1
			}

			if bricks < min {
				min = bricks
			}
		}
	}
	return min
}


import "container/list"

type state struct {
	node *TreeNode
	lvl  int
}

func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	queue := list.New()
	queue.PushBack(&state{root, 1})

	for queue.Len() > 0 {
		v := queue.Front()
		q := v.Value.(*state)
		p, lvl := q.node, q.lvl
		queue.Remove(v)
		if p.Left == nil && p.Right == nil {
			return lvl
		}
		if p.Left != nil {
			queue.PushBack(&state{p.Left, lvl + 1})
		}
		if p.Right != nil {
			queue.PushBack(&state{p.Right, lvl + 1})
		}
	}
	return 0
}

func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	q := []*TreeNode{root}

	depth := 1
	for len(q) > 0 {
		s := len(q)
		for i := 0; i < s; i++ {
			curr := q[0]
			q = q[1:]

			if curr.Left == nil && curr.Right == nil {
				return depth
			}

			if curr.Left != nil {
				q = append(q, curr.Left)
			}

			if curr.Right != nil {
				q = append(q, curr.Right)
			}
		}
		depth++
	}

	return -1
}


func openLock(deadends []string, target string) int {
	visited := make(map[string]bool)
	for _, dead := range deadends {
		visited[dead] = true
	}

	if visited["0000"] {
		return -1
	}
	visited["0000"] = true
	queue := []string{"0000"}
	rotate := 0
	for len(queue) > 0 {
		newQueue := []string{}
		for _, s := range queue {
			if s == target {
				return rotate
			}

			for i := 0; i < len(s); i++ {
				j := (int(s[i]-'0') + 1) % 10
				ns := s[:i] + strconv.Itoa(j) + s[i+1:]
				if !visited[ns] {
					visited[ns] = true
					newQueue = append(newQueue, ns)
				}
				j = (int(s[i]-'0') - 1 + 10) % 10
				ns = s[:i] + strconv.Itoa(j) + s[i+1:]
				if !visited[ns] {
					visited[ns] = true
					newQueue = append(newQueue, ns)
				}
			}
		}

		queue = newQueue
		rotate++
	}
	return -1
}