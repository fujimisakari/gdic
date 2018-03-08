package cli

import (
	"bufio"
	"fmt"
	"os/exec"
	"strings"

	"github.com/pkg/errors"
)

type ExecCLI struct {
	id   string
	rows []string
}

func (c *ExecCLI) SetID(row string) {
	r := strings.Fields(row)
	last := len(r) - 1
	c.id = r[last]
}

func (c *ExecCLI) GetRowsAsString() string {
	return strings.Join(c.rows, "\n")
}

func (c *ExecCLI) UpdateRows() {}

func (c *ExecCLI) Exec() error {
	return nil
}

func (c *ExecCLI) Output() string {
	if c.id != "" {
		return fmt.Sprintf("docker exec -it %s /bin/bash", c.id)
	}
	return ""
}

func (c *ExecCLI) isOnceCLI() bool {
	return true
}

func NewExecCLI() (*ExecCLI, error) {
	cmd := exec.Command("docker", "ps")
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return nil, errors.Wrap(err, "NewExecCLI Error")
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
	return &ExecCLI{rows: rows}, nil
}
