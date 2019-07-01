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

const (
	left  = iota + 1 // 1
	down             // 2
	right            // 3
	up               // 4
)

func (curr point) step() point {
	return point{0, 0}
}

func walkMaze(maze [][]int, start point, end point) {

}

func main() {
	maze := readMaze("./main/maze/maze.in")
	fmt.Println(maze)
	walkMaze(maze, point{0, 0}, point{len(maze), len(maze[0])})
	fmt.Println(rundirs)
	fmt.Println(left, down, right, up)
}
