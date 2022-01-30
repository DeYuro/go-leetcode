package problem141
import (
	"github.com/deyuro/go-leetcode/structures"
	"reflect"
	"testing"
)

type test struct {
	input *structures.ListNode
	expected bool
}

func getTestcases() []test {
	return []test{
		{
			structures.CreateListNode([]int{3,2,0,-4}),
			true,
		},
		{
			structures.CreateListNode([]int{1,2}),
			true,
		},
		{
			structures.CreateListNode([]int{1}),
			false,
		},
	}
}

//TODO testcases because leetcode pass pos of linked list to make cycle
func TestHasCycle(t *testing.T)  {
	for _,v := range getTestcases() {
		output := hasCycle(v.input)
		if !reflect.DeepEqual(output, v.expected) {
			t.Errorf("Expect %v: got %v", v.expected, output)
		}
	}
}