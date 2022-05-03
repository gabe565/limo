package action

//go:generate stringer -type=Action

type Action int64

const (
	Unknown Action = iota
	Upload
	View
	Delete
)
