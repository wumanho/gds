package main

import (
	"fmt"
	"os/exec"
)

func main() {
	var (
		cmd *exec.Cmd
		err error
		output []byte
	)
	cmd = exec.Command("/bin/bash", "-c", "sleep 5;ls -l")

	if output,err = cmd.CombinedOutput();err != nil{
		fmt.Println(err)
	}

	fmt.Println(string(output))
}
