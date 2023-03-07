package util

import (
	"fmt"
	"strings"
)

// Struct for holding information about a file
type File struct {
	// Name of the file ex. data.txt
	Name string
	// Full path of the file ex. /home/user/nobodyatall/desktop/data.txt
	Fullpath string
	// Extension type of the file, if one exists ex. txt
	Extension string
	// Size of the file in bytes ex. 2342
	Size int64
	// Flag to separate files from directories
	IsDir bool
	// A collection of file types that represent the subdirectory of the current dir
	Sub []File
	// Representation of Unix-style permissions display
	Permissions string
	// Representation of latest file modification time
	Mod string
	// Representation of file user
	User string
	// Representation of file group
	Group string
	// Flag to identify symlinks
	SymLink bool
}

// Simple function to print file info
func (file File) Print(verbose bool, indent int) {
	fmt.Println(strings.Repeat("-", 30))
	indentation := strings.Repeat("--", indent*2) + "|"
	fmt.Println(indentation, "Name:          ", file.Name)
	fmt.Println(indentation, "Absolute Path: ", file.Fullpath)
	fmt.Println(indentation, "Size:          ", file.Size)
	fmt.Println(indentation, "Permissions:   ", file.Permissions)
	fmt.Println(indentation, "Modified:      ", file.Mod)
	fmt.Println(indentation, "User:          ", file.User)
	fmt.Println(indentation, "Group:         ", file.Group)
	fmt.Println(indentation, "Symlink?       ", file.SymLink)
	fmt.Println(indentation, "Directory?     ", file.IsDir)
	if file.IsDir {
		fmt.Println(indentation, "Contents Count:", len(file.Sub))
	} else {
		fmt.Println(indentation, "Extension:     ", file.Extension)
	}
	if verbose {
		for _, entry := range file.Sub {
			entry.Print(verbose, indent+1)
		}
	}
}
