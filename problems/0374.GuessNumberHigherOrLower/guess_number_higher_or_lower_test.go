package problem374

import (
	"reflect"
	"testing"
)

type test struct {
	input int
	pick int
	expected int
}

func getTestcases() []test {
	return []test{
		{
			10,
			6,
			6,
		},
		{
			1,
			1,
			1,
		},
		{
			2,
			1,
			1,
		},
	}
}

func TestGuessNumber(t *testing.T) {
	for _, v := range getTestcases() {
		output := guessNumber(v.input, v.pick)
		if !reflect.DeepEqual(output, v.expected) {
			t.Errorf("Expect %v: got %v", v.expected, output)
		}
	}
}

func BenchmarkGuessNumber(b *testing.B) {
	test := getTestcases()[0]
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		guessNumber(test.input, test.pick)
	}
}