package main

import (
	"bufio"
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
	cmd := exec.Command(os.Args[1], os.Args[2:]...)
	cmdOut, err := cmd.StdoutPipe()
	if err != nil {
		fmt.Printf("Error get output pipe %q", err)
		os.Exit(1)
	}
	start := time.Now()
	if err = cmd.Start(); err != nil {
		fmt.Printf("Error run command: %q", err)
		os.Exit(1)
	}
	sc := bufio.NewScanner(cmdOut)
	for sc.Scan() {
		fmt.Println(sc.Text())
	}
	if err = cmd.Wait(); err != nil {
		fmt.Printf("Error wait command: %q", err)
	}
	//todo fix russian encoding in windows
	fmt.Printf("Time elapsed: %v", time.Since(start))
}
