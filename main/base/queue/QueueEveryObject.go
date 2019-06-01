package queue

import "fmt"

type QueueEveryObject []interface{}

func (q *QueueEveryObject) Push(v interface{}) {
	*q = append(*q, v)
}

func (q *QueueEveryObject) Pop() (interface{}, error) {
	if nil == q || 0 == len(*q) {
		fmt.Println("queue is empty")
		return -1, fmt.Errorf("queue is empty")
	}
	*q = (*q)[1:]
	return (*q)[0], nil
}

func (p *QueueEveryObject) IsEmpty() bool {
	return nil == p || 0 == len(*p)
}
