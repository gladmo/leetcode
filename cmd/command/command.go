package command

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var (
	Version = "v1.0.0"
)

var cmd = &cobra.Command{
	Use:   "leetcode",
	Short: "leetcode cli",
}

func init() {
	cmd.AddCommand(versionCmd, getCmd, clearCmd, infoCmd, testCmd, baseCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of leetcode cli",
	Long:  `All software has versions. This is leetcode's`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("leetcode cli", Version)
	},
}

// Execute cmd entrance
func Execute() {
	if err := cmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
