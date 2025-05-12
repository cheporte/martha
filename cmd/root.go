package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use: "martha",
	Short: "Martha - She, who gives life to your ideas",
	Long: `Martha is a CLI tool that helps you create, organize,
and manage a vault of code projects.

Built to keep memory, emotion, and structure all in one place.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("🌸 Martha is here. Type `martha help` to see what she can do.")
	},
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println("[Error]", err)
		os.Exit(1)
	}
}