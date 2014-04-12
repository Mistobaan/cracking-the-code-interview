package collections

type LinkedList struct {
	head *Node
}

type Node struct {
	Key  interface{}
	next *Node
}

// Single Linked List
func NewLinkedList() *LinkedList {
	return &LinkedList{}
}

func NewNode(value interface{}) *Node {
	return &Node{
		Key: value,
	}
}

func (l *LinkedList) Insert(n *Node) {
	if n == nil {
		return
	}

	if l.head != nil {
		n.next = l.head.next
	}

	l.head = n
}

func (l *LinkedList) Delete(n *Node) {
	if n == nil {
		return
	}

	p := l.head

	for p != nil && p.next != n {
		p = p.next
	}

	if p != nil {
		p.next = n.next
	}
}

func (l *LinkedList) Search(value interface{}) *Node {
	p := l.head

	for p != nil && p.Key != value {
		p = p.next
	}

	return p
}
