package stuff

// интерфейс прототипа

type Inode interface {
	Print(string)
	Clone() Inode
}
