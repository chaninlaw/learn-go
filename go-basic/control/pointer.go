package control

type Employee struct {
	Name   string
	Salary int
}

type ListNode struct {
	Value int
	Next  *ListNode
}

func Prepend(head **ListNode, value int) {
	newNode := ListNode{Value: value, Next: *head}
	*head = &newNode
}

func ChangeValue(ptr *int) {
	*ptr = 50
}

func GiveRaise(e *Employee, raise int) {
	e.Salary += raise

}
