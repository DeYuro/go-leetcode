package atm

import (
	"reflect"
	"testing"
)

type test struct {
	input int
	bill map[int]int
	expected []int
}

func getTestcases() []test {
	return []test{
		{
			5000,
			map[int]int{
				50: 1000,
				100: 1000,
				200: 1000,
				500: 1000,
				1000: 1000,
				2000: 1000,
				5000: 1000,
			},
			[]int{5000},
		},
		{
			1500,
			map[int]int{
				50: 1000,
				100: 1000,
				200: 1000,
				500: 1000,
				1000: 1000,
				2000: 1000,
				5000: 1000,
			},
			[]int{1000, 500},
		},
		{
			1500,
			map[int]int{
				50: 1000,
				100: 1,
				200: 4,
				500: 1,
				1000: 0,
				2000: 1000,
				5000: 1000,
			},
			[]int{500,200,200,200,200,100,50,50},
		},
		//{
		//	210,
		//	map[int]int{
		//		50: 1000,
		//		100: 1000,
		//		200: 1000,
		//		500: 1000,
		//		1000: 1000,
		//		2000: 1000,
		//		5000: 1000,
		//	},
		//	[]int{},
		//},
		{
			210,
			map[int]int{
				30: 1000,
				50: 1000,
				100: 1000,
				200: 1000,
				500: 1000,
				1000: 1000,
				2000: 1000,
				5000: 1000,
			},
			[]int{100,50,30,30},
		},
		{
			120,
			map[int]int{
				30: 1000,
				50: 1000,
				100: 1000,
				200: 1000,
				500: 1000,
				1000: 1000,
				2000: 1000,
				5000: 1000,
			},
			[]int{30,30,30,30},
		},
	}
}

func TestAtm(t *testing.T) {
	for _,v := range getTestcases() {
		output := atm(v.input, v.bill)
		if !reflect.DeepEqual(output, v.expected) {
			t.Errorf("Expect %v: got %v", v.expected, output)
		}
	}
}