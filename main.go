package main

import (
	"fmt"
	"os"

	"github.com/cheporte/martha/cmd"
)

func main() {
	if err := cmd.Execute(); err != nil {
		fmt.Println("[Error]", err)
		os.Exit(1)
	}
}
