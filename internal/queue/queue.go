package queue

type Queue[T any] interface {
	Enqueue(T)
	Dequeue() T
	Size() uint
}

type queue[T any] struct {
	values []T
	size   uint
}

func NewQueue[T any]() Queue[T] {
	return &queue[T]{
		values: make([]T, 0),
	}
}

func (q *queue[T]) Enqueue(value T) {
	q.size++
	q.values = append(q.values, value)
}

func (q *queue[T]) Dequeue() T {
	first := q.values[0]

	q.values = q.values[1:]
	q.size--

	return first
}

func (q *queue[T]) Size() uint {
	return q.size
}
