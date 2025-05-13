package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	rootCmd = &cobra.Command{
		Use:   "martha",
		Short: "Martha - She, who gives life to your ideas",
		Long: `Martha is a CLI tool that helps you create, organize,
and manage a vault of code projects. Built to keep memory, emotion, and structure all in one place.`,

		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("🌸 Martha is here. Type `martha help` to see what she can do.")
		},
	}
)

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	rootCmd.AddCommand(initCmd)

	initCmd.Flags().StringVarP(&vaultPath, "vault", "v", "", "Custom path for your project vault")
}
