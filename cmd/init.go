package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"os/user"

	"github.com/spf13/cobra"
)

var vaultPath string

var initCmd = &cobra.Command{
	Use:   "bloom",
	Short: "Plant the first seed of your vault 🌱",
	Long: `bloom plants the roots of your creative sanctuary — a vault where Martha
will tend to your projects with care, quiet strength, and just a touch of memory.`,

	Run: func(cmd *cobra.Command, args []string) {
		usr, err := user.Current()
		if vaultPath == "" {
			if err != nil {
				fmt.Println("[Error] Couldn't get user home directory:", err)
				return
			}
			vaultPath = filepath.Join(usr.HomeDir, "martha-home")
		} else if vaultPath == "." {
			absPath, err := os.Getwd()
			if err != nil {
				fmt.Println("[Error] Couldn't resolve current directory:", err)
				return
			}
			vaultPath = absPath
		} else {
			vaultPath, err = filepath.Abs(vaultPath)
			if err != nil {
				fmt.Println("[Error] Couldn't resolve given path:", err)
				return
			}
		}

		err = os.MkdirAll(vaultPath, os.ModePerm)
		if err != nil {
			fmt.Println("[Error] Could not create vault directory:", err)
			return
		}

		configFile := filepath.Join(vaultPath, ".martha")
		_, err = os.Create(configFile)
		if err != nil {
			fmt.Println("[Warning] Vault created, but failed to write config file:", err)
		} else {
			fmt.Println("Config file created successfully.")
		}

		fmt.Println()
		fmt.Println("🌸 Your vault has bloomed at:", filepath.Dir(vaultPath))
		fmt.Println("📜 Martha will remember this path, even if the world forgets.")
	},
}
