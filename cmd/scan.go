package cmd

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(scanCMD)
}

var scanCMD = &cobra.Command{
	Use: "scan",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("<!---STARTING---!>")
		rootpath := filepath.Base(args[0])
		_, err := os.Stat(rootpath)
		if err == nil {
			FS := os.DirFS(rootpath)
			fs.WalkDir(FS, rootpath, func(path string, d fs.DirEntry, err error) error {
				if err != nil {
					fmt.Println(err)
				}
				fmt.Println(path)
				return nil
			})
		} else {
			fmt.Println(args[0] + " is an invalid path!")
		}
		fmt.Println("<!---STOPPING---!>")
	},
}
