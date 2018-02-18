package cli

import (
	"testing"
)

func newRunCLI() *RunCLI {
	rows := []string{
		"mysql  5.7     7d83a47ab2d2  2 months ago  408 MB",
		"redis  latest  aaf79d45ddb1  5 months ago  107 MB",
	}
	return &RunCLI{rows: rows}
}

func TestRunCLIGetRowsAsString(t *testing.T) {
	runCLI := newRunCLI()
	actual := runCLI.GetRowsAsString()
	expected := "mysql  5.7     7d83a47ab2d2  2 months ago  408 MB\nredis  latest  aaf79d45ddb1  5 months ago  107 MB"
	if actual != expected {
		t.Errorf("Invalid operation: %#v == %#v", expected, actual)
	}
}

func TestRunCLISetID(t *testing.T) {
	runCLI := newRunCLI()
	runCLI.SetID(runCLI.rows[0])
	actual := runCLI.id
	expected := "mysql:5.7"
	if actual != expected {
		t.Error("Invalid operation: %#v == %#v", expected, actual)
	}
}

func TestRunCLIOutput(t *testing.T) {
	runCLI := newRunCLI()
	runCLI.SetID(runCLI.rows[0])
	actual := runCLI.Output()
	expected := "docker run -it mysql:5.7 /bin/bash"
	if actual != expected {
		t.Error("Invalid operation: %#v == %#v", expected, actual)
	}
}

func TestRunCLIisOnceCLI(t *testing.T) {
	runCLI := newRunCLI()
	actual := runCLI.isOnceCLI()
	expected := true
	if actual != expected {
		t.Error("Invalid operation: %#v == %#v", expected, actual)
	}
}
