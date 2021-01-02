package main

import (
	"fmt"
	"github.com/gorhill/cronexpr"
)

func main() {
	var (
		expr *cronexpr.Expression
		err  error
	)
	//每分钟执行一次
	if expr, err = cronexpr.Parse("* * * * *123"); err != nil {
		fmt.Println(err)
		return
	}

	expr = expr
}
