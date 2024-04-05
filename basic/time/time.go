package main

import (
	"fmt"
	"log"
	"time"
)

func main() {
	t := time.Now().Unix()      //外部传入的时间戳（秒为单位），必须为int64类型
	t1 := "2019-01-08 13:50:30" //外部传入的时间字符串

	//时间转换的模板，golang里面只能是 "2006-01-02 15:04:05"
	timeTemplate1 := "2006-01-02 15:04:05" //常规类型
	timeTemplate2 := "2006/01/02 15:04:05" //其他类型
	timeTemplate3 := "2006-01-02"          //其他类型
	timeTemplate4 := "15:04:05"            //其他类型

	// ======= 将时间戳格式化为日期字符串 =======
	log.Println(time.Unix(t, 0).Format(timeTemplate1)) //输出：2019-01-08 13:50:30
	log.Println(time.Unix(t, 0).Format(timeTemplate2)) //输出：2019/01/08 13:50:30
	log.Println(time.Unix(t, 0).Format(timeTemplate3)) //输出：2019-01-08
	log.Println(time.Unix(t, 0).Format(timeTemplate4)) //输出：13:50:30
	fmt.Println(time.Now().Format(timeTemplate1))

	// ======= 将时间字符串转换为时间戳 =======
	stamp, _ := time.ParseInLocation(timeTemplate1, t1, time.Local) //使用parseInLocation将字符串格式化返回本地时区时间
	log.Println(stamp.Unix())                                       //输出：1546926630
}
