package cli

import (
	"bufio"
	"fmt"
	"os/exec"
	"strings"

	"github.com/pkg/errors"
)

type ExecContainerCLI struct {
	id   string
	rows []string
}

func (c *ExecContainerCLI) SetID(row string) {
	r := strings.Fields(row)
	last := len(r) - 1
	c.id = r[last]
}

func (c *ExecContainerCLI) GetRowsAsString() string {
	return strings.Join(c.rows, "\n")
}

func (c *ExecContainerCLI) UpdateRows() {}

func (c *ExecContainerCLI) Exec() error {
	return nil
}

func (c *ExecContainerCLI) Output() string {
	if c.id != "" {
		return fmt.Sprintf("docker exec -it %s /bin/bash", c.id)
	}
	return ""
}

func (c *ExecContainerCLI) isOnceCLI() bool {
	return true
}

func NewExecContainerCLI() (*ExecContainerCLI, error) {
	cmd := exec.Command("docker", "ps", "-a")
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return nil, errors.Wrap(err, "NewExecContainerCLI Error")
	}

	var i uint64
	rows := []string{}
	cmd.Start()
	scanner := bufio.NewScanner(stdout)
	for scanner.Scan() {
		if i == 0 {
			i++
			continue
		}
		rows = append(rows, scanner.Text())
	}
	cmd.Wait()
	return &ExecContainerCLI{rows: rows}, nil
}
