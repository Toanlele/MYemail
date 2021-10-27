package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
	"time"

	"github.com/fatih/color"
)

func main() {
	logo()
	f, err := os.Open("data.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		anle := (scanner.Text())
		time.Sleep(time.Duration(2) * time.Second)
		qqcn := "@qq.com"
		qqmail := anle + qqcn

		color.Blue("正在给:%s发送邮件！", qqmail)
		cmd := exec.Command("bash", "mail.sh", qqmail)
		out, err := cmd.CombinedOutput()
		if err != nil {
			fmt.Printf("combined out:\n%s\n", string(out))
			log.Fatalf("cmd.Run() failed with %s\n", err)
		}
		fmt.Printf("combined out:\n%s\n", string(out))
		time.Sleep(time.Duration(2) * time.Second)
		color.Red("发送成功！")
	}

}

func logo() {

	beeStrUp := `
		____o__ __o____                                   
		/   \   /   \                                    
			 \o/                                         
			  |        o__ __o       o__ __o/  \o__ __o  
			 < >      /v     v\     /v     |    |     |> 
			  |      />       <\   />     / \  / \   / \ 
			  o      \         /   \      \o/  \o/   \o/ 
			 <|       o       o     o      |    |     |  
			 / \      <\__ __/>     <\__  / \  / \   / \ 	v1.7.0		
`
	beeStrDown := `
├── Version     : 1.7.0
├── GoVersion : go 1.6.2
├── GOOS      : Linux
├── GOARCH    : amd64
├── NumCPU    : 8
├── Anthor    :3143246477@qq.com
├── school    :重庆市建筑工程学院
└── Date      : Friday, 25 Aug 2021`
	color.Set(color.FgMagenta, color.Bold)
	defer color.Unset()
	fmt.Println(beeStrUp)
	color.Set(color.FgGreen, color.Bold)
	fmt.Println(beeStrDown)
}
