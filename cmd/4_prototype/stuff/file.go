package stuff

import "fmt"

// конкретный прототип

type File struct {
	Name string
}

func (f *File) Print(identation string) {
	fmt.Println(identation + f.Name)
}

func (f *File) Clone() Inode {
	return &File{
		Name: f.Name + "_clone",
	}
}
