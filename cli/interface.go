package cli

type CLI interface {
	SetID(row string)
	GetRowsAsString() string
	UpdateRows()
	Exec() error
	Output() string
	isOnceCLI() bool
}
