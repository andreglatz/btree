package main

import (
	"fmt"

	"github.com/andreglatz/tree/internal/graph"
	"github.com/andreglatz/tree/internal/queue"
)

func main() {
	tree := graph.NewNode(1, graph.NewNode(2, graph.NewNode(4), graph.NewNode(5)), graph.NewNode(3, graph.NewNode(6)))
	q := queue.NewQueue[*graph.Node[int]]()

	q.Enqueue(tree)

	for q.Size() > 0 {
		levelSize := q.Size()

		for i := 0; i < int(levelSize); i++ {
			node := q.Dequeue()
			fmt.Print(node.Value, " ")

			for _, child := range node.Children {
				q.Enqueue(child)
			}
		}

		fmt.Println()
	}
}
