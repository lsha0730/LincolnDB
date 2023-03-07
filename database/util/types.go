package util

// Basic types
type Driver struct {
	name string
	root string
}

// Query types
const (
	READ   string = "read"
	WRITE         = "write"
	LIST          = "list"
	MAKEDB        = "newdb"
)

type ReadQuery struct {
	db   string
	op   string
	path string
}

type WriteQuery struct {
	db    string
	op    string
	path  string
	value interface{}
}

type ListQuery struct {
	db   string
	op   string
	path string
}

type MakeQuery struct {
	op   string
	name string
}
