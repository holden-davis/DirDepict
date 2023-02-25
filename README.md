# DirDepict
CLI program for analyzing file system storage consumption written in Go

## Example desired output of `dirdepict scan .` assuming currently in ~:

Scanned folder: /home/user/nobodyatall/

Items: 5 (3 dir, 2 file)

Size: 1.75 GB (3% of disk sda1)

| Item      | Type | Size    | Percentage |
|-----------|------|---------|------------|
| /folder1  | dir  | 1.0 GB  | 70%        |
| /folder2  | dir  | 0.5 GB  | 20%        |
| /folder3  | file | 0.25 GB | 10%        |
| script.sh | file | 20 KB   | 0.2%       |
| data.txt  | file | 1 B     | 0.6%       |

### Concept

Recursively search through directories, collecting information on files and directories therein

Need a data structure to hold recursive data
Information on a file should just be a structure: it needs to have string for name, string for type, string for size
Information on a folder needs the same fields but a list of folder constructs as well for recursion

Need a recursive function that gets called for each directory

Function should return a Dir struct representing the directory

At the beginning of the function, collect the path name and filename of the directory

For each file, create a File struct and add it to the Sub slice

For each subdirectory, call the function again: eventually it will return a Dir struct which can be added to Sub

Before returning, the Dir should set its size by summing the size of everything in the Sub collection

What does each function call need?
* Either the full path of the directory or just the name of the directory, just something to generate a DirEntry object from