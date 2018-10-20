package main

import (
	"bytes"
	"fmt"
	"io"
	"os/exec"
)

func main() {
	cmd := exec.Command("echo", "-n", "My first command comes from golang")

	pipe, err := cmd.StdoutPipe()
	if err != nil {
		fmt.Printf("error: cannot obtain the stdout pipe:%s\n", err)
		return
	}

	if err := cmd.Start(); err != nil {
		fmt.Printf("error: The command cannot be startup:%s\n", err)
		return
	}

	//方式一
	var data bytes.Buffer
	for {
		tmp := make([]byte, 5)
		n, err := pipe.Read(tmp)
		if err != nil {
			if io.EOF == err {
				break
			} else {
				fmt.Printf("error: cannot read data from pipe:%s\n", err)
				return
			}
		}
		if n > 0 {
			fmt.Printf("%s\n", tmp)
			data.Write(tmp[:n])
		}
	}
	fmt.Printf("%s\n", data.String())

	//方式二
	// reader := bufio.NewReader(pipe)
	// data, _, err := reader.ReadLine()
	// if err != nil {
	// 	fmt.Printf("error: cannot read data from the pipe:%s\n", err)
	// 	return
	// }
	// fmt.Printf("%s\n", string(data))
}
