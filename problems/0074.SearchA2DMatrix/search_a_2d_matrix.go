package problem74

func searchMatrix(matrix [][]int, target int) bool {
	l, h := 0, len(matrix)-1

	for l <= h {
		m := (l + h) / 2
		if target < matrix[m][0] {
			h = m - 1
		}

		if target > matrix[m][len(matrix[m])-1] {
			l = m + 1
		}

		if target >= matrix[m][0] && target <= matrix[m][len(matrix[m])-1] {
			return bs(matrix[m], target)
		}
	}

	return false
}

func bs(s []int, t int) bool {
	l, h := 0, len(s)-1

	for l <= h {
		m := (l + h) / 2

		if t > s[m] {
			l = m + 1
		}
		if t < s[m]{
			h = m - 1
		}

		if t == s[m] {
			return true
		}
	}

	return false
}


func searchMatrix2(matrix [][]int, target int) bool {
	m, n := len(matrix), len(matrix[0])
	low, high := 0, (m * n) - 1
	for ; low <= high; {
		mid := (low + high) / 2
		r, c := mid / n, mid % n
		println(matrix[r][c])
		if matrix[r][c] == target {
			return true
		} else if matrix[r][c] > target {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return false
}
//{ 1, 3, 5, 7},
//{10,11(middle),16,20},
//{23,30,34,60}