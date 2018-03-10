package cli

import (
	"bufio"
	"fmt"
	"os/exec"
	"strings"

	"github.com/pkg/errors"
)

type RunCLI struct {
	id   string
	rows []string
}

func (c *RunCLI) SetID(row string) {
	r := strings.Fields(row)
	c.id = fmt.Sprintf("%s:%s", r[0], r[1])
}

func (c *RunCLI) GetRowsAsString() string {
	return strings.Join(c.rows, "\n")
}

func (c *RunCLI) UpdateRows() {}

func (c *RunCLI) Exec() string {
	return ""
}

func (c *RunCLI) Output() string {
	if c.id != "" {
		return fmt.Sprintf("docker run -it %s /bin/bash", c.id)
	}
	return ""
}

func (c *RunCLI) isOnceCLI() bool {
	return true
}

func NewRunCLI() (*RunCLI, error) {
	cmd := exec.Command("docker", "images")
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return nil, errors.Wrap(err, "NewRunCLI Error")
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
	return &RunCLI{rows: rows}, nil
}
