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

// 两个 返回值	1.返回的是 给定 点坐标的 作为横纵坐标 数组对应的 值
//			   	2.返回的是 这个点 是否 是一个 合法的点。
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

			// 准备尝试去探索下一个坐标
			// 那么 下一个 坐标 是否能被探索， 有两个重要的 条件：
			// 1.也就是说 下一个 点 对应的 强的 位置 必须能走。不能为1 （因为 1-不能走，0-能走）
			// 且也不能为start
			//		maze as next is 0 and next != start
			// 2.也就是说 下一个 点 的 下一步 ，  只能为 初始化的0 的。否则 就代表 有人走过
			// 		steps at next is 0
			// 只有条件全满足 才能 继续 走下一步

			value, ok := next.at(maze)
			// 迷宫的下一个 位置 不能越界。且 值 也不能为1 （为1 就撞墙了 ， 不能被通过）
			if (!ok) || value == 1 {
				continue
			}

			// 不能是 起始点
			if next.i == start.i && next.j == start.j {
				continue
			}

			// 存储 走过 路径 （走的每一步 记录 的 数组）， 与迷宫 对应相同位置的地方 也不能 大于 0。大于0 代表 有人走过
			value1, ok1 := next.at(steps)
			if (!ok1) || value1 > 0 {
				continue
			}

			q.push(next)

			// 打标记
			steps[next.i][next.j] = steps[cur.i][cur.j] + 1
		}
	}

	for _, val := range steps {
		fmt.Println(val)
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
	//fmt.Println(maze)
	walkMaze(maze, point{0, 0}, point{len(maze), len(maze[0])})
	//fmt.Println(rundirs)
}
