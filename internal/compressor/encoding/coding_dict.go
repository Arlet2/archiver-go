package encoding

type CodingDict map[int32]string

const (
	leftCode  = "0"
	rightCode = "1"
)

func CreateCodingDict(tree CodingTree) (dict CodingDict) {
	dict = make(CodingDict, 0)

	if IsTreeEmpty(tree) {
		return
	}

	generateNewCode(Node(tree), dict, "")

	return dict
}

func generateNewCode(node Node, dict CodingDict, code string) {
	if node.LeftChild == nil && node.RightChild == nil {
		dict[node.Value] = code
		return
	}

	if node.LeftChild != nil {
		generateNewCode(*node.LeftChild, dict, code+leftCode)
	}

	if node.RightChild != nil {
		generateNewCode(*node.RightChild, dict, code+rightCode)
	}
}
