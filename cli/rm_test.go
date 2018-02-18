package cli

import (
	"testing"
)

func newRmCLI() *RmCLI {
	rows := []string{
		"2e3aa87d412b  sample_mysql  \"docker-entrypoint...\"  4 days ago Exited (0) 23 hours ago  sample_mysql_1",
		"560f277ebacb  sample_redis  \"docker-entrypoint...\"  2 days ago Exited (0) 23 hours ago  sample_redis_1",
	}
	return &RmCLI{rows: rows}
}

func TestRmCLIGetRowsAsString(t *testing.T) {
	execCLI := newRmCLI()
	actual := execCLI.GetRowsAsString()
	expected := "2e3aa87d412b  sample_mysql  \"docker-entrypoint...\"  4 days ago Exited (0) 23 hours ago  sample_mysql_1\n560f277ebacb  sample_redis  \"docker-entrypoint...\"  2 days ago Exited (0) 23 hours ago  sample_redis_1"
	if actual != expected {
		t.Errorf("Invalid operation: %#v == %#v", expected, actual)
	}
}

func TestRmCLISetID(t *testing.T) {
	execCLI := newRmCLI()
	execCLI.SetID(execCLI.rows[0])
	actual := execCLI.id
	expected := "2e3aa87d412b"
	if actual != expected {
		t.Error("Invalid operation: %#v == %#v", expected, actual)
	}
}

func TestRmCLIUpdateRows(t *testing.T) {
	execCLI := newRmCLI()
	execCLI.SetID(execCLI.rows[0])
	execCLI.UpdateRows()
	actual := len(execCLI.rows)
	expected := 1
	if actual != expected {
		t.Error("Invalid operation: %#v == %#v", expected, actual)
	}
}

func TestRmCLIOutput(t *testing.T) {
	execCLI := newRmCLI()
	actual := execCLI.Output()
	expected := ""
	if actual != expected {
		t.Error("Invalid operation: %#v == %#v", expected, actual)
	}
}

func TestRmCLIisOnceCLI(t *testing.T) {
	execCLI := newRmCLI()
	actual := execCLI.isOnceCLI()
	expected := false
	if actual != expected {
		t.Error("Invalid operation: %#v == %#v", expected, actual)
	}
}
