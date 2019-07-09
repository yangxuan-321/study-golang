package main

import (
	"fmt"
	"regexp"
	"study-golang/main/crawler/engine"
)

func main() {
	aarseProfile1()
}

// 人物 内心独白 -- 简称OS
const peopleOSRegex = `<div class="m-title" data-v-bff6f798>内心独白</div> <div class="m-content-box m-des" data-v-bff6f798><span data-v-bff6f798>(([^<].*)[^>])</span><!----></div>`
const test = `<div class="m-title" data-v-bff6f798>内心独白</div> <div class="m-content-box m-des" data-v-bff6f798><span data-v-bff6f798>一直生活在成都。阿坝是填错的地方。自由职业。画故事书。生活内心都很干净。希望你有肌肉爱健身，居家。温暖。希望和你去爬一次武功山也想和你去看海还想和你每天买菜做饭看剧过家家。始终相信再迟也会遇到温柔岁月的那个人。所以，我不泡吧不抽烟不乱搞，我就每天好好赚钱做好保养控制好体重打理好自己慢慢的变好看变优秀慢慢的等你出现。</span><!----></div>`

func aarseProfile() engine.ParseResult {
	re := regexp.MustCompile(peopleOSRegex)
	submatch := re.FindStringSubmatch(test)
	for _, r := range submatch {
		fmt.Printf("%s\n", r)
	}
	return engine.ParseResult{}

}

const regex1 = `<div class="m-title" data-v-bff6f798>个人资料</div> ` +
	`<div class="m-content-box" data-v-bff6f798>` +
	`<div class="purple-btns" data-v-bff6f798>` +
	`<div class="m-btn purple" data-v-bff6f798>([^<].*)</div>` +
	`<div class="m-btn purple" data-v-bff6f798>([^<]\d+)岁</div>` +
	`<div class="m-btn purple" data-v-bff6f798>([^<].*)</div>` +
	`<div class="m-btn purple" data-v-bff6f798>([^<]\d+)cm</div>` +
	`<div class="m-btn purple" data-v-bff6f798>([^<]\d+)kg</div>` +
	`<div class="m-btn purple" data-v-bff6f798>工作地\:([^<].*)</div>` +
	`<div class="m-btn purple" data-v-bff6f798>月收入\:([^<].*)</div>` +
	`<div class="m-btn purple" data-v-bff6f798>([^<].*)</div>` +
	`<div class="m-btn purple" data-v-bff6f798>([^<].*)</div>` +
	`</div> ` +
	`<div class="pink-btns" data-v-bff6f798>` +
	`<div class="m-btn pink" data-v-bff6f798>([^<].*)</div>` +
	`<div class="m-btn pink" data-v-bff6f798>籍贯\:([^<].*)</div>` +
	`<div class="m-btn pink" data-v-bff6f798>体型\:([^<].*)</div>` +
	`<div class="m-btn pink" data-v-bff6f798>([^<].*)</div>` +
	`<div class="m-btn pink" data-v-bff6f798>([^<].*)</div>` +
	`<div class="m-btn pink" data-v-bff6f798>([^<].*)</div>` +
	`<div class="m-btn pink" data-v-bff6f798>([^<].*)</div>` +
	`<div class="m-btn pink" data-v-bff6f798>([^<].*)</div>` +
	`<div class="m-btn pink" data-v-bff6f798>是否想要孩子\:([^<].*)</div>` +
	`<div class="m-btn pink" data-v-bff6f798>何时结婚\:([^<].*)</div>` +
	`</div>` +
	`</div>`

func aarseProfile1() {
	test := `<div class="m-title" data-v-bff6f798>个人资料</div> ` +
		`<div class="m-content-box" data-v-bff6f798>` +
		`<div class="purple-btns" data-v-bff6f798>` +
		`<div class="m-btn purple" data-v-bff6f798>离异</div>` +
		`<div class="m-btn purple" data-v-bff6f798>26岁</div>` +
		`<div class="m-btn purple" data-v-bff6f798>天蝎座(10.23-11.21)</div>` +
		`<div class="m-btn purple" data-v-bff6f798>156cm</div>` +
		`<div class="m-btn purple" data-v-bff6f798>47kg</div>` +
		`<div class="m-btn purple" data-v-bff6f798>工作地:阿坝汶川</div>` +
		`<div class="m-btn purple" data-v-bff6f798>月收入:1.2-2万</div>` +
		`<div class="m-btn purple" data-v-bff6f798>画家</div>` +
		`<div class="m-btn purple" data-v-bff6f798>大学本科</div>` +
		`</div> ` +
		`<div class="pink-btns" data-v-bff6f798>` +
		`<div class="m-btn pink" data-v-bff6f798>汉族</div>` +
		`<div class="m-btn pink" data-v-bff6f798>籍贯:甘肃兰州</div>` +
		`<div class="m-btn pink" data-v-bff6f798>体型:保密</div>` +
		`<div class="m-btn pink" data-v-bff6f798>不吸烟</div>` +
		`<div class="m-btn pink" data-v-bff6f798>稍微喝一点酒</div>` +
		`<div class="m-btn pink" data-v-bff6f798>和家人同住</div>` +
		`<div class="m-btn pink" data-v-bff6f798>未买车</div>` +
		`<div class="m-btn pink" data-v-bff6f798>没有小孩</div>` +
		`<div class="m-btn pink" data-v-bff6f798>是否想要孩子:想要孩子</div>` +
		`<div class="m-btn pink" data-v-bff6f798>何时结婚:时机成熟就结婚</div>` +
		`</div>` +
		`</div>`
	compile := regexp.MustCompile(regex1)
	submatch := compile.FindStringSubmatch(test)[1:]
	fmt.Println(len(submatch))
}
