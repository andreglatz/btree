package graph

type Node[T any] struct {
	Value    T
	Children []*Node[T]
}

func NewNode[T any](value T, children ...*Node[T]) *Node[T] {
	return &Node[T]{
		Value:    value,
		Children: children,
	}
}
