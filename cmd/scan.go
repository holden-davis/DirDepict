package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/holden-davis/DirDepict/util"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(scanCMD)
}

var scanCMD = &cobra.Command{
	Use: "scan",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			fmt.Println("Scan needs an argument!")
			return
		}
		rootpath, _ := filepath.Abs(args[0])
		_, err := os.Stat(rootpath)
		if err != nil {
			fmt.Println(args[0], "is not a valid file or directory!")
			return
		}
		root := scan(rootpath)
		root.Print()
	},
}

func scan(path string) util.File {
	root := util.File{}
	fi, _ := os.Stat(path)
	root.Fullpath = path
	root.Name = filepath.Base(path)
	root.Extension = filepath.Ext(path)
	if fi.IsDir() {
		entries, _ := os.ReadDir(path)
		root.IsDir = true
		for _, entry := range entries {
			root.Sub = append(root.Sub, scan(filepath.Join(root.Fullpath, entry.Name())))
		}
		for _, entry := range root.Sub {
			root.Size += entry.Size
		}
	} else {
		root.Size = fi.Size()
		root.IsDir = false
	}
	return root
}
