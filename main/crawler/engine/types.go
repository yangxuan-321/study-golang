package engine

type Request struct {
	Url        string
	ParserFunc func([]byte) (ParseResult, error)
}

type ParseResult struct {
	// 返回的 将要载入 队列 接下来请求的相关请求
	Requests []Request
	// 解析出的内容列表
	Items []Item
}

type Item struct {
	// 之所以把 Id Type Url 抽象出来，我们认为大多数爬虫都是需要 这些属性的。而Id和Type最好还是由 Item的本身来定，也就是parser来定。
	// ElasticSearch中的Id
	Id string
	// ElasticSearch中的表名
	Type    string
	Url     string
	Payload interface{}
}

// 什么都不做的 Parser
func NilParser(contents []byte) (ParseResult, error) {
	return ParseResult{}, nil
}
