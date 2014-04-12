package binary

type Node struct {
	parent      *Node
	left, right *Node
	Key         int
}

type Tree struct {
	root *Node
}

func (T *Tree) Insert(z *Node) {
	var y *Node
	x := T.root

	// Search for
	for x != nil {
		y = x // keep track of the parent
		if z.Key < x.Key {
			x = x.left
		} else {
			x = x.right
		}
	}
	z.parent = y
	if y == nil {
		T.root = z
	} else if z.Key < y.Key {
		y.left = z
	} else {
		y.right = z
	}

}

func (T *Tree) transplant(u, v *Node) {
	if u.parent == nil {
		T.root = v
	} else if u == u.parent.left {
		u.parent.left = v
	} else {
		u.parent.right = v
	}
	if v != nil {
		v.parent = u.parent
	}
}

func (T *Tree) Min(z *Node) *Node {
	p := z
	var y *Node

	for p != nil {
		y = p
		p = p.left
	}

	return y
}

func (T *Tree) Delete(z *Node) {
	if z == nil {
		return
	}

	if z.left == nil {
		T.transplant(z, z.right)
	} else if z.right == nil {
		T.transplant(z, z.left)
	} else {
		y := T.Min(z.right)
		if y.parent != z {
			T.transplant(y, y.right)
			y.right = z.right
			y.right.parent = y
		}
		T.transplant(z, y)
		y.left = z.left
		y.left.parent = y
	}
}

func (T *Tree) Search(value int) *Node {
	p := T.root
	for p != nil {
		if p.Key == value {
			return p
		} else if p.Key < value {
			p = p.right
		} else {
			p = p.left
		}
	}
	return p
}
