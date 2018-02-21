package cli

import (
	"testing"
)

func newStopCLI() *StopCLI {
	rows := []string{
		"2e3aa87d412b  sample_mysql  \"docker-entrypoint...\"  4 days ago Exited (0) 23 hours ago  sample_mysql_1",
		"560f277ebacb  sample_redis  \"docker-entrypoint...\"  2 days ago Exited (0) 23 hours ago  sample_redis_1",
	}
	return &StopCLI{rows: rows}
}

func TestStopCLISetID(t *testing.T) {
	execCLI := newStopCLI()
	execCLI.SetID(execCLI.rows[0])
	actual := execCLI.id
	expected := "2e3aa87d412b"
	if actual != expected {
		t.Error("Invalid operation: %#v == %#v", expected, actual)
	}
}

func TestStopCLIGetRowsAsString(t *testing.T) {
	execCLI := newStopCLI()
	actual := execCLI.GetRowsAsString()
	expected := "2e3aa87d412b  sample_mysql  \"docker-entrypoint...\"  4 days ago Exited (0) 23 hours ago  sample_mysql_1\n560f277ebacb  sample_redis  \"docker-entrypoint...\"  2 days ago Exited (0) 23 hours ago  sample_redis_1"
	if actual != expected {
		t.Errorf("Invalid operation: %#v == %#v", expected, actual)
	}
}

func TestStopCLIUpdateRows(t *testing.T) {
	execCLI := newStopCLI()
	execCLI.SetID(execCLI.rows[0])
	execCLI.UpdateRows()
	actual := len(execCLI.rows)
	expected := 1
	if actual != expected {
		t.Error("Invalid operation: %#v == %#v", expected, actual)
	}
}

func TestStopCLIOutput(t *testing.T) {
	execCLI := newStopCLI()
	actual := execCLI.Output()
	expected := ""
	if actual != expected {
		t.Error("Invalid operation: %#v == %#v", expected, actual)
	}
}

func TestStopCLIisOnceCLI(t *testing.T) {
	execCLI := newStopCLI()
	actual := execCLI.isOnceCLI()
	expected := false
	if actual != expected {
		t.Error("Invalid operation: %#v == %#v", expected, actual)
	}
}
