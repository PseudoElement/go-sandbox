package datastructures

func TestTree() {
	tree := NewHtmlTree()

	node2 := NewTreeNode(2, "html-2")
	node3 := NewTreeNode(3, "html-3")

	tree.Root.AppendChild(node2)
	tree.Root.AppendChild(node3)

	node4 := NewTreeNode(4, "html-4")
	node5 := NewTreeNode(5, "html-5")
	node6 := NewTreeNode(6, "html-6")

	node2.AppendChild(node4)
	node2.AppendChild(node5)
	node2.AppendChild(node6)

	node7 := NewTreeNode(7, "html-7")
	node8 := NewTreeNode(8, "html-8")
	node9 := NewTreeNode(9, "html-9")

	node4.AppendChild(node7)

	node5.AppendChild(node8)
	node5.AppendChild(node9)

	node10 := NewTreeNode(10, "html-10")
	node11 := NewTreeNode(11, "html-11")

	node7.AppendChild(node10)
	node7.AppendChild(node11)

	// node, stepCount := tree.FindByIdDFSRecursive(tree.Root, "html-8", new(int))

	// fmt.Printf("node %+v, stepCount %v\n", node, stepCount)
}

// TASK EXAMPLE
// find all combinations on values in tree, which sum equals to n

type TreeNode struct {
	Id       string
	Value    int
	Children []*TreeNode
}

func NewTreeNode(value int, id string) *TreeNode {
	return &TreeNode{Id: id, Value: value, Children: make([]*TreeNode, 0)}
}

func (node *TreeNode) AppendChild(child *TreeNode) {
	node.Children = append(node.Children, child)
}

type Tree struct {
	Root *TreeNode
}

func NewHtmlTree() *Tree {
	return &Tree{Root: NewTreeNode(1, "html")}
}

/*
 * BFS
 * Breadth-first search
 */
func (t *Tree) FindByIdBFS(id string) *TreeNode {
	queue := make([]*TreeNode, 0)
	queue = append(queue, t.Root)

	for len(queue) > 0 {
		nextUp := queue[0]
		queue = queue[1:]
		if nextUp.Id == id {
			return nextUp
		}
		if len(nextUp.Children) > 0 {
			queue = append(queue, nextUp.Children...)
		}
	}

	return nil
}

/*
 * DFS
 * Depth-first search
 * + counts how many steps was performed to find node
 */
func (t *Tree) FindByIdDFSRecursive(node *TreeNode, id string, stepCount *int) (*TreeNode, int) {
	*stepCount++

	if node.Id == id {
		return node, *stepCount
	}

	if len(node.Children) > 0 {
		for _, child := range node.Children {
			found, _ := t.FindByIdDFSRecursive(child, id, stepCount)
			if found != nil {
				return found, *stepCount
			}
		}
	}

	return nil, *stepCount
}

func (t *Tree) FindNodeByIdStack(node *TreeNode, id string, stepCount *int) (*TreeNode, int) {
	if node == nil {
		return nil, 0
	}

	// Use a stack for iterative DFS
	stack := []*TreeNode{node}

	for len(stack) > 0 {
		*stepCount++

		// Pop the last node from stack
		current := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		// Check if current node matches
		if current.Id == id {
			return current, *stepCount
		}

		// Push children to stack in reverse order to maintain DFS order
		for i := len(current.Children) - 1; i >= 0; i-- {
			stack = append(stack, current.Children[i])
		}
	}

	return nil, *stepCount
}
