package utils

import "fmt"

type Queue struct {
	slice []int;
}

func (q *Queue) PushBack(v int) {
	q.slice = append(q.slice, v);
}

func (q *Queue) PopFront() {
	q.slice = q.slice[1:];
}

func (q *Queue) Empty() bool {
	return len(q.slice) == 0;
}

func (q *Queue) Front() int {
	return q.slice[0];
}

func (q *Queue) PrintQueue() {
	fmt.Println(q.slice);
}