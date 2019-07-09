package main

import "strings"

// ---------------------------- 接口的组合 -------------------------------
/**

 */
// ---------------------------- 接口的组合 -------------------------------

// Read 接口 负责读
type Read interface {
	ReadFile() []byte
}

type ReadImpl struct {
	FilePath string
}

func (r ReadImpl) ReadFile() []byte {
	if 0 == len(r.FilePath) {
		panic("file path i")
	}

	//判断字符串相等 也可以 使用 == 代替
	if !strings.EqualFold("123.txt", r.FilePath) {
		panic("the file is not found")
	}

	return []byte("12345")
}

// Wtite 接口 负责写
type Write interface {
	WriteFile(contents []byte) bool
}

type WriteImpl struct {
	FilePath string
}

func (w WriteImpl) WriteFile(contents []byte) bool {
	if 0 == len(w.FilePath) {
		panic("file path i")
	}

	//判断字符串相等 也可以 使用 == 代替
	if !strings.EqualFold("123.txt", w.FilePath) {
		panic("the file is not found")
	}

	if 0 == len(contents) {
		panic("cant not write empty array byte for file")
	}

	return true
}

// Copy 接口 组合 Read和Write 接口
type Copy interface {
	Read
	Write
}

type CopyFile struct {
}

func ccc(c Copy) {
	contents := c.ReadFile()
	c.WriteFile(contents)
}

func main() {
	copy := Copy(ReadImpl{"123.txt"}, WriteImpl{"123.txt"})
	ccc(copy)
}
