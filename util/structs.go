package util

// Struct for holding information about a file
type File struct {
	//Name of the file ex. data.txt
	Name string
	//Full path of the file ex. /home/user/nobodyatall/desktop/data.txt
	Fullpath string
	//Extension type of the file, if one exists ex. txt
	Extension string
	//Size of the file in bytes ex. 2342
	Size int64
	//Flag to separate files from directories
	IsDir bool
	//A collection of file types that represent the subdirectory of the current dir
	Sub []File
}
