package main

import (
	"fmt"
	"regexp"
	"study-golang/main/crawler/engine"
	"unicode/utf8"
)

func main() {
	aarseProfile()
}

// 人物 内心独白 -- 简称OS
const peopleOSRegex = `<div class="m-title" data-v-bff6f798>内心独白</div> <div class="m-content-box m-des" data-v-bff6f798><span data-v-bff6f798>(.*)</span></div>`
const test = `<span data-v-bff6f798>一直生活在成都。阿坝是填错的地方。自由职业。画故事书。生活内心都很干净。希望你有肌肉爱健身，居家。温暖。希望和你去爬一次武功山也想和你去看海还想和你每天买菜做饭看剧过家家。始终相信再迟也会遇到温柔岁月的那个人。所以，我不泡吧不抽烟不乱搞，我就每天好好赚钱做好保养控制好体重打理好自己慢慢的变好看变优秀慢慢的等你出现。</span>`

func aarseProfile() engine.ParseResult {
	re := regexp.MustCompile(peopleOSRegex)
	submatch := re.FindSubmatch(utf8.DecodeRuneInString())
	for _, r := range submatch {
		fmt.Printf("%s", string(r[1]))
		fmt.Println()
		fmt.Printf("%s", "123")
	}
	return engine.ParseResult{}

}
