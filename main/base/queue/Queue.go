package queue

import "fmt"

//使用 别名 的方式 将 []int 扩展为 队列
type Queue []int

func (q *Queue) Push(v int) {
	*q = append(*q, v)
}

func (q *Queue) Pop() (int, error) {
	if nil == q || 0 == len(*q) {
		fmt.Println("queue is empty")
		return -1, fmt.Errorf("queue is empty")
	}
	*q = (*q)[1:]
	return (*q)[0], nil
}

func (p *Queue) IsEmpty() bool {
	return nil == p || 0 == len(*p)
}

func main() {

}
