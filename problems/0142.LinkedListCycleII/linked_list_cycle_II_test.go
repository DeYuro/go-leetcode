package problem142
import (
	"github.com/deyuro/go-leetcode/structures"
	"reflect"
	"testing"
)

type test struct {
	input *structures.ListNode
	expected *structures.ListNode
}

func getTestcases() []test {
	return []test{
		{
			structures.CreateListNode([]int{3,2,0,-4}),
			structures.CreateListNode([]int{2}),
		},
		{
			structures.CreateListNode([]int{1,2}),
			structures.CreateListNode([]int{2}),
		},
		{
			structures.CreateListNode([]int{1}),
			nil,
		},
	}
}

//TODO testcases because leetcode pass pos of linked list to make cycle
func TestDetectCycle(t *testing.T)  {
	for _,v := range getTestcases() {
		output := detectCycle(v.input)
		if !reflect.DeepEqual(output, v.expected) {
			t.Errorf("Expect %v: got %v", v.expected, output)
		}
	}
}