package parser

type NodeStack struct {
	Nodes []*Node
}

func NewNodeStack() *NodeStack {
	return &NodeStack{}
}

func (s *NodeStack) Push(node *Node) {
	s.Nodes = append(s.Nodes, node)
}

func (s *NodeStack) Pop() *Node {
	node := s.Nodes[len(s.Nodes)-1]
	s.Nodes = s.Nodes[:len(s.Nodes)-1]
	return node
}

func (s *NodeStack) Peek() *Node {
	return s.Nodes[len(s.Nodes)-1]
}

func (s *NodeStack) IsEmpty() bool {
	return len(s.Nodes) == 0
}

func (s *NodeStack) HasSequence(sequenceNodeType []int) bool {
	if len(s.Nodes) < len(sequenceNodeType) {
		return false
	}

	for i := 0; i < len(sequenceNodeType); i++ {
		if s.Nodes[len(s.Nodes)-1-i].Token.Type != sequenceNodeType[len(sequenceNodeType)-1-i] {
			return false
		}
	}

	return true
}

func (s *NodeStack) Size() int {
	return len(s.Nodes)
}
