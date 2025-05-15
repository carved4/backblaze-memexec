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
	// Define flags
	var (
		args     string
		fileName string
	)

	flag.StringVar(&args, "args", "", "Comma-separated list of arguments to pass to the executable")
	flag.StringVar(&fileName, "file", DefaultExecutable, "Name of the executable to download and run (default: test.exe)")
	flag.Parse()

	// Process the command-line arguments
	var filePath string
	if flag.NArg() > 0 {
		// User provided a custom file path
		filePath = flag.Arg(0)
	} else {
		// Use the default or flag-specified file
		filePath = fileName
	}

	// Parse arguments if provided
	var execArgs []string
	if args != "" {
		execArgs = splitArgs(args)
	}

	// Print information about what we're doing
	fmt.Printf("Using Backblaze B2 bucket: %s\n", exec.DefaultBucket)
	fmt.Printf("Downloading and executing %s in memory\n", filePath)

	// Download and execute the file in memory
	err := pulldown.PulldownAndExec(filePath, execArgs)

	if err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("Done!")
}

// splitArgs splits a comma-separated string into a slice of arguments
func splitArgs(args string) []string {
	if args == "" {
		return nil
	}
	// Split the string by commas and trim whitespace from each argument
	rawArgs := strings.Split(args, ",")
	trimmedArgs := make([]string, 0, len(rawArgs))
	
	for _, arg := range rawArgs {
		trimmedArgs = append(trimmedArgs, strings.TrimSpace(arg))
	}
	
	return trimmedArgs
} 