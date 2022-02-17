package problem239

type deque []int

func (d *deque) PushLeft(x interface{}) {
	*d = append([]int{x.(int)}, *d...)
}
func (d *deque) PushRight(x interface{}) {
	*d = append(*d,x.(int))
}

func (d *deque) PopLeft()interface{} {
	old := *d
	n := len(old)
	x := old[0]
	*d = old[1:n]
	return x
}

func (d *deque) PopRight()interface{} {
	old := *d
	n := len(old)
	x := old[n-1]
	*d = old[:n-1]
	return x
}

func (d *deque) seeLeft()interface{} {
	old := *d
	x := old[0]
	return x
}

func (d *deque) seeRight()interface{} {
	old := *d
	n := len(old)
	x := old[n-1]
	return x
}

func (d deque) isEmpty() bool {
	return len(d) == 0
}

func maxSlidingWindow(nums []int, k int) []int {
	res := []int{}
	l,r := 0,0

	var d deque

	for r < len(nums){
		for !d.isEmpty() && nums[d.seeRight().(int)] < nums[r] {
			d.PopRight()
		}
		d.PushRight(r)

		if l > d.seeLeft().(int) {
			d.PopLeft()
		}

		if r + 1 >= k {
			res = append(res, nums[d.seeLeft().(int)])
			l++
		}

		r++
	}

	return res
}

func maxSlidingWindow2(nums []int, k int) []int {
	var deq deque
	var res []int
	for i := 0; i < len(nums); i++ {

		for !deq.IsEmpty() && (i-k == deq.PeekBack().(int)) {
			deq.PopBack()
		}

		for !deq.IsEmpty() && nums[deq.PeekFront().(int)] < nums[i] {
			deq.PopFront()
		}

		deq.PushFront(i)

		if i >= k - 1 {
			res = append(res, nums[deq.PeekBack().(int)])
		}
	}

	return res
}