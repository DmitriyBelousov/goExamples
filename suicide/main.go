package main

import (
	"fmt"
	"os"
	"os/exec"
	"strconv"
)

func main() {
	pid := os.Getpid()
	str := strconv.Itoa(pid)
	fmt.Println("Process identifier: ", str)
	ret, _ := exec.Command("kill", "1", str).Output()
	fmt.Println("this will never be printed: ", ret)
}