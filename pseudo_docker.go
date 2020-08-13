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
        case "stop":
             stop()
        case "rm":
             rm()
        case "commit":
             commit()
        default:
             panic("bad command")
        } 
}

func run() {
        fmt.Printf("Running %v as %d\n", os.Args[2:], os.Getpid())

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
        fmt.Printf("Running %v as %d\n", os.Args[2:], os.Getpid())
	deploy := exec.Command("tar","xvfz","/home/pseudo_docker/myimages/myroot.tar.gz","-C","/home/pseudo_docker")
	deploy.Run()

        cmd := exec.Command(os.Args[2], os.Args[3:]...)
        cmd.Stdin = os.Stdin
        cmd.Stdout = os.Stdout
        cmd.Stderr = os.Stderr

	syscall.Sethostname([]byte("container"))
	syscall.Chroot("/home/pseudo_docker/myroot")
	syscall.Chdir("/")
	syscall.Mount("proc", "proc", "proc", 0, "")

	err := cmd.Run()
	if err != nil {
		panic(nil)
	}
	syscall.Unmount("/proc", 0)
}

func stop() {
        fmt.Printf("Stopping now ...\n")
        stop := exec.Command("tar","cvfz","/home/pseudo_docker/exit/myroot.tar.gz","-C","/home/pseudo_docker","myroot")
	stop.Run()
}

func rm() {
        fmt.Printf("Removing now ...\n")
        rm := exec.Command("rm","-r","/home/pseudo_docker/myroot")
	rm.Run()
}

func commit() {
        fmt.Printf("Committing now ...\n")
        commit := exec.Command("mv","/home/pseudo_docker/exit/myroot.tar.gz","/home/pseudo_docker/myimages")
	commit.Run()
}
