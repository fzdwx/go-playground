package main

import (
	"fmt"
	"net/http"
	"os"
	"os/exec"
)

func main() {

	http.HandleFunc("/cmd/", func(writer http.ResponseWriter, request *http.Request) {
		command := request.URL.Path[len("/cmd/"):]
		script := "/root/project/fzdwx/deploy.sh " + command
		cmd, err := Script(script)
		if err != nil {
			fmt.Println("script["+command+"]初始化失败  :", err)
			return
		}
		output, err := cmd.Output()
		if err != nil {
			fmt.Println("执行失败:", err)
			return
		}
		writer.Write(output)
	})
	err := http.ListenAndServe(":9999", nil)
	if err != nil {
		panic(err)
	}
	/*https://github.com/unknwon/the-way-to-go_ZH_CN/blob/master/eBook/15.3.md*/
}

// Script returns a command to execute a script through a shell.
func Script(script string) (*exec.Cmd, error) {
	shell := "/bin/sh"
	if other := os.Getenv("SHELL"); other != "" {
		shell = other
	}
	return exec.Command(shell, "-c", script), nil
}
