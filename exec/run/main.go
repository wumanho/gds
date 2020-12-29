package main

import (
	"context"
	"fmt"
	"os/exec"
	"time"
)

type result struct {
	err error
	output []byte
}

func main() {
	var (
		ctx       context.Context
		cancelFun context.CancelFunc
		cmd       *exec.Cmd
		//该队列用于收集bash执行结果，发送到主协程
		resultChan chan *result
		res *result
	)
	//初始化结果队列
	resultChan = make(chan *result,1000)
	ctx, cancelFun = context.WithCancel(context.TODO())

	go func() {
		var (
			output []byte
			err    error
		)
		//执行任务，捕获输出结果
		cmd = exec.CommandContext(ctx, "/bin/bash", "-c", "sleep 5;echo 'hello'")
		output, err = cmd.CombinedOutput()
		//将结果投递到队列中
		resultChan <- &result{
			err:err,
			output: output,
		}
	}()

	time.Sleep(2 * time.Second)
	cancelFun()
	//从协程中获取数据，打印结果
	res  = <- resultChan
	fmt.Println(res.err,string(res.output))
}
