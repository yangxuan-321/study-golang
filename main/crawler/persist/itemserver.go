package persist

import (
	"context"
	"fmt"
	"gopkg.in/olivere/elastic.v5"
)

func ItemServer() chan interface{} {
	out := make(chan interface{})
	go func() {
		itemCount := 0
		for {
			item := <-out
			itemCount += 1
			fmt.Printf("Got item: %d th, info = %v", itemCount, item)
		}
	}()
	return out
}

// 数据库
const Index = "dating_profile"

// 表
const Type = "zhenai"

// elasticsearch 的 host
const Host = "http://eat_chat:9200/"

//使用elasticsearch的客户端来保存数据
//选择客户端 https://github.com/olivere/elastic
//安装 go get -v gopkg.in/olivere/elastic.v5
//导入 import "gopkg.in/olivere/elastic.v5"
func save(item interface{}) (id string, err error) {
	client, err := elastic.NewClient(
		elastic.SetURL(Host),
		// Must trun off sniff in docker
		elastic.SetSniff(false),
	)
	if nil != err {
		return "", err
	}

	// Index 动作就是存数据的意思
	//client.Index().Index(Index).Type(Type).Id().BodyJson()
	// 不加Id，Id会自动生成
	reponse, err := client.Index().
		Index(Index).
		Type(Type).
		BodyJson(item).
		Do(context.Background())
	if nil != err {
		return "", err
	}

	fmt.Printf("%+v", reponse)
	return reponse.Id, nil
}
