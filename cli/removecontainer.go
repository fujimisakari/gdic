package cli

import (
	"bufio"
	"os/exec"
	"strings"

	"github.com/pkg/errors"
)

type RemoveContainerCLI struct {
	id   string
	rows []string
}

func (c *RemoveContainerCLI) SetID(row string) {
	c.id = c.pickUpID(row)
}

func (c *RemoveContainerCLI) GetRowsAsString() string {
	return strings.Join(c.rows, "\n")
}

func (c *RemoveContainerCLI) UpdateRows() {
	newRows := []string{}
	for i := range c.rows {
		id := c.pickUpID(c.rows[i])
		if c.id != id {
			newRows = append(newRows, c.rows[i])
		}
	}
	c.rows = newRows
}

func (c *RemoveContainerCLI) pickUpID(row string) string {
	r := strings.Fields(row)
	return r[0]
}

func (c *RemoveContainerCLI) Exec() error {
	return exec.Command("docker", "rm", c.id).Run()
}

func (c *RemoveContainerCLI) Output() string {
	return ""
}

func (c *RemoveContainerCLI) isOnceCLI() bool {
	return false
}

func NewRemoveContainerCLI() (*RemoveContainerCLI, error) {
	cmd := exec.Command("docker", "ps", "-a")
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return nil, errors.Wrap(err, "NewRemoveContainerCLI Error")
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
	return &RemoveContainerCLI{rows: rows}, nil
}
