package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"syscall"
	"time"
)

var numChild = flag.Int(
	"numChild",
	10,
	"hi",
)

func main() {
	flag.Parse()
	fmt.Println(os.Args)
	wd, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	fmt.Println("wd:", wd)
	p := "/home/vcap/app/dumb"

	if *numChild > 0 {
		pid, err := syscall.ForkExec(p, []string{p, "-numChild", strconv.Itoa(*numChild - 1)}, &syscall.ProcAttr{
			Env: os.Environ(),
			Dir: wd,
		})

		if err != nil {
			panic(err)
		}
		fmt.Println("i made a first pid! it is:", pid)

		pid, err = syscall.ForkExec(p, []string{p, "-numChild", strconv.Itoa(*numChild - 1)}, &syscall.ProcAttr{
			Env: os.Environ(),
			Dir: wd,
		})

		if err != nil {
			panic(err)
		}
		fmt.Println("i made a second pid! it is:", pid)
	}

	for {
		fmt.Println("heartbeat")
		time.Sleep(time.Second)
	}
}
