package parser

import (
	"log"
	"regexp"
	"study-golang/main/crawler/engine"
)

// 城市列表解析器

const citysRegex = `<a href="(http://www.zhenai.com/zhenghun/[0-9a-z]+)"[^>]*>([^<]+)</a>`

func ParseCityList(contents []byte) engine.ParseResult {
	// 使用 `` 可以 保证 不会涉及到复杂的 转义。 例如 如果 用 "" 那么 在
	// 在正则里面 需要 用到 . 那么就要用 \\. 而在`` 只需要 \.
	// ^> 代表 只要不是 右括号
	compile := regexp.MustCompile(citysRegex)
	all := compile.FindAllSubmatch(contents, -1)
	result := engine.ParseResult{}
	limit := 10
	for _, a := range all {
		// 放入 城市的名字
		result.Items = append(result.Items, "City: "+string(a[2]))
		result.Requests = append(result.Requests, engine.Request{Url: string(a[1]), ParserFunc: ParseCity})
		limit--
		if limit == 0 {
			break
		}
	}

	// 找到 470 个
	log.Printf("count:%d", len(all))

	return result
}
