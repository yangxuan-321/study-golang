package parser

import (
	"regexp"
	"study-golang/main/crawler/engine"
)

// 城市解析器

//const cityUserRegex = `<a href="http://album.zhenai.com/u/([0-9]+)[^>].*">([^<].*[^>].*)</a>`
const cityUserRegex = `<a href="(http://album.zhenai.com/u/[0-9]+)" target="_blank">([^<]*)</a>`

var cityUserRe = regexp.MustCompile(cityUserRegex)

const cityUrlRegex = `<a href="(http://www.zhenai.com/zhenghun/aba/[^"]+)">`

var cityUrlRe = regexp.MustCompile(cityUrlRegex)

func ParseCity(contents []byte) (engine.ParseResult, error) {
	all := cityUserRe.FindAllSubmatch(contents, -1)
	result := engine.ParseResult{}
	for _, m := range all {
		// 放入 用户的名字
		//result.Items = append(result.Items, "User: "+string(m[2]))
		result.Requests = append(result.Requests,
			engine.Request{
				Url: string(m[1]),
				ParserFunc: func(bytes []byte) (engine.ParseResult, error) {
					return ParseProfile(bytes, string(m[2]))
				}})
	}

	cityUrls := cityUrlRe.FindAllSubmatch(contents, -1)
	for _, m := range cityUrls {
		result.Requests = append(result.Requests, engine.Request{
			Url:        string(m[1]),
			ParserFunc: ParseCity,
		})
	}
	return result, nil
}
