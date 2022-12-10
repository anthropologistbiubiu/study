package main

import (
	"context"
	"encoding/json"
	"fmt"
	"reflect"

	"github.com/olivere/elastic/v7"
)

var client *elastic.Client

var host = "http://localhost:9200"

type Employee struct {
	FirstName string   `json:"first_name"`
	LastName  string   `json:"last_name"`
	Age       int      `json:"age"`
	About     string   `json:"about"`
	Interests []string `json:"interests"`
}

// 初始化
func init() {
	//errorlog := log.New(os.Stdout, "APP", log.LstdFlags)
	var err error
	//这个地方有个小坑 不加上elastic.SetSniff(false) 会连接不上
	client, err = elastic.NewClient(elastic.SetSniff(false), elastic.SetURL(host))
	if err != nil {
		panic(err)
	}
	_, _, err = client.Ping(host).Do(context.Background())
	if err != nil {
		panic(err)
	}

	_, err1 := client.ElasticsearchVersion(host)
	if err != nil {
		panic(err1)
	}
	//fmt.Printf("Elasticsearch version %s\n", esversion)
}

// 创建
func create() {
	//使用结构体
	e1 := Employee{"Jane", "Smith", 32, "I like to collect rock albums", []string{"music"}}
	put1, err := client.Index().
		Index("megacorp").
		Type("employee").
		Id("1").
		BodyJson(e1).
		Do(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("Indexed tweet %s to index s%s, type %s\n", put1.Id, put1.Index, put1.Type)

	//使用字符串
	e2 := `{"first_name":"John","last_name":"Smith","age":25,"about":"I love to go rock climbing","interests":["sports","music"]}`
	put2, err := client.Index().
		Index("megacorp").
		Type("employee").
		Id("2").
		BodyJson(e2).
		Do(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("Indexed tweet %s to index s%s, type %s\n", put2.Id, put2.Index, put2.Type)

	e3 := `{"first_name":"Douglas","last_name":"Fir","age":35,"about":"I like to build cabinets","interests":["forestry"]}`
	put3, err := client.Index().
		Index("megacorp").
		Type("employee").
		Id("3").
		BodyJson(e3).
		Do(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("Indexed tweet %s to index s%s, type %s\n", put3.Id, put3.Index, put3.Type)
}
func gets() {
	//通过id查找
	for i := 2; i < 4; i++ {
		id := fmt.Sprintf("%d", i)
		get1, err := client.Get().Index("megacorp").Type("employee").Id(id).Do(context.Background())
		if err != nil {
			panic(err)
		}
		if get1.Found {
			//fmt.Printf("Got document %s in version %d from index %s, type %s\n", get1.Id, get1.Version, get1.Index, get1.Type)
			var bb Employee
			err := json.Unmarshal(get1.Source, &bb)
			if err != nil {
				fmt.Println(err)
			}
			fmt.Println(bb)
		}

	}
}

func delete() {
	res, err := client.Delete().Index("megacorp").
		Type("employee").
		Id("1").
		Do(context.Background())
	if err != nil {
		println(err.Error())
		return
	}
	fmt.Printf("delete result %s\n", res.Result)
}

func update() {
	res, err := client.Update().
		Index("megacorp").
		Type("employee").
		Id("2").
		Doc(map[string]interface{}{"first_name": "sun", "last_name": "weiming", "age": 24}).
		Do(context.Background())
	if err != nil {
		println(err.Error())
	}
	fmt.Printf("update age %s\n", res.Result)
}

// //搜索
func query() {
	var res *elastic.SearchResult
	var err error
	//取所有
	res, err = client.Search("megacorp").Type("employee").Do(context.Background())
	if err != nil {
		fmt.Println(err)
	}
	printEmployee(res, err)
	////字段相等
	sq := elastic.NewQueryStringQuery("last_name:weiming")
	res, err = client.Search("megacorp").Type("employee").Query(sq).Do(context.Background())
	if err != nil {
		println(err.Error())
	}
	printEmployee(res, err)
	//条件查询
	//年龄大于30岁的
	boolQ := elastic.NewBoolQuery()
	mq := boolQ.Must(elastic.NewMatchQuery("last_name", "weiming"))
	res, err = client.Search("megacorp").Type("employee").Query(mq).Do(context.Background())
	printEmployee(res, err)
	//短语搜索 搜索about字段中有 rock climbing
	matchPhraseQuery := elastic.NewMatchPhraseQuery("about", "rock climbing")
	res, err = client.Search("megacorp").Type("employee").Query(matchPhraseQuery).Do(context.Background())
	printEmployee(res, err)

	//分析 interests
	aggs := elastic.NewTermsAggregation().Field("interests")
	fmt.Println(aggs)
	fmt.Println("___________________________________________________________________________")
	res, err = client.Search("megacorp").Type("employee").Aggregation("all_interests", aggs).Do(context.Background())
	printEmployee(res, err)

}

// //简单分页
func list(size, page int) {
	if size < 0 || page < 1 {
		fmt.Printf("param error")
		return
	}
	res, err := client.Search("megacorp").
		Type("employee").
		Size(size).
		From((page - 1) * size).
		Do(context.Background())
	printEmployee(res, err)

}

// 打印查询到的Employee
func printEmployee(res *elastic.SearchResult, err error) {
	//fmt.Printf("%+v\n", res)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	var typ Employee
	for _, item := range res.Each(reflect.TypeOf(typ)) { //从搜索结果中取数据的方法
		t := item.(Employee)
		fmt.Printf("%#v\n", t)
	}
}

func main() {
	query()
}
