package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/mitchellh/go-ps"
)

type TProcess struct {
	ps.Process
	UpperExecutable string
}

func line(n int) string {
	if n == 0 {
		return ""
	}
	if n == 1 {
		return "`-"
	}
	return strings.Repeat(" ", n-1) + "`-"
}

func mains(args []string) error {
	processes, err := ps.Processes()
	if err != nil {
		return err
	}
	processMap := map[int]*TProcess{}
	for _, p := range processes {
		processMap[p.Pid()] = &TProcess{
			Process:         p,
			UpperExecutable: strings.ToUpper(p.Executable()),
		}
	}
	seperator := ""
	for _, exeName := range args {
		upperExeName := strings.ToUpper(exeName)
		for _, p := range processMap {
			if strings.HasPrefix(p.UpperExecutable, upperExeName) {
				q := p
				indent := 0
				fmt.Print(seperator)
				seperator = "\n"
				for {
					fmt.Printf("%s%-5d %-5d %s\n",
						line(indent),
						q.Process.Pid(),
						q.Process.PPid(),
						q.Process.Executable())
					var ok bool
					q, ok = processMap[q.PPid()]
					if !ok {
						break
					}
					indent += 2
				}
			}
		}
	}
	return nil
}

func main() {
	if err := mains(os.Args[1:]); err != nil {
		fmt.Println(os.Stderr, err.Error())
		os.Exit(1)
	}
}
