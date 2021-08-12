package main

import (
	"fmt"
	"io/ioutil"
	"os/exec"
)

func main() {

	dateCmd := exec.Command("date")
	dateOut, err := dateCmd.Output()
	if err != nil {
		panic(err)
	}
	fmt.Println("> date")
	fmt.Println(string(dateOut))

	grepCommand := exec.Command("grep", "hello")

	grepIn, _ := grepCommand.StdinPipe()
	grepOut, _ := grepCommand.StdoutPipe()

	grepCommand.Start()

	grepIn.Write([]byte("hello grep\ngoodbye grep"))
	grepIn.Close()

	grepBytes, _ := ioutil.ReadAll(grepOut)

	grepCommand.Wait()

	fmt.Println("> grep hello")
	fmt.Println(string(grepBytes))

	lsCmd := exec.Command("bash", "-c", "ls -a -l -h")
	lsOut, err := lsCmd.Output()
	if err != nil {
		panic(err)
	}
	fmt.Println("> ls -a -l -h")
	fmt.Println(string(lsOut))
}
