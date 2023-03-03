package cmd

import (
	"fmt"
	"os"

	"github.com/holden-davis/DirDepict/util"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(scanCMD)
}

var scanCMD = &cobra.Command{
	Use: "scan",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("<!---STARTING---!>")
		_, err := os.ReadDir(args[0])
		if err != nil {
			fmt.Println(args[0] + " is not a valid directory!")
		} else {
			rootentry, _ := os.Stat(args[0])
			root := util.File{}
			root.Size = rootentry.Size()
			fmt.Println(root)
		}
		fmt.Println("<!---STOPPING---!>")
	},
}
