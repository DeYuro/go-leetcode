package main

//dp[i] = deque[0] + nums[i]

func maxResult(nums []int, k int) int {
	n := len(nums)
	dp := make([]int, n) //dp[i] : the max score at index i
	dp[0] = nums[0]

	//deque is a desc monotonous queue
	//the deque maintains DP values in desc order of the sliding window with max size K
	deque := make([]int, 0)
	deque = append(deque, 0)

	for i := 1; i < n; i++ {
		//remove max index out of the deque
		for len(deque) > 0 && deque[0] < i - k {
			deque = deque[1:]
		}

		//dp[i] = deque[0] + nums[i]
		dp[i] = dp[deque[0]] + nums[i]

		//maintain the deque
		for len(deque) > 0 && dp[deque[len(deque) - 1]] < dp[i] {
			deque = deque[:len(deque) - 1]
		}
		deque = append(deque, i) // be noted append the index into the queue
	}
	return dp[n - 1]
}

func canJump(nums []int) bool {
	lastPos := len(nums)-1 // last position can reach the end index
	for i := len(nums)-1; i >= 0; i-- {
		if i + nums[i] >= lastPos {
			lastPos = i
		}
	}
	return lastPos == 0
}

//----------------------
func jump(nums []int) int {
	max, res, curMax := 0, 0, 0
	for i := 0; i < len(nums) - 1; i++ {
		max = Max(max, nums[i] + i)
		if curMax == i {
			res++
			curMax = max
		}
	}
	return res
}

func Max(a, b int) int {
	if a > b { return a }
	return b
}

//---------------------------
func canReach(arr []int, start int) bool {
	queue, l := []int{ start }, len(arr)
	seen := make([]bool, l)
	seen[start] = true
	for len(queue) > 0 {
		index := queue[0]
		queue = queue[1:]
		if arr[index] == 0 { return true }
		for _, nextIndex := range []int{ index + arr[index], index - arr[index] } {
			if 0 <= nextIndex && nextIndex < l && seen[nextIndex] == false {
				seen[nextIndex] = true
				queue = append(queue, nextIndex)
			}
		}
	}
	return false
}

///

func maxJumps(arr []int, d int) int {
	dp := make([]int, len(arr))
	jumps := 0
	for i := range arr {
		if curJump := dynamicProgramming(dp, arr, i, d); curJump > jumps {
			jumps = curJump
		}
	}
	return jumps
}

func dynamicProgramming(dp, arr []int, index, k int) int {
	if dp[index] != 0 {
		return dp[index]
	}
	jump := 1
	for i := 1; i <= k; i++ {
		if index-i < 0 {
			break
		}
		if arr[index] <= arr[index-i] {
			break
		}
		if currJump := dynamicProgramming(dp, arr, index-i, k) + 1; currJump > jump {
			jump = currJump
		}
	}
	for i := 1; i <= k; i++ {
		if index+i >= len(arr) {
			break
		}
		if arr[index] <= arr[index+i] {
			break
		}
		if currJump := dynamicProgramming(dp, arr, index+i, k) + 1; currJump > jump {
			jump = currJump
		}
	}
	dp[index] = jump
	return dp[index]
}