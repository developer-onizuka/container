package main

import (
        "fmt"
        "os"
        "os/exec"
        "syscall"
)

// docker         run image <cmd> <param>
// go run main.go run       <cmd> <param>


func main() {
        switch os.Args[1] {
        case "run":
             run()
        case "child":
             child()
        default:
             panic("bad command")
        } 
}

func run() {
        fmt.Printf("Running %v\n as %d\n", os.Args[2:], os.Getpid())

	cmd := exec.Command("/proc/self/exe", append([]string{"child"}, os.Args[2:]...)...)
        cmd.Stdin = os.Stdin
        cmd.Stdout = os.Stdout
        cmd.Stderr = os.Stderr

        cmd.SysProcAttr = &syscall.SysProcAttr{
              Cloneflags: syscall.CLONE_NEWUTS | syscall.CLONE_NEWPID,
        }

	err := cmd.Run()
	if err != nil {
		panic(nil)
	}
}

func child() {
        fmt.Printf("Running %v\n as %d\n", os.Args[2:], os.Getpid())

        cmd := exec.Command(os.Args[2], os.Args[3:]...)
        cmd.Stdin = os.Stdin
        cmd.Stdout = os.Stdout
        cmd.Stderr = os.Stderr

	syscall.Sethostname([]byte("container"))

	err := cmd.Run()
	if err != nil {
		panic(nil)
	}
}
