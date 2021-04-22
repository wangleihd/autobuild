package main

import (
	"fmt"
	"os"
	"os/exec"
	"time"
)

const (
	fileName string = "/build.log"
	gitPath  string = "./conf"
)

func main() {
	showtime := getDate()
	buildText := fmt.Sprintf("Auto build a commit by %s", showtime)

	add := fmt.Sprintf(" cd %s && git commit -a -m \"%s\"", gitPath, buildText)
	commit := fmt.Sprintf(" cd %s && git commit -a -m \"%s\"", gitPath, buildText)
	push := fmt.Sprintf(" cd %s && git commit -a -m \"%s\"", gitPath, buildText)

	writeFile(buildText)

	Command(add)
	Command(commit)
	Command(push)

	fmt.Println(showtime)
}

func writeFile(log string) {

	filename := fmt.Sprintf("%s%s", gitPath, fileName)

	file, err := os.OpenFile(filename, os.O_RDWR, 0666)
	defer file.Close()
	if err != nil {
		fmt.Println(err)
		return
	}

	file.WriteString(fmt.Sprintf("%s\n", log))

}

func getDate() string {
	timeUnix := time.Now().Unix()
	formatTimeStr := time.Unix(timeUnix, 0).Format("2006-01-02 15:04:05")

	return formatTimeStr

}

// 这里为了简化，我省去了stderr和其他信息
func Command(cmd string) error {
	c := exec.Command("bash", "-c", cmd)
	// 此处是windows版本
	// c := exec.Command("cmd", "/C", cmd)
	output, err := c.CombinedOutput()
	fmt.Println(string(output))
	return err
}
