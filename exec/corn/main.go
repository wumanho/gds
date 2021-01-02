package main

import (
	"fmt"
	"github.com/gorhill/cronexpr"
	"time"
)

func main() {
	var (
		expr *cronexpr.Expression
		err  error
		now time.Time
		nextTime time.Time
	)

	if expr, err = cronexpr.Parse("*/6 * * * *"); err != nil {
		fmt.Println(err)
		return
	}

	//获取当前时间
	now = time.Now()
	//获取下次调度时间
	nextTime = expr.Next(now)
	fmt.Printf("当前时间:%v\n 下次调度时间:%v",now,nextTime)
}
