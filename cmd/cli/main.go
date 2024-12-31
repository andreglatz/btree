package main

import (
	"fmt"

	"github.com/andreglatz/btree/internal/graph"
	"github.com/andreglatz/btree/internal/queue"
)

func main() {
	tree := graph.NewNode(1, graph.NewNode(2, graph.NewNode(4), graph.NewNode(5)), graph.NewNode(3, graph.NewNode(6)))
	q := queue.NewQueue[[]*graph.Node[int]]()

	q.Enqueue([]*graph.Node[int]{tree})

	for q.Size() > 0 {
		nodes := q.Dequeue()
		var level []*graph.Node[int]

		for _, v := range nodes {
			fmt.Print(v.Value, " ")
			level = append(level, v.Children...)
		}
		fmt.Println()

		if len(level) > 0 {
			q.Enqueue(level)
		}
	}
}
