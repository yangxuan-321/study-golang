package parser

import (
	"regexp"
	"study-golang/main/crawler/engine"
)

// 城市解析器

//const cityUserRegex = `<a href="http://album.zhenai.com/u/([0-9]+)[^>].*">([^<].*[^>].*)</a>`
const cityUserRegex = `<a href="(http://album.zhenai.com/u/[0-9]+)" target="_blank">([^<]*)</a>`

var cityUserRe = regexp.MustCompile(cityUserRegex)

func ParseCity(contents []byte) engine.ParseResult {
	all := cityUserRe.FindAllSubmatch(contents, -1)
	result := engine.ParseResult{}
	for _, m := range all {
		// 放入 用户的名字
		result.Items = append(result.Items, "User: "+string(m[2]))
		result.Requests = append(result.Requests,
			engine.Request{
				Url: string(m[1]),
				ParserFunc: func(bytes []byte) engine.ParseResult {
					return ParseProfile(bytes, string(m[2]))
				}})
	}

	return result
}
