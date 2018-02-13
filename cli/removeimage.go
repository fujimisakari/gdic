package cli

import (
	"bufio"
	"os/exec"
	"strings"

	"github.com/pkg/errors"
)

type RemoveImageCLI struct {
	id   string
	rows []string
}

func (c *RemoveImageCLI) SetID(row string) {
	c.id = c.pickUpID(row)
}

func (c *RemoveImageCLI) GetRowsAsString() string {
	return strings.Join(c.rows, "\n")
}

func (c *RemoveImageCLI) UpdateRows() {
	newRows := []string{}
	for i := range c.rows {
		id := c.pickUpID(c.rows[i])
		if c.id != id {
			newRows = append(newRows, c.rows[i])
		}
	}
	c.rows = newRows
}

func (c *RemoveImageCLI) pickUpID(row string) string {
	r := strings.Fields(row)
	return r[2]
}

func (c *RemoveImageCLI) Exec() error {
	return exec.Command("docker", "rmi", c.id).Run()
}

func (c *RemoveImageCLI) Output() string {
	return ""
}

func (c *RemoveImageCLI) isOnceCLI() bool {
	return false
}

func NewRemoveImageCLI() (*RemoveImageCLI, error) {
	cmd := exec.Command("docker", "images")
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return nil, errors.Wrap(err, "NewRemoveImageCLI Error")
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
	return &RemoveImageCLI{rows: rows}, nil
}
