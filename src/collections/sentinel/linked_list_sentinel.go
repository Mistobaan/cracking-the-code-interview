package sentinel

type Node struct {
	Key  interface{}
	next *Node
	prev *Node
}

// Circular doubly linked list with a sentinel
type LinkedList struct {
	null *Node
}

func NewLinkedList() *LinkedList {
	ll := &LinkedList{
		null: &Node{},
	}
	ll.null.next = ll.null
	ll.null.prev = ll.null
	return ll
}

func (l *LinkedList) Insert(n *Node) {
	n.next = l.null.next
	l.null.next.prev = n
	l.null.next = n
	n.prev = l.null
}

func (l *LinkedList) Delete(n *Node) {
	n.prev.next, n.next.prev = n.next, n.prev
}

func (l *LinkedList) Search(value interface{}) *Node {
	n := l.null.next
	for n != l.null && n.Key != value {
		n = n.next
	}
	return n
}
