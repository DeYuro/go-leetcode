package problem374

//We are playing the Guess Game. The game is as follows:
//
//I pick a number from 1 to n. You have to guess which number I picked.
//
//Every time you guess wrong, I will tell you whether the number I picked is higher or lower than your guess.
//
//You call a pre-defined API int guess(int num), which returns three possible results:
//
//-1: Your guess is higher than the number I picked (i.e. num > pick).
//1: Your guess is lower than the number I picked (i.e. num < pick).
//0: your guess is equal to the number I picked (i.e. num == pick).
//Return the number that I picked.
//
//
//
//Example 1:
//
//Input: n = 10, pick = 6
//Output: 6
//Example 2:
//
//Input: n = 1, pick = 1
//Output: 1
//Example 3:
//
//Input: n = 2, pick = 1
//Output: 1

func guessNumber(num, pick int) int {
	low, high := 0, num

	for low <= high {
		mid := (low + high) / 2
		guess := guess(mid, pick)

		if guess > 0 {
			low = mid + 1
		}

		if guess < 0 {
			high = high - 1
		}

		if guess == 0 {
			return mid
		}
	}

	return -1
}

func guess(num, pick int) int {
	if num < pick {
		return 1
	}

	if num > pick {
		return -1
	}

	return 0
}