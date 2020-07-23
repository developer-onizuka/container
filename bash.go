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
        cmd.SysProcAttr = &syscall.SysProcAttr{
              Cloneflags: syscall.CLONE_NEWUTS,
        }

//	syscall.Sethostname([]byte("container"))

	err := cmd.Run()
	if err != nil {
		panic(nil)
	}

        state := cmd.ProcessState
        fmt.Printf("%s\n", state.String())
        fmt.Printf(" Pid: %d\n", state.Pid())
        fmt.Printf(" Exited: %v\n", state.Exited())
        fmt.Printf(" Success: %v\n", state.Success())
        fmt.Printf(" System: %v\n", state.SystemTime())
        fmt.Printf(" User: %v\n", state.UserTime())

}

