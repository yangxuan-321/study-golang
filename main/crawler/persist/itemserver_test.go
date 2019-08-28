package persist

import (
	"context"
	"encoding/json"
	"fmt"
	"gopkg.in/olivere/elastic.v5"
	"study-golang/main/crawler/model"
	"testing"
)

func TestSave(t *testing.T) {
	profile := model.Profile{
		Age:        34,
		Height:     162,
		Weight:     57,
		Income:     "3001-5000元",
		Gender:     "女",
		Name:       "安静的雪",
		XingZuo:    "牡羊座",
		Occupation: "人事/行政",
		Marriage:   "离异",
		Hourse:     "已购房",
		HuKou:      "山东菏泽",
		Education:  "大学本科",
		Car:        "未购车",
		PeopleOS:   "东边日出西边雨, 道却无情似有情",
	}

	id, err := save(profile)

	if nil != err {
		panic(err)
	}

	client, err := elastic.NewClient(
		elastic.SetURL(Host),
		elastic.SetSniff(false),
	)
	if nil != err {
		panic(err)
	}

	reponse, err := client.Get().Index(Index).Type(Type).Id(id).Do(context.Background())
	if nil != err {
		panic(err)
	}

	//fmt.Printf("%+v", reponse)
	var actual model.Profile
	//bytes, err := reponse.Source.MarshalJSON()
	if nil != err {
		panic(err)
	}
	err = json.Unmarshal(*reponse.Source, &actual)

	if nil != err {
		panic(err)
	}

	if actual != profile {
		panic("------------not same------------")
	} else {
		fmt.Printf("------good - pass-------")
	}
}
