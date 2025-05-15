package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"memb2/pkg/exec"
	"memb2/pkg/pulldown"
)

const (
	// Default executable to run if no filename is provided
	DefaultExecutable = "test.exe"
)

func main() {
	var (
		args     string
		fileName string
	)

	flag.StringVar(&args, "args", "", "Comma-separated list of arguments to pass to the executable")
	flag.StringVar(&fileName, "file", DefaultExecutable, "Name of the executable to download and run (default: test.exe)")
	flag.Parse()


	var filePath string
	if flag.NArg() > 0 {
	
		filePath = flag.Arg(0)
	} else {
		filePath = fileName
	}

	var execArgs []string
	if args != "" {
		execArgs = splitArgs(args)
	}


	err := pulldown.PulldownAndExec(filePath, execArgs)

	if err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("Done!")
}

func splitArgs(args string) []string {
	if args == "" {
		return nil
	}
	
	rawArgs := strings.Split(args, ",")
	trimmedArgs := make([]string, 0, len(rawArgs))
	
	for _, arg := range rawArgs {
		trimmedArgs = append(trimmedArgs, strings.TrimSpace(arg))
	}
	
	return trimmedArgs
} 