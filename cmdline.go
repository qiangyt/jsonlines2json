package main

import (
	"fmt"
	"os"
)

const (
	// AppVersion ...
	AppVersion = "v0.9.0"
)

// PrintVersion ...
func PrintVersion() {
	fmt.Println(AppVersion)
}

// PrintHelp ...
func PrintHelp() {
	fmt.Println("\nTool to convert JSON lines to array of JSON objects")
	PrintVersion()
	fmt.Println()

	fmt.Println("Usage:")
	fmt.Println("  jl2ja  [option...] <your JSON line file path>")
	fmt.Println()

	fmt.Println("Options:")
	fmt.Printf("  -h,  --help       Display this information\n")
	fmt.Printf("  -V,  --version    Display app version information\n")
	fmt.Println()
}

// CommandLineT ...
type CommandLineT struct {
	LogFilePath string
}

// CommandLine ...
type CommandLine = *CommandLineT

// ParseCommandLine ...
func ParseCommandLine() (bool, CommandLine) {

	r := &CommandLineT{}

	for i := 0; i < len(os.Args); i++ {
		if i == 0 {
			continue
		}

		arg := os.Args[i]

		if arg[0:1] == "-" {
			if arg == "-h" || arg == "--help" {
				PrintHelp()
				return false, nil
			} else if arg == "-V" || arg == "--version" {
				PrintVersion()
				return false, nil
			} else {
				fmt.Printf("Unknown option: '%s'\n\n", arg)
				PrintHelp()
				return false, nil
			}
		} else {
			r.LogFilePath = arg
		}
	}

	if len(r.LogFilePath) == 0 {
		fmt.Printf("Please specify the json line file path")
		PrintHelp()
		return false, nil
	}

	return true, r
}
