package main

import (
	"fmt"
	"time"
)

func main() {
	// 获取每天的零点时间戳, 一个小时的时间戳是3600
	timeStr := time.Now().Format("2006-01-02")
	times, _ := time.ParseInLocation("2006-01-02", timeStr, time.Local)
	fmt.Println(times.Unix())

	timeStr = time.Now().Format("2006-01-02")
	t, _ := time.ParseInLocation("2006-01-02", timeStr, time.Local)
	timeUnix := t.Unix()
	fmt.Println(timeUnix)
}
