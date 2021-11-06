package atm

import (
	"reflect"
	"testing"
)

type test struct {
	input int
	bill map[int]int
	expected map[int]int
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
			map[int]int{5000:1},
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
			map[int]int{1000:1, 500:1},
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
			map[int]int{500:1, 200: 4, 100:1, 50:2},
		},
		{
			210,
			map[int]int{
				50: 1000,
				100: 1000,
				200: 1000,
				500: 1000,
				1000: 1000,
				2000: 1000,
				5000: 1000,
			},
			nil,
		},
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
			map[int]int{100:1, 50:1, 30:2},
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
			map[int]int{30:4},
		},
	}
}

func TestAtm(t *testing.T) {
	for _,v := range getTestcases() {
		output := atm2(v.input, v.bill)
		if !reflect.DeepEqual(output, v.expected) {
			t.Errorf("Expect %v: got %v", v.expected, output)
		}
	}
}