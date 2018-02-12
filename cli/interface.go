package cli

type CLI interface {
	GetRowsAsString() string
	Exec(id string) error
	PickUpID(row string) string
	UpdateRows(id string)
	IsExitCLI(row string) bool
}
