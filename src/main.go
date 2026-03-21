package main

import (
	"fmt"
	"os"
    "path/filepath"
	"strconv"
	"time"
)

func outputPath(input string) string {
    ext := filepath.Ext(input)
    base := input
    if ext != "" {
        base = input[:len(input)-len(ext)]
    }
    return base + "-voxed.obj"
}

func main() {
	if len(os.Args) < 3 {
        fmt.Println("Usage: go run ./src <input.obj> <max-depth>")
        os.Exit(1)
    }

    input := os.Args[1]
    output := outputPath(input)

    depth, err := strconv.Atoi(os.Args[2])
    if err != nil || depth < 0 {
        fmt.Println("Error: <max-depth> can't be negative")
        os.Exit(1)
    }

}