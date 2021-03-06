package main

import (
	"fmt"
	"os"
	"os/exec"
	"time"
)

const (
	fileName string = "/build.txt"
	gitPath  string = "/data/limitoo"
)

func main() {
	showtime := getDate()
	buildText := fmt.Sprintf("Auto build %s", showtime)

	pull := fmt.Sprintf(" cd %s && git pull", gitPath)
	add := fmt.Sprintf(" cd %s && git add .", gitPath)
	commit := fmt.Sprintf(" cd %s && git commit -a -m \"%s\"", gitPath, buildText)
	push := fmt.Sprintf(" cd %s && git push", gitPath)

	writeFile(buildText)

	Command(pull)
	Command(add)
	Command(commit)
	Command(push)

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
