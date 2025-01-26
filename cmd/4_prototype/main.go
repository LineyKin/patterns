package main

import (
	"fmt"
	"patterns/cmd/4_prototype/stuff"
)

func main() {
	file1 := &stuff.File{Name: "file1"}
	file2 := &stuff.File{Name: "file2"}
	file3 := &stuff.File{Name: "file3"}

	folder1 := &stuff.Folder{
		Children: []stuff.Inode{file1},
		Name:     "Folder1",
	}

	folder2 := &stuff.Folder{
		Children: []stuff.Inode{folder1, file2, file3},
		Name:     "Folder2",
	}

	fmt.Println("\nFolder2 hierarchy:")
	folder2.Print(" ")

	cloneFolder2 := folder2.Clone()
	fmt.Println("\nFolder2 clone hierarchy:")
	cloneFolder2.Print(" ")
}
