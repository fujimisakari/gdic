package cli

import (
	"bufio"
	"os/exec"
	"strings"

	"github.com/pkg/errors"
)

type StopContainerCLI struct {
	id   string
	rows []string
}

func (c *StopContainerCLI) SetID(row string) {
	c.id = c.pickUpID(row)
}

func (c *StopContainerCLI) GetRowsAsString() string {
	return strings.Join(c.rows, "\n")
}

func (c *StopContainerCLI) UpdateRows() {
	newRows := []string{}
	for i := range c.rows {
		id := c.pickUpID(c.rows[i])
		if c.id != id {
			newRows = append(newRows, c.rows[i])
		}
	}
	c.rows = newRows
}

func (c *StopContainerCLI) pickUpID(row string) string {
	r := strings.Fields(row)
	return r[0]
}

func (c *StopContainerCLI) Exec() error {
	return exec.Command("docker", "stop", c.id).Run()
}

func (c *StopContainerCLI) Output() string {
	return ""
}

func (c *StopContainerCLI) isOnceCLI() bool {
	return false
}

func NewStopContainerCLI() (*StopContainerCLI, error) {
	cmd := exec.Command("docker", "ps", "-a")
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return nil, errors.Wrap(err, "NewStopContainerCLI Error")
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
	return &StopContainerCLI{rows: rows}, nil
}
