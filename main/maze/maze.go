package main

import (
	"fmt"
	"os"
)

func readMaze(filename string) [][]int {
	file, e := os.Open(filename)
	if nil != e {
		panic(e)
	}
	var rowSize, colSize int
	fmt.Fscanf(file, "%d %d", &rowSize, &colSize)

	// 初始化二维数组(切片)
	maze := make([][]int, rowSize)
	for i := range maze {
		maze[i] = make([]int, colSize)
		for j := range maze[i] {
			fmt.Fscanf(file, "%d", &maze[i][j])
		}
	}

	return maze
}

type point struct {
	i, j int
}

var rundirs = [4]point{
	{-1, 0}, //左
	{0, -1}, //下
	{1, 0},  //右
	{0, 1},  //上
}

func (curr point) step(dir point) point {
	return point{curr.i + dir.i, curr.j + dir.j}
}

func (p point) at(grid [][]int) (int, bool) {
	if p.i < 0 || p.i >= len(grid) {
		return 0, false
	}

	if p.j < 0 || p.j >= len(grid[p.i]) {
		return 0, false
	}

	return grid[p.i][p.j], true
}

func walkMaze(maze [][]int, start point, end point) {
	// 初始化 steps 用于存放走过的路径
	steps := make([][]int, len(maze))
	for i := range steps {
		steps[i] = make([]int, len(maze[i]))
	}

	// 待探索的 队列
	q := &queue{}

	// 放入起始点进入队列
	q.push(start)

	// 队列为空 说明 已经不存在
	for len(*q) > 0 {
		cur := q.pop()

		// 如果 当前点 == 终点
		if cur == end {
			break
		}

		for _, dir := range rundirs {
			next := cur.step(dir)
			q.push(next)

			value, ok := next.at(maze)
		}

	}
}

type queue []point

func (q *queue) push(p point) {
	if nil == q {
		panic("queue can not be null")
	}

	*q = append(*q, p)
}

func (q *queue) pop() point {
	if nil == q {
		panic("queue can not be null")
	}

	p := (*q)[0]
	*q = (*q)[1:]
	return p
}

func main() {
	maze := readMaze("./main/maze/maze.in")
	fmt.Println(maze)
	walkMaze(maze, point{0, 0}, point{len(maze), len(maze[0])})
	fmt.Println(rundirs)
	fmt.Println(left, down, right, up)
}
