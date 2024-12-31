# Tree Level-Order Traversal

This project demonstrates a simple implementation of **level-order traversal** (or breadth-first traversal) of a tree. The program processes each level of the tree and outputs the values grouped by their level.

---

## How It Works

The algorithm uses a **queue-based approach** to traverse the tree level by level:

1. Start with the root node.
2. Enqueue the nodes of the current level as a list.
3. Dequeue the list of nodes, print their values, and gather their children.
4. Enqueue the gathered children (if any) for the next level.
5. Repeat until all levels are processed.

The result is an output where each line represents the values at a specific level of the tree.

---

## Example Output

Given a tree structure:

```
        1
       / \
      2   3
     / \   \
    4   5   6
```

The program produces the following output:

```
1
2 3
4 5 6
```

Each line corresponds to the values at a specific level of the tree.
