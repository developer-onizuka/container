package main

import (
	"fmt"
	"os"
	"os/exec"
)

// docker          run image <cmd> <param>
// go run main.go  run       <cmd> <param>

func main() {
	switch os.Args[1] {
	case "run":
		run()

	default:
		panic("bad command")
	}
}

func run() {
        fmt.Printf("Running %v\n", os.Args[2:])

        cmd := exec.Command(os.Args[2], os.Args[3:]...)
        cmd.Stdin = os.Stdin
        cmd.Stdout = os.Stdout
        cmd.Stderr = os.Stderr

        cmd.Run()
}

