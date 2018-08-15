package main

import (
	"errors"
	"fmt"
	"math"
	"os"
	"os/exec"
)

func main() {
	deferpanicrecover()
	subprocessexec()
}

func ConvertInt64ToInt(a int64) int {
	if math.MinInt32 <= a && a <= math.MaxInt32 {
		return int(a)
	}
	panic("too big")
}

func IntFromInt64(a int64) (p int, err error) {
	defer func() {
		if e := recover(); e != nil {
			err = errors.New("123")
		}
	}()
	p = ConvertInt64ToInt(a)
	return p, nil
}

func deferpanicrecover() {
	l := int64(456123)
	if i, err := IntFromInt64(l); err != nil {
		fmt.Printf("The conversion of %d to an int32 resulted in an error: %s\n", l, err.Error())
	} else {
		fmt.Printf("%d converted to an int32 is %d\n", l, i)
	}
}

func subprocess() {
	env := os.Environ()
	procAttr := &os.ProcAttr{
		Env: env,
		Files: []*os.File{
			os.Stdin,
			os.Stdout,
			os.Stderr,
		},
	}
	fmt.Println(os.Getpid())
	pid, err := os.StartProcess("/bin/ls", []string{"-l"}, procAttr)
	if err != nil {
		fmt.Printf("Error %v starting process!\n", err) //
		os.Exit(1)
	}
	fmt.Printf("The process id is %v\n", pid)

	fmt.Println()

	pid, err = os.StartProcess("/bin/ps", []string{"-e", "-opid,ppid,comm"}, procAttr)
	if err != nil {
		fmt.Printf("Error %v starting process!\n", err) //
		os.Exit(1)
	}
	fmt.Printf("The process id is %v\n", pid)
}

func execute() {
	cmd := exec.Command("ls", "-al") // this opens a gedit-window
	err := cmd.Run()
	if err != nil {
		fmt.Printf("Error %v executing command!", err)
		os.Exit(1)
	}
	fmt.Printf("The command is %v", cmd)
}

func subprocessexec() {
	subprocess()
	execute()
}
