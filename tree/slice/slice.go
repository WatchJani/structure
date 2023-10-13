package tree

type Node struct {
	value int
	node  *Node
}

type Tree struct {
	LastNode *Node
	Memory   []*Node
}

func NewStruct(capacity int) Tree {
	return Tree{
		Memory: make([]*Node, 0, capacity),
	}
}

func New(value int) *Node {
	return &Node{
		value: value,
	}
}

func (t *Tree) Inset(value int) {
	newNode := New(value)

	if t.LastNode != nil {
		t.LastNode.node = newNode
	}

	t.LastNode = newNode
	t.Memory = append(t.Memory, newNode)
}

func (t *Tree) Read() {

}
