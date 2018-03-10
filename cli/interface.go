package cli

type CLI interface {
	SetID(row string)
	GetRowsAsString() string
	UpdateRows()
	Exec() string
	Output() string
	isOnceCLI() bool
}
