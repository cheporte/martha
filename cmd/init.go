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
	Use:   "init",
	Short: "Initialize a new vault for your projects",
	Long:  "Sets up the initial structure for your project vault where Martha will store and manage your creative code journeys.",

	Run: func(cmd *cobra.Command, args []string) {
		if vaultPath == "" {
			usr, err := user.Current()
			if err != nil {
				fmt.Println("[Error] Could not get user home directory:", err)
				return
			}
			vaultPath = filepath.Join(usr.HomeDir, "martha-home")
		}

		err := os.MkdirAll(vaultPath, os.ModePerm)
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

		fmt.Println("Vault successfully created at:", vaultPath)
	},
}
