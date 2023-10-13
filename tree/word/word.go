package word

const (
	ENGLISH_ALPHABET = 26
	ALLOCATED_NODE   = 2
)

type Node struct {
	node  map[rune]*Node
	isEnd bool
}

type Tree struct {
	Memory []*Node
}

func New() *Node {
	return &Node{
		node: make(map[rune]*Node, ALLOCATED_NODE), //every node can have unlimited nodes, but for memory we allocated 13 nodes for every node
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
		currentNode = currentNode.node[char] //switch to next linked node (both if it exists and if it doesn't exist)
	}

	currentNode.isEnd = true
}

func (t *Tree) IsExist(word string) *Node {
	currentNode := t.Memory[0] //root

	for _, char := range word {
		if _, ok := currentNode.node[char]; !ok {
			return nil
		}
		currentNode = currentNode.node[char] //switch to next linked node (both if it exists and if it doesn't exist)
	}

	return currentNode
}

// check if exist complete word in structure
func (t *Tree) IsExistWord(word string) bool {
	return t.IsExist(word).isEnd
}

// check if exist prefix of some word
func (t *Tree) IsExistPrefix(word string) bool {
	return t.IsExist(word) != nil
}

func (t *Tree) ReadAllNodes(word string) []string {
	root := t.IsExist(word)

	if root == nil {
		return nil
	}

	var result []string = make([]string, 0, 10)

	queue := make([]struct {
		currentNode *Node
		prefix      string
	}, 0, 10)

	queue = append(queue, struct {
		currentNode *Node
		prefix      string
	}{root, word})

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		currentNode, prefix := current.currentNode, current.prefix

		if currentNode.isEnd {
			result = append(result, prefix)
		}

		for char, child := range currentNode.node {
			childPrefix := prefix + string(char)
			queue = append(queue, struct {
				currentNode *Node
				prefix      string
			}{child, childPrefix})
		}
	}

	return result
}

// func (t *Tree) ReadAllNodes(word string) []string {
// 	root := t.IsExist(word)

// 	var mySuggestions []string = make([]string, 0, 10)

// 	stack := make([]struct {
// 		currentNode *Node
// 		prefix      *strings.Builder
// 	}, 0, 10)

// 	stack = append(stack, struct {
// 		currentNode *Node
// 		prefix      *strings.Builder
// 	}{
// 		currentNode: root,
// 		prefix:      &strings.Builder{},
// 	})

// 	stack[0].prefix.WriteString(word) // Dodajte početni prefiks u builder

// 	for len(stack) > 0 {
// 		current := stack[len(stack)-1]
// 		stack = stack[:len(stack)-1]

// 		currentNode, prefix := current.currentNode, current.prefix

// 		if currentNode.isEnd {
// 			mySuggestions = append(mySuggestions, prefix.String())
// 		}

// 		for char, child := range currentNode.node {
// 			newPrefix := prefix // Koristite isti builder za svaku granu
// 			newPrefix.WriteRune(char)
// 			stack = append(stack, struct {
// 				currentNode *Node
// 				prefix      *strings.Builder
// 			}{
// 				currentNode: child,
// 				prefix:      newPrefix,
// 			})
// 		}
// 	}

// 	return mySuggestions
// }
