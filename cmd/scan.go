package cmd

import (
	"fmt"
	"os"
	"os/user"
	"path/filepath"
	"strconv"
	"syscall"

	"github.com/holden-davis/DirDepict/util"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var verbose bool

func init() {
	scanCMD.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "verbose output")
	viper.BindPFlag("verbose", scanCMD.PersistentFlags().Lookup("verbose"))
	rootCmd.AddCommand(scanCMD)
}

var scanCMD = &cobra.Command{
	Use:     "scan",
	Short:   "Scans a directory or file",
	Long:    "Scans file or recursively scans directory, printing an overview of the results",
	Example: "dirdepict scan .",
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		rootpath, _ := filepath.Abs(args[0])
		_, err := os.Stat(rootpath)
		if err != nil {
			fmt.Println(args[0], "is not a valid file or directory!")
			return
		}
		root := scan(rootpath)
		root.Print(viper.GetBool("verbose"), 0)
	},
}

func scan(path string) util.File {
	root := util.File{}
	fiSYM, _ := os.Lstat(path)
	root.SymLink = fiSYM.Mode()&os.ModeSymlink != 0
	fi, _ := os.Stat(path)
	root.Fullpath = path
	root.Name = filepath.Base(path)
	root.Extension = filepath.Ext(path)
	root.Permissions = fi.Mode().String()
	root.Mod = fi.ModTime().String()
	fileUID := fi.Sys().(*syscall.Stat_t).Uid
	fileGID := fi.Sys().(*syscall.Stat_t).Gid
	fileUser, _ := user.LookupId(strconv.Itoa(int(fileUID)))
	fileGroup, _ := user.LookupGroupId(strconv.Itoa(int(fileGID)))
	root.User = fileUser.Username
	root.Group = fileGroup.Name
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
