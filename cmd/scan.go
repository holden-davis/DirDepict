package cmd

import (
	"fmt"
	"github.com/holden-davis-uca/DirDepict/util"
	"github.com/spf13/cobra"
	"os"
	"path"
	"path/filepath"
)

func init() {
	rootCmd.AddCommand(scanCMD)
}

var scanCMD = &cobra.Command{
	Use: "scan",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("STARTING\n")
		abspath, _ := filepath.Abs(args[0])
		_, err := os.ReadDir(abspath)
		if err == nil {
			startingdir := recursiveScan(abspath)
			fmt.Println(startingdir.Name)
			fmt.Println(startingdir.Fullpath)
			fmt.Println(startingdir.Extension)
			fmt.Println(startingdir.Bytes)
			fmt.Println(startingdir.IsDir)
			for _, entry := range startingdir.Sub {
				fmt.Println(entry)
			}
		} else {
			fmt.Println("Error: " + args[0] + " does not exist.")
		}
		fmt.Println("\nSTOPPING")
	},
}

func recursiveScan(dirstring string) util.File {
	localdir := util.File{
		Name:      path.Base(dirstring),
		Fullpath:  dirstring,
		Extension: "",
		Bytes:     0,
		IsDir:     true,
		Sub:       nil,
	}
	direntries, _ := os.ReadDir(dirstring)
	for _, entry := range direntries {
		abs, _ := filepath.Abs(entry.Name())
		if !entry.IsDir() {
			filedata, _ := entry.Info()
			newfile := util.File{
				Name:      filedata.Name(),
				Fullpath:  abs,
				Extension: path.Ext(abs),
				Bytes:     filedata.Size(),
				IsDir:     filedata.IsDir(),
				Sub:       nil,
			}
			localdir.Sub = append(localdir.Sub, newfile)
		} else {
			localdir.Sub = append(localdir.Sub, recursiveScan(abs))
		}
	}
	for _, entry := range localdir.Sub {
		localdir.Bytes += entry.Bytes
	}
	return localdir
}
