package cli

import (
	"testing"
)

func newExecCLI() *ExecCLI {
	rows := []string{
		"2e3aa87d412b  sample_mysql  \"docker-entrypoint...\"  4 days ago Exited (0) 23 hours ago  sample_mysql_1",
		"560f277ebacb  sample_redis  \"docker-entrypoint...\"  2 days ago Exited (0) 23 hours ago  sample_redis_1",
	}
	return &ExecCLI{rows: rows}
}

func TestExecCLIGetRowsAsString(t *testing.T) {
	execCLI := newExecCLI()
	actual := execCLI.GetRowsAsString()
	expected := "2e3aa87d412b  sample_mysql  \"docker-entrypoint...\"  4 days ago Exited (0) 23 hours ago  sample_mysql_1\n560f277ebacb  sample_redis  \"docker-entrypoint...\"  2 days ago Exited (0) 23 hours ago  sample_redis_1"
	if actual != expected {
		t.Errorf("Invalid operation: %#v == %#v", expected, actual)
	}
}

func TestExecCLISetID(t *testing.T) {
	execCLI := newExecCLI()
	execCLI.SetID(execCLI.rows[0])
	actual := execCLI.id
	expected := "sample_mysql_1"
	if actual != expected {
		t.Error("Invalid operation: %#v == %#v", expected, actual)
	}
}

func TestExecCLIOutput(t *testing.T) {
	execCLI := newExecCLI()
	execCLI.SetID(execCLI.rows[0])
	actual := execCLI.Output()
	expected := "docker exec -it sample_mysql_1 /bin/bash"
	if actual != expected {
		t.Error("Invalid operation: %#v == %#v", expected, actual)
	}
}

func TestExecCLIisOnceCLI(t *testing.T) {
	execCLI := newExecCLI()
	actual := execCLI.isOnceCLI()
	expected := true
	if actual != expected {
		t.Error("Invalid operation: %#v == %#v", expected, actual)
	}
}
