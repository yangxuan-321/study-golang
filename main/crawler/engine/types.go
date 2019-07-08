package engine

type Request struct {
	Url        string
	ParserFunc func([]byte) ParseResult
}

type ParseResult struct {
	// 返回的 将要载入 队列 接下来请求的相关请求
	Requests []Request
	// 解析出的内容列表
	Items []interface{}
}

// 什么都不做的 Parser
func NilParser(contents []byte) ParseResult {
	return ParseResult{}
}
