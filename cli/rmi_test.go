package cli

import (
	"testing"
)

func newRmiCLI() *RmiCLI {
	rows := []string{
		"mysql   5.7     7d83a47ab2d2  2 months ago   408 MB",
		"<none>  <none>  cec50792cfe0  5 minutes ago  168 MB",
	}
	return &RmiCLI{rows: rows}
}

func TestRmiCLIGetRowsAsString(t *testing.T) {
	execCLI := newRmiCLI()
	actual := execCLI.GetRowsAsString()
	expected := "mysql   5.7     7d83a47ab2d2  2 months ago   408 MB\n<none>  <none>  cec50792cfe0  5 minutes ago  168 MB"
	if actual != expected {
		t.Errorf("Invalid operation: %#v == %#v", expected, actual)
	}
}

func TestRmiCLISetID(t *testing.T) {
	execCLI := newRmiCLI()
	execCLI.SetID(execCLI.rows[0])
	actual := execCLI.id
	expected := "mysql"
	if actual != expected {
		t.Error("Invalid operation: %#v == %#v", expected, actual)
	}
}

func TestRmiCLISetIDNone(t *testing.T) {
	execCLI := newRmiCLI()
	execCLI.SetID(execCLI.rows[1])
	actual := execCLI.id
	expected := "cec50792cfe0"
	if actual != expected {
		t.Error("Invalid operation: %#v == %#v", expected, actual)
	}
}

func TestRmiCLIUpdateRows(t *testing.T) {
	execCLI := newRmiCLI()
	execCLI.SetID(execCLI.rows[0])
	execCLI.UpdateRows()
	actual := len(execCLI.rows)
	expected := 1
	if actual != expected {
		t.Error("Invalid operation: %#v == %#v", expected, actual)
	}
}

func TestRmiCLIOutput(t *testing.T) {
	execCLI := newRmiCLI()
	actual := execCLI.Output()
	expected := ""
	if actual != expected {
		t.Error("Invalid operation: %#v == %#v", expected, actual)
	}
}

func TestRmiCLIisOnceCLI(t *testing.T) {
	execCLI := newRmiCLI()
	actual := execCLI.isOnceCLI()
	expected := false
	if actual != expected {
		t.Error("Invalid operation: %#v == %#v", expected, actual)
	}
}
