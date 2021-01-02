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

	if expr, err = cronexpr.Parse("*/5 * * * * * *"); err != nil {
		fmt.Println(err)
		return
	}

	//获取当前时间
	now = time.Now()
	//获取下次调度时间
	nextTime = expr.Next(now)
	//设置一个定时器和回调函数
	time.AfterFunc(nextTime.Sub(now), func() {
		fmt.Println("被调度了，下次执行时间：",nextTime)
	})

	time.Sleep(5 * time.Second)
}
