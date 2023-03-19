package main

import "fmt"

func main1() {
	dict := map[string]int{"key1": 1, "key2": 2}
	if value, ok := dict["key1"]; ok {
		fmt.Println(value)
	} else {
		fmt.Println("key1 不存在")
	}
}

type Student struct {
	name string
	id   int
}

func main() {
	i := make([]interface{}, 3) // 定义 interface 数组
	i[0] = 1                    //int
	i[1] = "hello go"           //string
	i[2] = Student{"mike", 18}  //Student

	//类型查询，类型断言
	//第一个返回下标，第二个返回下标对应的值，data分别是i[0],i[1],i[2]
	for _, data := range i {
		switch value := data.(type) {
		case int:
			_, ok := data.(int)
			fmt.Println(ok, value)
			//fmt.Printf("x[%d] 类型为int, 内容为%d\n", index, value)
		case string:
			_, ok := data.(string)
			fmt.Println(value, ok)
			//fmt.Printf("x[%d] 类型为string, 内容为%s\n", index, value)
		case Student:
			v, ok := data.(Student)
			fmt.Println(v, ok)
			//fmt.Printf("x[%d] 类型为Student, 内容为%s\n", index, value.name)
		}
	}
}
