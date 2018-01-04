package main

import (
	"fmt"
	"os"
	"os/exec"
	"time"
)

func main() {
	if len(os.Args) == 1 {
		fmt.Println("No argument was provided")
		os.Exit(1)
	}
	var cmd *exec.Cmd
	_, err := exec.LookPath(os.Args[1])
	if err != nil {
		cmd = exec.Command("cmd.exe", "/C")
		cmd.Args = append(cmd.Args, os.Args[1:]...)
	} else {
		cmd = exec.Command(os.Args[1], os.Args[2:]...)
	}
	start := time.Now()
	cmdOut, _ := cmd.CombinedOutput()
	/*if err != nil {
		fmt.Printf("Error executing command: %v", err)
		os.Exit(1)
	}*/
	//todo fix russian encoding in windows
	fmt.Printf("%s\n", cmdOut)
	fmt.Printf("Time elapsed: %v", time.Since(start))
}
