package stuff

import "fmt"

// конкретный прототип

type Folder struct {
	Children []Inode
	Name     string
}

func (f *Folder) Print(identation string) {
	fmt.Println(identation + f.Name)
	for _, i := range f.Children {
		i.Print(identation + identation)
	}
}

func (f *Folder) Clone() Inode {
	cloneFolder := &Folder{Name: f.Name + "_clone"}
	var tempChildren []Inode
	for _, i := range f.Children {
		copy := i.Clone()
		tempChildren = append(tempChildren, copy)
	}

	cloneFolder.Children = tempChildren

	return cloneFolder
}
