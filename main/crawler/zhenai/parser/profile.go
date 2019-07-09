package parser

import (
	"fmt"
	"regexp"
	"strconv"
	"study-golang/main/crawler/engine"
	"study-golang/main/crawler/model"
)

// 人物 内心独白 -- 简称OS
const peopleOSRegex = `<div class="m-title" data-v-bff6f798>内心独白</div> <div class="m-content-box m-des" data-v-bff6f798><span data-v-bff6f798>([^<].*[^>])</span><!----></div>`

const peopleBaseInfoRegex = `<div class="m-title" data-v-bff6f798>个人资料</div> ` +
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

func ParseProfile(contents []byte) engine.ParseResult {

	profile := model.Profile{}

	// 解析独白
	re := regexp.MustCompile(peopleOSRegex)
	peopleOSInfo := re.FindSubmatch(contents)

	if len(peopleOSInfo) == 2 {
		profile.PeopleOS = string(peopleOSInfo[1])
	}

	// 解析其他个人信息
	parsePeopleBaseInfo()

	return engine.ParseResult{}
}

func parsePeopleBaseInfo(contents []byte, profile *model.Profile) error {
	// 解析独白
	re := regexp.MustCompile(peopleBaseInfoRegex)
	peopleBaseInfo := re.FindAllSubmatch(contents, -1)
	if len(peopleBaseInfo) != 1 {
		return fmt.Errorf("parse error, len(peopleBaseInfo) must 1, len(peopleBaseInfo):%d", len(peopleBaseInfo))
	}

	peopleBaseInfo0 := peopleBaseInfo[0]
	if len(peopleBaseInfo0)-1 < 19 {
		return fmt.Errorf("parse error, len(peopleBaseInfo0) must 19, len(peopleBaseInfo0):%d", len(peopleBaseInfo))
	}

	peopleBaseInfo0 = peopleBaseInfo0[1:]

	profile.Marriage = string(peopleBaseInfo0[0])

	age, e := strconv.Atoi(string(peopleBaseInfo0[1]))
	if nil != e {
		return fmt.Errorf("age is illegal:%s", string(age))
	}
	profile.Age = age

	profile.XingZuo = string(peopleBaseInfo0[2])

	height, e := strconv.Atoi(string(peopleBaseInfo0[3]))
	if nil != e {
		return fmt.Errorf("height is illegal:%s", string(height))
	}
	profile.Height = height

	weight, e := strconv.Atoi(string(peopleBaseInfo0[4]))
	if nil != e {
		return fmt.Errorf("weight is illegal:%s", string(weight))
	}
	profile.Weight = weight

	profile.Income = string(peopleBaseInfo0[6])

	profile.Occupation = string(peopleBaseInfo0[7])

	profile.Education = string(peopleBaseInfo0[8])

	profile.HuKou = string(peopleBaseInfo0[10])

	profile.Hourse = string(peopleBaseInfo0[14])

	profile.Car = string(peopleBaseInfo0[15])
}
