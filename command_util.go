package util

import (
	"bytes"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"fmt"
)

func ExecShell(s string) (string, error) {
	//函数返回一个*Cmd，用于使用给出的参数执行name指定的程序
	cmd := exec.Command("/bin/bash", "-c", s)

	//读取io.Writer类型的cmd.Stdout，再通过bytes.Buffer(缓冲byte类型的缓冲器)将byte类型转化为string类型(out.String():这是bytes类型提供的接口)
	var out bytes.Buffer
	cmd.Stdout = &out

	//Run执行c包含的命令，并阻塞直到完成。  这里stdout被取出，cmd.Wait()无法正确获取stdin,stdout,stderr，则阻塞在那了
	err := cmd.Run()
	checkErr(err)

	return out.String(), err
}

//错误处理函数
func checkErr(err error) {
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
}

func exec_shell_script(s string) (string, error) {
	//函数返回一个*Cmd，用于使用给出的参数执行name指定的程序

	stript := fmt.Sprintf(`tell application "Terminal" to do script "%s"`, s)
	cmd := exec.Command(`osascript`, "-s", "h", "-e", stript)
	//读取io.Writer类型的cmd.Stdout，再通过bytes.Buffer(缓冲byte类型的缓冲器)将byte类型转化为string类型(out.String():这是bytes类型提供的接口)
	var out bytes.Buffer
	cmd.Stdout = &out

	stderr, err := cmd.StderrPipe()
	log.SetOutput(os.Stderr)

	if err != nil {
		log.Fatal(err)
	}

	if err := cmd.Start(); err != nil {
		log.Fatal(err)
	}

	slurp, _ := ioutil.ReadAll(stderr)
	fmt.Printf("%s\n", slurp)

	if err := cmd.Wait(); err != nil {
		log.Fatal(err)
	}

	return string(slurp), err
}