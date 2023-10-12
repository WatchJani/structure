package word

const ENGLISH_ALPHABET = 26

type Node struct {
	node  map[rune]*Node
	isEnd bool
}

type Tree struct {
	Memory []*Node
}

func New() *Node {
	return &Node{
		node: make(map[rune]*Node, ENGLISH_ALPHABET/2),
	}
}

func (n *Node) InsertNode(char rune, pointer *Node) {
	n.node[char] = pointer
}

func WordCollectorInit(capacity int) *Tree {
	Tree := Tree{
		Memory: make([]*Node, 0, capacity),
	}

	Tree.Memory = append(Tree.Memory, New())

	return &Tree
}

func (t *Tree) Insert(word string) {
	currentNode := t.Memory[0] //root

	for _, char := range word {
		if _, ok := currentNode.node[char]; !ok {
			t.Memory = append(t.Memory, New())                      // create new empty node
			currentNode.InsertNode(char, t.Memory[len(t.Memory)-1]) //on current node add new node with link
		}
		currentNode = currentNode.node[char] //switch to next linked node
	}

	currentNode.isEnd = true
}
