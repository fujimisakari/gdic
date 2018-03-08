package cli

import (
	"bufio"
	"os/exec"
	"strings"

	"github.com/pkg/errors"
)

type StopCLI struct {
	id   string
	rows []string
}

func (c *StopCLI) SetID(row string) {
	c.id = c.pickUpID(row)
}

func (c *StopCLI) GetRowsAsString() string {
	return strings.Join(c.rows, "\n")
}

func (c *StopCLI) UpdateRows() {
	newRows := []string{}
	for i := range c.rows {
		id := c.pickUpID(c.rows[i])
		if c.id != id {
			newRows = append(newRows, c.rows[i])
		}
	}
	c.rows = newRows
}

func (c *StopCLI) pickUpID(row string) string {
	r := strings.Fields(row)
	return r[0]
}

func (c *StopCLI) Exec() error {
	return exec.Command("docker", "stop", c.id).Run()
}

func (c *StopCLI) Output() string {
	return ""
}

func (c *StopCLI) isOnceCLI() bool {
	return false
}

func NewStopCLI() (*StopCLI, error) {
	cmd := exec.Command("docker", "ps")
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return nil, errors.Wrap(err, "NewStopCLI Error")
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
	return &StopCLI{rows: rows}, nil
}
