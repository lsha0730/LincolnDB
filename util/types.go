package util

const (
	READ           string = "read"
	WRITE                 = "write"
	LIST                  = "list"
	MAKECOLLECTION        = "newcol"
	MAKEDOCUMENT          = "newdoc"
)

type Query interface{}

type ReadQuery struct {
	op   string
	path string
}

type WriteQuery struct {
	op    string
	path  string
	value interface{}
}

type ListQuery struct {
	op   string
	path string
}

type MakeQuery struct {
	op   string
	path string
	name string
}
