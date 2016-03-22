package main

import (
	"fmt"
	"os"
)

// PrintEnv print env
func PrintEnv() {
	for _, item := range os.Environ() {
		fmt.Println(item)
	}
}

func main() {
	fmt.Println(os.Getenv("TERM"))
	fmt.Println(os.Getenv("HOME"))
	fmt.Println(os.Getenv("RUN_TERM"))

	//PrintEnv()
}
